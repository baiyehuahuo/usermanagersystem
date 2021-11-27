package main

import (
	"log"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/service/htmlcontrol"
	"usermanagersystem/service/logincontrol"
	"usermanagersystem/service/usercontrol"
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := utils.ConfigRead(); err != nil {
		log.Fatal(err)
	}

	if err := utils.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	if err := utils.ConnectToRedis(); err != nil {
		log.Fatal(err)
	}

	if err := utils.NewCache(consts.AuthCodeContinueTime, consts.AuthCodeCacheFlushTime); err != nil {
		log.Fatal(err)
	}

	if err := UploadFilePathCreate(); err != nil {
		log.Fatal(err)
	}
	utils.EmailAuthCodeControllerCreate()

}

func main() {
	router := gin.Default()
	htmlManager := htmlcontrol.New()
	router.LoadHTMLGlob("templates/*")   // html 文件
	router.Static("/static", "./static") // 静态文件映射
	router.Static("/avatar", "./uploadfiles/avatars")
	router.GET("/", htmlManager.ToLogin)
	router.GET("/UserManage", htmlManager.ToUserManage)

	handle := handleManager{
		lm: logincontrol.New(),
		um: usercontrol.New(),
	}

	router.GET("/UserLogin", handle.UserLogin)
	router.GET("/GetUserMessage", handle.GetUserMessageByCookie)
	router.GET("/UserRegedit", handle.UserRegedit)

	router.POST("/UploadAvatar", handle.UploadAvatar)
	router.POST("/UploadFile", handle.UploadFile)
	router.POST("/ModifyPassword", handle.ModifyPassword)

	user := model.User{
		Account:   "000",
		Password:  "",
		Email:     "1770194225@qq.com",
		NickName:  "范伟锋",
		AvatarExt: "c6f057b86584942e415435ffb1fa93d4",
	}
	log.Println(user)
	utils.RedisNew().SetUser(user)
	utils.RedisNew().GetUser("000")
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

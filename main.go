package main

import (
	"log"
	"usermanagersystem/consts"
	"usermanagersystem/service/htmlcontrol"
	"usermanagersystem/service/logincontrol"
	"usermanagersystem/service/usercontrol"
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	var err error
	if err = utils.ConfigRead(); err != nil {
		log.Fatal(err)
	}

	if err = utils.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	if err = utils.ConnectToRedis(); err != nil {
		log.Fatal(err)
	}

	if err = utils.NewCache(consts.AuthCodeContinueTime, consts.AuthCodeCacheFlushTime); err != nil {
		log.Fatal(err)
	}

	if err = UploadFilePathCreate(); err != nil {
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
	router.Static("/movie", "./uploadfiles/movies")
	router.GET("/", htmlManager.ToLogin)
	router.GET("/UserManage", htmlManager.ToUserManage)

	handle := handleManager{
		lm: logincontrol.New(),
		um: usercontrol.New(),
	}

	router.GET("/CheckAuthCode", handle.CheckAuthCode)
	router.GET("/GetUserMessage", handle.GetUserMessageByCookie)
	router.GET("/UserLogin", handle.UserLogin)
	router.GET("/UserRegedit", handle.UserRegedit)
	router.GET("/SendAuthCode", handle.SendAuthCode)

	router.POST("/UploadAvatar", handle.UploadAvatar)
	router.POST("/UploadFile", handle.UploadFile)
	router.POST("/ModifyPassword", handle.ModifyPassword)
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

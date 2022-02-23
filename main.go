package main

import (
	"log"
	"os"
	"usermanagersystem/consts"
	"usermanagersystem/service/html_control"
	"usermanagersystem/service/login_control"
	"usermanagersystem/service/user_control"
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	var err error

	if err = SetLog(); err != nil {
		log.Fatal(err)
	}

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

	if err = SetTimer(); err != nil {
		log.Fatal(err)
	}

	utils.EmailAuthCodeControllerCreate()
}

func main() {
	router := gin.Default()
	router.Use(Cors())
	htmlManager := html_control.New()
	router.LoadHTMLGlob("templates/*")                                // html 文件
	router.Static(consts.DefaultStaticPath, consts.DefaultStaticPath) // 静态文件映射
	router.Static(consts.DefaultAvatarPath, consts.DefaultAvatarPath)
	// router.Static("/movie", "./uploadfiles/movies")
	router.GET("/", htmlManager.ToLogin)
	router.GET("/UserManage", htmlManager.ToUserManage)

	handle := handleManager{
		lm: login_control.New(),
		um: user_control.New(),
	}
	router.GET("/CheckAuthCode", handle.CheckAuthCode)
	router.GET("/CheckEmailAvailable", handle.CheckEmailAvailable)
	router.GET("/GetUserMessage", handle.GetUserMessageByCookie)
	router.GET("/RestoreMySQL", handle.RestoreMySQL)
	router.GET("/UserLogin", handle.UserLogin)
	router.GET("/UserRegister", handle.UserRegister)
	router.GET("/SendAuthCode", handle.SendAuthCode)

	router.POST("/ModifyPassword", handle.ModifyPassword)
	router.POST("/UploadAvatar", handle.UploadAvatar)
	router.POST("/UploadFile", handle.UploadFile)
	if err := os.MkdirAll(consts.DefaultAvatarPath, os.ModePerm); err != nil {
		log.Fatal("目录创建失败 ", err)
	}
	// log.Println(utils.GetDB().ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Where(model.User{Email: "6@qq.com"}).First(&model.User{})
	// }))
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

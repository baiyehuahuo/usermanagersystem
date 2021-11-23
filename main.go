package main

import (
	"log"
	"usermanagersystem/service/filecontrol"
	"usermanagersystem/service/htmlcontrol"
	"usermanagersystem/service/logincontrol"
	"usermanagersystem/service/passwordcontrol"
	"usermanagersystem/service/regeditcontrol"
	"usermanagersystem/utils/configread"
	"usermanagersystem/utils/databasecontrol"
	"usermanagersystem/utils/emailauthcode"
	"usermanagersystem/utils/rediscontrol"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := configread.ConfigRead(); err != nil {
		log.Fatal(err)
	}

	if err := databasecontrol.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	if err := rediscontrol.ConnectToRedis(); err != nil {
		log.Fatal(err)
	}

	emailauthcode.EmailAuthCodeControllerCreate()
}

func main() {
	router := gin.Default()
	htmlManager := htmlcontrol.New()
	router.LoadHTMLGlob("templates/*")   // html 文件
	router.Static("/static", "./static") // 静态文件映射
	router.GET("/", htmlManager.ToLogin)
	router.GET("/UserManage", htmlManager.ToUserManage)

	handle := handleManager{
		loginManager:       logincontrol.New(),
		regeditManager:     regeditcontrol.New(),
		fileControlManager: filecontrol.New(),
		passwordManager:    passwordcontrol.New(),
	}
	router.GET("/UserLogin", handle.UserLogin)
	router.GET("/UserRegedit", handle.UserRegedit)
	router.POST("/UploadFile", handle.FileUpload)
	router.POST("/ModifyPassword", handle.ModifyPassword)
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

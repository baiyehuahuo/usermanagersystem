package main

import (
	"log"
	"usermanagersystem/service/filecontrol"
	"usermanagersystem/service/html"
	"usermanagersystem/service/login"
	"usermanagersystem/service/regedit"
	"usermanagersystem/utils/configReader"
	"usermanagersystem/utils/database"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := configReader.ConfigRead(); err != nil {
		log.Fatal(err)
	}
	if err := database.ConnectDatabase(); err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	htmlManager := html.New()
	router.LoadHTMLGlob("templates/*")   // html 文件
	router.Static("/static", "./static") // 静态文件映射
	router.GET("/", htmlManager.ToLogin)
	router.GET("/UserManage", htmlManager.ToUserManage)

	handle := handleManager{
		loginManager:       login.New(),
		regeditManager:     regedit.New(),
		fileControlManager: filecontrol.New(),
	}
	router.GET("/UserLogin", handle.UserLogin)
	router.GET("/UserRegedit", handle.UserRegedit)
	router.POST("/UploadFile", handle.FileUpload)
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

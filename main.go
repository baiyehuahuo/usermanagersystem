package main

import (
	"log"
	"usermanagersystem/service/html"
	"usermanagersystem/service/login"
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
	loginManager := login.New()
	router.LoadHTMLGlob("templates/*")   // html 文件
	router.Static("/static", "./static") // 静态文件映射
	router.GET("/", htmlManager.ToLogin)
	router.GET("/UserLogin", loginManager.UserLogin)
	router.Run()
}

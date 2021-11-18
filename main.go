package main

import (
	"UserManageSystem/service/html"
	"UserManageSystem/service/login"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	htmlManager := html.New()
	loginManager := login.New()
	router.LoadHTMLGlob("templates/*")   // html 文件
	router.Static("/static", "./static") // 静态文件映射
	router.GET("/", htmlManager.ToLogin)
	router.GET("/HelloWorld", loginManager.UserLogin)
	router.Run()
}

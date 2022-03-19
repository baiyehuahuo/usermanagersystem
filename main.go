package main

import (
	"log"
	"mime"
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

	if err = UploadFilePathCreate(); err != nil {
		log.Fatal(err)
	}

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
	router.Static(consts.DefaultUserPngRootPath, consts.DefaultUserPngRootPath)
	_ = mime.AddExtensionType(".js", "application/javascript")
	router.GET("/", htmlManager.ToLogin)

	handle := handleManager{
		lm: login_control.New(),
		um: user_control.New(),
	}

	router.GET("/CheckEmailAvailable", handle.CheckEmailAvailable)
	router.GET("/GetUserFilesPath", handle.GetUserFilesPath)
	router.GET("/GetUserMessage", handle.GetUserMessageByCookie)
	router.GET("/PredictPng", handle.PredictPng)
	router.GET("/RestoreMySQL", handle.RestoreMySQL)
	router.GET("/UserLogin", handle.UserLogin)
	router.GET("/UserRegister", handle.UserRegister)
	router.GET("/SendAuthCode", handle.SendAuthCode)

	router.POST("/ModifyPassword", handle.ModifyPassword)
	router.POST("/ForgetPassword", handle.ForgetPassword)
	router.POST("/UploadAvatar", handle.UploadAvatar)
	router.POST("/UploadPng", handle.UploadPng)
	router.POST("/DeletePng", handle.DeletePng)
	if err := os.MkdirAll(consts.DefaultAvatarPath, os.ModePerm); err != nil {
		log.Fatal("目录创建失败 ", err)
	}
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

// todo 添加一个删除文件的接口
// todo rabbitmq与模型识别结合 减去加载模型的时间
// todo 丰富错误信息处理

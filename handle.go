package main

import (
	"log"
	"net/http"
	"os"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/service/logincontrol"
	"usermanagersystem/service/usercontrol"

	"github.com/gin-gonic/gin"
)

type handleManager struct {
	lm logincontrol.LoginController
	um usercontrol.UserController
}

// CheckAuthCode 验证码检测处理接口
func (handle *handleManager) CheckAuthCode(c *gin.Context) {
	if err := handle.lm.CheckAuthCode(c); err != nil {
		log.Printf("CheckAuthCode Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.CheckAuthCodeFail)
		return
	}
	c.JSON(http.StatusOK, consts.CheckAuthCodeSuccess)
}

// GetUserMessageByCookie 通过Cookie获取用户信息处理接口
func (handle *handleManager) GetUserMessageByCookie(c *gin.Context) {
	user, err := handle.um.GetUserMessageByCookie(c)
	if err != nil {
		log.Printf("GetUserMessage Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.GetUserMessageFail)
		return
	}
	result := model.UserMessage{
		Account:  user.Account,
		Email:    user.Email,
		NickName: user.NickName,
	}
	c.JSON(http.StatusOK, result)
}

// ModifyPassword 修改密码处理接口
func (handle *handleManager) ModifyPassword(c *gin.Context) {
	if err := handle.um.ModifyPassword(c); err != nil {
		log.Printf("ModifyPassword Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.ModifyPasswordFail)
		return
	}
	c.JSON(http.StatusOK, consts.ModifyPasswordSuccess)
}

// UserLogin 用户登录处理接口
func (handle *handleManager) UserLogin(c *gin.Context) {
	if err := handle.lm.UserLogin(c); err != nil {
		log.Printf("Login Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.LoginFail)
		return
	}
	c.JSON(http.StatusOK, consts.LoginSuccess)
}

// UserRegedit 用户注册处理接口
func (handle *handleManager) UserRegedit(c *gin.Context) {
	if err := handle.lm.UserRegedit(c); err != nil {
		log.Printf("Regedit Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.RegeditFail)
		return
	}
	c.JSON(http.StatusOK, consts.RegeditSuccess)
}

// UploadFile 用户文件上传处理接口
func (handle *handleManager) UploadFile(c *gin.Context) {
	if err := handle.um.UploadFile(c); err != nil {
		log.Printf("UploadFile Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.UploadFail)
		return
	}
	c.JSON(http.StatusOK, consts.UploadSuccess)
}

// UploadAvatar 用户头像上传处理接口
func (handle *handleManager) UploadAvatar(c *gin.Context) {
	if err := handle.um.UploadAvatar(c); err != nil {
		log.Printf("UploadAvatar Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.UploadFail)
		return
	}
	c.JSON(http.StatusOK, consts.UploadSuccess)
}

// SendAuthCode 发送验证码处理接口
func (handle *handleManager) SendAuthCode(c *gin.Context) {
	if err := handle.lm.SendAuthCode(c); err != nil {
		log.Printf("SendAuthCode Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.SendAuthCodeFail)
	}
	c.JSON(http.StatusOK, consts.SendAuthCodeSuccess)
}

func UploadFilePathCreate() (err error) {
	if err = os.MkdirAll(consts.DefaultUserFilePath, os.ModePerm); err != nil {
		log.Print("目录创建失败", err)
		return err
	}
	if err = os.MkdirAll(consts.DefaultAvatarPath, os.ModePerm); err != nil {
		log.Print("目录创建失败", err)
		return err
	}
	return nil
}

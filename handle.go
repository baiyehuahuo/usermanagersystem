package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
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

// get post 参数在本层校验

// CheckAuthCode 验证码检测处理接口
func (handle *handleManager) CheckAuthCode(c *gin.Context) {
	email := c.Query("email")
	authCode, err := strconv.Atoi(c.Query("auth_code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, consts.InputParamsError)
		return
	}
	if err := handle.lm.CheckAuthCode(c, email, authCode); err != nil {
		log.Printf("CheckAuthCode Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.CheckAuthCodeFail)
		return
	}
	c.JSON(http.StatusOK, consts.CheckAuthCodeSuccess)
}

// CheckEmailAvailable 验证码检测处理接口
func (handle *handleManager) CheckEmailAvailable(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}
	if err := handle.lm.CheckEmailAvaiable(c, email); err != nil {
		log.Printf("CheckEmailAvailable Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.EmailUnavailable)
		return
	}
	c.JSON(http.StatusOK, consts.EmailAvailable)
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
		Account:    user.Account,
		Email:      user.Email,
		NickName:   user.NickName,
		AvatarPath: "", // todo 计算avatar路径
	}
	c.JSON(http.StatusOK, result)
}

// ModifyPassword 修改密码处理接口
func (handle *handleManager) ModifyPassword(c *gin.Context) {
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	if oldPassword == "" || newPassword == "" {
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}

	if err := handle.um.ModifyPassword(c, oldPassword, newPassword); err != nil {
		log.Printf("ModifyPassword Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.ModifyPasswordFail)
		return
	}
	c.JSON(http.StatusOK, consts.ModifyPasswordSuccess)
}

// UserLogin 用户登录处理接口
func (handle *handleManager) UserLogin(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	if account == "" || password == "" {
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}

	if err := handle.lm.UserLogin(c, account, password); err != nil {
		log.Printf("Login Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.LoginFail)
		return
	}
	c.JSON(http.StatusOK, consts.LoginSuccess)
}

// UserRegedit 用户注册处理接口
func (handle *handleManager) UserRegedit(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	email := c.Query("email")
	nickName := c.Query("nick_name")
	if account == "" || password == "" || email == "" || nickName == "" {
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}

	if err := handle.lm.UserRegedit(c, account, password, email, nickName); err != nil {
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
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}

	if err := handle.lm.SendAuthCode(c, email); err != nil {
		log.Printf("SendAuthCode Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.SendAuthCodeFail)
		return
	}
	c.JSON(http.StatusOK, consts.SendAuthCodeSuccess)
}

func UploadFilePathCreate() (err error) {
	if err = os.MkdirAll(consts.DefaultUserFilePath, os.ModePerm); err != nil {
		log.Print("目录创建失败 ", err)
		return err
	}
	if err = os.MkdirAll(consts.DefaultAvatarPath, os.ModePerm); err != nil {
		log.Print("目录创建失败 ", err)
		return err
	}
	if err = os.MkdirAll(consts.DefaultMoviewPath, os.ModePerm); err != nil {
		log.Print("创建目录失败 ", err)
		return err
	}
	return nil
}

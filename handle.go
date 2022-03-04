package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/service/login_control"
	"usermanagersystem/service/user_control"
	"usermanagersystem/utils"

	"github.com/robfig/cron"

	"github.com/gin-gonic/gin"
)

type handleManager struct {
	lm login_control.LoginController
	um user_control.UserController
}

// get post 参数在本层校验

// CheckAuthCode 验证码检测处理接口
func (handle *handleManager) CheckAuthCode(c *gin.Context) {
	email := c.Query("email")
	authCode, err := strconv.Atoi(c.Query("auth_code"))
	if err != nil || !verifyEmailFormat(email) {
		log.Printf("CheckAuthCode fail: %s get auth code fail.", email)
		c.JSON(http.StatusBadRequest, consts.InputParamsError)
		return
	}

	if err := handle.lm.CheckAuthCode(c, email, authCode); err != nil {
		log.Printf("CheckAuthCode fail: %s check auth code fail.", email)
		c.JSON(http.StatusInternalServerError, consts.CheckAuthCodeFail)
		return
	}

	log.Printf("CheckAuthCode success: %s check auth code success.", email)
	c.JSON(http.StatusOK, consts.CheckAuthCodeSuccess)
}

// CheckEmailAvailable 验证码检测处理接口
func (handle *handleManager) CheckEmailAvailable(c *gin.Context) {
	email := c.Query("email")
	if !verifyEmailFormat(email) {
		log.Printf("CheckEmailAvailable fail: email is wrong.")
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}
	if err := handle.lm.CheckEmailAvailable(c, email); err != nil {
		log.Printf("CheckEmailAvailable fail: %s \terr: %v.", email, err)
		c.JSON(http.StatusInternalServerError, consts.EmailUnavailable)
		return
	}
	log.Printf("CheckEmailAvailable success: %s.", email)
	c.JSON(http.StatusOK, consts.EmailAvailable)
}

// ForgetPassword 验证码修改密码
func (handle *handleManager) ForgetPassword(c *gin.Context) {
	email := c.PostForm("email")
	authCode, err := strconv.Atoi(c.PostForm("auth_code"))
	newPassword := c.PostForm("new_password")
	if email == "" || err != nil || newPassword == "" {
		log.Printf("ForgetPassword fail: input error.")
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}
	if err = handle.lm.CheckAuthCode(c, email, authCode); err != nil {
		log.Printf("ForgetPassword fail: check auth code fail: %s \terr: %v.", email, err)
		c.JSON(http.StatusInternalServerError, consts.CheckAuthCodeFail)
		return
	}
	if err = handle.um.SetPassword(c, email, newPassword); err != nil {
		log.Printf("ForgetPassword fail: set password fail: %s \t %s\terr: %v.", email, newPassword, err)
		c.JSON(http.StatusInternalServerError, consts.ForgetPasswordFail)
		return
	}
	log.Printf("ForgetPassword success: %s.", email)
	c.JSON(http.StatusOK, consts.ForgetPasswordSuccess)
}

// GetUserMessageByCookie 通过Cookie获取用户信息处理接口
func (handle *handleManager) GetUserMessageByCookie(c *gin.Context) {
	user, err := handle.um.GetUserMessageByCookie(c)
	if err != nil {
		log.Printf("GetUserMessage fail: %s \terr: %v.", user, err)
		c.JSON(http.StatusInternalServerError, consts.GetUserMessageFail)
		return
	}
	result := model.UserMessage{
		Account:    user.Account,
		Email:      user.Email,
		NickName:   user.NickName,
		AvatarPath: utils.GetNetAvatarPath(user.Account, user.AvatarExt),
	}
	log.Printf("GetUserMessage success: %s.", user)
	c.JSON(http.StatusOK, result)
}

// ModifyPassword 修改密码处理接口
func (handle *handleManager) ModifyPassword(c *gin.Context) {
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	if oldPassword == "" || newPassword == "" {
		log.Printf("ModifyPassword fail: password is empty.")
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}
	user, err := handle.um.GetAccountByCookie(c)
	if err != nil {
		log.Printf("ModifyPassword fail: %s \terr: %v", user, err)
		c.JSON(http.StatusInternalServerError, consts.ModifyPasswordFail)
	}
	if err = handle.um.ModifyPassword(c, user, oldPassword, newPassword); err != nil {
		log.Printf("ModifyPassword fail: %s \terr: %v.", user, err)
		c.JSON(http.StatusInternalServerError, consts.ModifyPasswordFail)
		return
	}
	log.Printf("ModifyPassword success: %s", user)
	c.JSON(http.StatusOK, consts.ModifyPasswordSuccess)
}

// RestoreMySQL 恢复数据库
func (handle *handleManager) RestoreMySQL(c *gin.Context) {
	utils.RestoreMySQL()
}

// UserLogin 用户登录处理接口
func (handle *handleManager) UserLogin(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	if account == "" || password == "" {
		log.Printf("UserLogin fail: account or password is empty.")
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}

	if err := handle.lm.UserLogin(c, account, password); err != nil {
		log.Printf("UserLogin fail: %s \terr: %v.", account, err)
		c.JSON(http.StatusInternalServerError, consts.LoginFail)
		return
	}
	log.Printf("UserLogin success: %s", account)
	c.JSON(http.StatusOK, consts.LoginSuccess)
}

// UserRegister 用户注册处理接口
func (handle *handleManager) UserRegister(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	email := c.Query("email")
	nickName := c.Query("nick_name")
	authCode, err := strconv.Atoi(c.Query("auth_code"))
	if account == "" || password == "" || !verifyEmailFormat(email) || err != nil || nickName == "" {
		log.Printf("UserRegister fail: params has wrong.")
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}

	if err := handle.lm.UserRegister(c, account, password, email, authCode, nickName); err != nil {
		log.Printf("UserRegister fail: %s \terr: %v.", account, err)
		c.JSON(http.StatusInternalServerError, consts.RegeditFail)
		return
	}
	log.Printf("UserRegister success: %s", account)
	c.JSON(http.StatusOK, consts.RegeditSuccess)
}

// UploadFile 用户文件上传处理接口
func (handle *handleManager) UploadFile(c *gin.Context) {
	user, err := handle.um.GetAccountByCookie(c)
	if err != nil {
		log.Printf("UploadFile fail: user is not found.")
		c.JSON(http.StatusInternalServerError, consts.UploadFail)
		return
	}
	if err := handle.um.UploadFile(c); err != nil {
		log.Printf("UploadFile fail: %s \terr: %v.", user, err)
		c.JSON(http.StatusInternalServerError, consts.UploadFail)
		return
	}
	log.Printf("UploadFile success: %s.", user)
	c.JSON(http.StatusOK, consts.UploadSuccess)
}

// UploadAvatar 用户头像上传处理接口
func (handle *handleManager) UploadAvatar(c *gin.Context) {
	user, err := handle.um.GetAccountByCookie(c)
	if err != nil {
		log.Printf("UploadAvatar fail: user is not found.")
		c.JSON(http.StatusInternalServerError, consts.UploadFail)
		return
	}
	if err := handle.um.UploadAvatar(c); err != nil {
		log.Printf("UploadAvatar fail: %s \terr: %v.", user, err)
		c.JSON(http.StatusInternalServerError, consts.UploadFail)
		return
	}
	log.Printf("UploadAvatar success: %s.", user)

	c.JSON(http.StatusOK, consts.UploadSuccess)
}

// SendAuthCode 发送验证码处理接口
func (handle *handleManager) SendAuthCode(c *gin.Context) {
	email := c.Query("email")
	if !verifyEmailFormat(email) {
		log.Printf("SendAuthCode fail: email is wrong.")
		c.JSON(http.StatusInternalServerError, consts.InputParamsError)
		return
	}

	if err := handle.lm.SendAuthCode(c, email); err != nil {
		log.Printf("SendAuthCode fail: %s \terr: %v", email, err)
		c.JSON(http.StatusInternalServerError, consts.SendAuthCodeFail)
		return
	}
	log.Printf("SendAuthCode success: %s", email)
	c.JSON(http.StatusOK, consts.SendAuthCodeSuccess)
}

// UploadFilePathCreate 创建文件上传路径
func UploadFilePathCreate() (err error) {
	if err = os.MkdirAll(consts.DefaultUserFilePath, os.ModePerm); err != nil {
		log.Print("目录创建失败 ", err)
		return err
	}
	if err = os.MkdirAll(consts.DefaultAvatarPath, os.ModePerm); err != nil {
		log.Print("目录创建失败 ", err)
		return err
	}
	if err = os.MkdirAll(consts.SystemLogPath, os.ModePerm); err != nil {
		log.Print("目录创建失败 ", err)
		return err
	}
	// if err = os.MkdirAll(consts.DefaultMoviewPath, os.ModePerm); err != nil {
	// 	log.Print("创建目录失败 ", err)
	// 	return err
	// }
	return nil
}

// SetLog 设置日志路径
func SetLog() (err error) {
	logFile, err := os.OpenFile(consts.LogFilePath, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return nil
}

// SetTimer 启动定时器
func SetTimer() (err error) {
	c := cron.New() // 精确到秒

	utils.BackupMySQL()
	spec := "0 */5 * * * ?" // every 5 minutes
	if err := c.AddFunc(spec, utils.BackupMySQL); err != nil {
		return err
	}

	c.Start()
	return nil
}

// Cors 设置跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		r := c.Request
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET")
		// w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Allow-Headers", "Access-Token")
		// w.Header().Set("content-type", "application/javascript")
		c.Next()
	}
}

// verifyEmailFormat 验证邮箱格式
func verifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

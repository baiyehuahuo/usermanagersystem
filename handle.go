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

// CheckEmailAvailable 验证码检测处理接口
func (handle *handleManager) CheckEmailAvailable(c *gin.Context) {
	email := c.Query("email")
	if !verifyEmailFormat(email) {
		log.Printf("CheckEmailAvailable fail: email is wrong.")
		returnFail(c, model.Err{Code: consts.InputParamsWrong})
		return
	}
	if Err := handle.lm.CheckEmailAvailable(c, email); Err.Code != consts.OperateSuccess {
		log.Printf("CheckEmailAvailable fail: %s \terr: %v.", email, Err.Msg)
		returnFail(c, Err)
		return
	}
	log.Printf("CheckEmailAvailable success: %s.", email)
	returnSuccess(c)
}

// DeletePng 验证码检测处理接口
func (handle *handleManager) DeletePng(c *gin.Context) {
	account, Err := handle.um.GetAccountByCookie(c)
	if account == "" || Err.Code != consts.OperateSuccess {
		log.Printf("DeletePng fail: user is not found. %v", Err.Msg)
		returnFail(c, Err)
		return
	}
	png := c.PostForm("delete_png_name")
	if png == "" {
		log.Printf("DeletePng fail: png is nil.")
		returnFail(c, model.Err{Code: consts.InputParamsWrong})
		return
	}

	if Err := handle.um.DeletePng(c, account, png); Err.Code != consts.OperateSuccess {
		log.Printf("DeletePng fail: %s \terr: %v.", account, Err.Msg)
		returnFail(c, Err)
		return
	}

	log.Printf("DeletePng success: %s.", account)
	returnSuccess(c)
}

// ForgetPassword 验证码修改密码
func (handle *handleManager) ForgetPassword(c *gin.Context) {
	email := c.PostForm("email")
	authCode, err := strconv.Atoi(c.PostForm("auth_code"))
	newPassword := c.PostForm("new_password")
	if email == "" || err != nil || newPassword == "" {
		log.Printf("ForgetPassword fail: input error.")
		returnFail(c, model.Err{Code: consts.InputParamsWrong})
		return
	}
	if Err := handle.lm.CheckAuthCode(c, email, authCode); Err.Code != consts.OperateSuccess {
		log.Printf("ForgetPassword fail: check auth code fail: %s \terr: %v.", email, Err.Msg)
		returnFail(c, Err)
		return
	}
	if Err := handle.um.SetPassword(c, email, newPassword); Err.Code != consts.OperateSuccess {
		log.Printf("ForgetPassword fail: set password fail: %s \t %s\terr: %v.", email, newPassword, Err.Msg)
		returnFail(c, Err)
		return
	}
	log.Printf("ForgetPassword success: %s.", email)
	returnSuccess(c)
}

func (handle *handleManager) GetUserFilesPath(c *gin.Context) {
	account, Err := handle.um.GetAccountByCookie(c)
	if account == "" || Err.Code != consts.OperateSuccess {
		log.Printf("GetUserFilesPath fail: user is not found. %v", Err.Msg)
		returnFail(c, Err)
		return
	}
	var result []string
	if result, Err = handle.um.GetUserFilesPath(c, account); Err.Code != consts.OperateSuccess {
		log.Printf("GetUserFilesPath fail: %s \terr: %v.", account, Err.Msg)
		returnFail(c, Err)
		return
	}
	log.Printf("GetUserFilesPath success: %s.", account)

	c.JSON(http.StatusOK, model.Err{
		Code: consts.OperateSuccess,
		Msg:  consts.ErrCodeMessage[consts.OperateSuccess],
		Data: result,
	})
}

// GetUserMessageByCookie 通过Cookie获取用户信息处理接口
func (handle *handleManager) GetUserMessageByCookie(c *gin.Context) {
	user, Err := handle.um.GetUserMessageByCookie(c)
	if Err.Code != consts.OperateSuccess {
		log.Printf("GetUserMessage fail: %s \terr: %v.", user, Err.Msg)
		returnFail(c, Err)
		return
	}
	result := model.UserMessage{
		Account:    user.Account,
		Email:      user.Email,
		NickName:   user.NickName,
		AvatarPath: utils.GetNetAvatarPath(user.Account, user.AvatarExt),
	}
	log.Printf("GetUserMessage success: %s.", user)
	c.JSON(http.StatusOK, model.Err{
		Code: consts.OperateSuccess,
		Msg:  consts.ErrCodeMessage[consts.OperateSuccess],
		Data: result,
	})
}

// ModifyPassword 修改密码处理接口
func (handle *handleManager) ModifyPassword(c *gin.Context) {
	oldPassword := c.PostForm("oldPassword")
	newPassword := c.PostForm("newPassword")
	if oldPassword == "" || newPassword == "" {
		log.Printf("ModifyPassword fail: password is empty.")
		returnFail(c, model.Err{Code: consts.InputParamsWrong})
		return
	}
	user, Err := handle.um.GetAccountByCookie(c)
	if Err.Code != consts.OperateSuccess {
		log.Printf("ModifyPassword fail: %s \terr: %v", user, Err)
		returnFail(c, Err)
		return
	}
	if Err = handle.um.ModifyPassword(c, user, oldPassword, newPassword); Err.Code != consts.OperateSuccess {
		log.Printf("ModifyPassword fail: %s \terr: %v.", user, Err)
		returnFail(c, Err)
		return
	}
	log.Printf("ModifyPassword success: %s", user)
	returnSuccess(c)
}

// PredictPng 分割Png
func (handle *handleManager) PredictPng(c *gin.Context) {
	predictPngName := c.Query("predict_png_name")
	if predictPngName == "" {
		log.Printf("PredictPng fail: params has wrong.")
		returnFail(c, model.Err{Code: consts.InputParamsWrong})
		return
	}

	account, Err := handle.um.GetAccountByCookie(c)
	if account == "" || Err.Code != consts.OperateSuccess {
		log.Printf("Predict fail: user is not found.")
		returnFail(c, Err)
		return
	}

	var predictPath string
	if predictPath, Err = handle.um.PredictPng(c, account, predictPngName); Err.Code != consts.OperateSuccess {
		log.Printf("Predict fail: %v", Err.Msg)
		returnFail(c, Err)
		return
	}
	log.Printf("Predict success: %s\t%s", account, predictPngName)
	c.JSON(http.StatusOK, model.Err{
		Code: consts.OperateSuccess,
		Msg:  consts.ErrCodeMessage[consts.OperateSuccess],
		Data: predictPath,
	})
}

// RestoreMySQL 恢复数据库
func (handle *handleManager) RestoreMySQL(c *gin.Context) {
	if Err := utils.RestoreMySQL(); Err.Code != consts.OperateSuccess {
		returnFail(c, Err)
	}
	returnSuccess(c)
}

// UserLogin 用户登录处理接口
func (handle *handleManager) UserLogin(c *gin.Context) {
	account := c.Query("account")
	password := c.Query("password")
	if account == "" || password == "" {
		log.Printf("UserLogin fail: account or password is empty.")
		returnFail(c, model.Err{Code: consts.InputParamsWrong})
		return
	}

	if Err := handle.lm.UserLogin(c, account, password); Err.Code != consts.OperateSuccess {
		log.Printf("UserLogin fail: %s \terr: %v.", account, Err)
		returnFail(c, Err)
		return
	}
	log.Printf("UserLogin success: %s", account)
	returnSuccess(c)
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
		returnFail(c, model.Err{Code: consts.InputParamsWrong})
		return
	}

	if Err := handle.lm.UserRegister(c, account, password, email, authCode, nickName); Err.Code != consts.OperateSuccess {
		log.Printf("UserRegister fail: %s \terr: %v.", account, err)
		returnFail(c, Err)
		return
	}
	log.Printf("UserRegister success: %s", account)
	returnSuccess(c)
}

// UploadFile 用户文件上传处理接口
func (handle *handleManager) UploadPng(c *gin.Context) {
	account, Err := handle.um.GetAccountByCookie(c)
	if account == "" || Err.Code != consts.OperateSuccess {
		log.Printf("UploadFile fail: user is not found.")
		returnFail(c, Err)
		return
	}
	if Err = handle.um.UploadPng(c, account); Err.Code != consts.OperateSuccess {
		log.Printf("UploadFile fail: %s \terr: %v.", account, Err.Msg)
		returnFail(c, Err)
		return
	}
	log.Printf("UploadFile success: %s.", account)
	returnSuccess(c)
}

// UploadAvatar 用户头像上传处理接口
func (handle *handleManager) UploadAvatar(c *gin.Context) {
	user, Err := handle.um.GetAccountByCookie(c)
	if Err.Code != consts.OperateSuccess {
		log.Printf("UploadAvatar fail: user is not found.")
		returnFail(c, Err)
		return
	}
	if Err = handle.um.UploadAvatar(c); Err.Code != consts.OperateSuccess {
		log.Printf("UploadAvatar fail: %s \terr: %v.", user, Err.Msg)
		returnFail(c, Err)
		return
	}
	log.Printf("UploadAvatar success: %s.", user)
	returnSuccess(c)
}

// SendAuthCode 发送验证码处理接口
func (handle *handleManager) SendAuthCode(c *gin.Context) {
	email := c.Query("email")
	if !verifyEmailFormat(email) {
		log.Printf("SendAuthCode fail: email is wrong.")
		returnFail(c, model.Err{Code: consts.InputParamsWrong})
		return
	}

	if Err := handle.lm.SendAuthCode(c, email); Err.Code != consts.OperateSuccess {
		log.Printf("SendAuthCode fail: %s \terr: %v", email, Err.Msg)
		returnFail(c, Err)
		return
	}
	log.Printf("SendAuthCode success: %s", email)
	returnSuccess(c)
}

// UploadFilePathCreate 创建文件上传路径
func UploadFilePathCreate() (err error) {
	if err = os.MkdirAll(consts.DefaultUserPngPath, os.ModePerm); err != nil {
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

func returnSuccess(c *gin.Context) {
	// requestFuncName := c.Request.RequestURI[1:]
	// if index := strings.Index(requestFuncName, "?"); index != -1 {
	// 	requestFuncName = requestFuncName[:index]
	// }
	c.JSON(http.StatusOK, model.Err{
		Code: consts.OperateSuccess,
		Msg:  consts.ErrCodeMessage[consts.OperateSuccess],
	})
}

func returnFail(c *gin.Context, Err model.Err) {
	Err.Msg = consts.ErrCodeMessage[Err.Code]
	c.JSON(http.StatusOK, Err)
}

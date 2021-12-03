package logincontrol

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/utils"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type loginControllerImpl struct {
	rc utils.RedisController
}

func (loginController *loginControllerImpl) CheckAuthCode(c *gin.Context) (err error) {
	email := c.Query("email")
	authCode, err := strconv.Atoi(c.Query("auth_code"))
	if err != nil {
		return err
	}
	if !utils.GetEACC().CheckAuthCodeByEmail(email, authCode) {
		return errors.New(consts.CheckAuthCodeFail)
	}
	return nil
}

func (loginController *loginControllerImpl) CheckEmailAvaiable(c *gin.Context) error {
	email := c.Query("email")
	if email == "" {
		return errors.New(consts.InputParamsError)
	}

	user := model.User{Email: email}
	if err := utils.GetDB().Where(&user).Take(&user).Error; err != gorm.ErrRecordNotFound {
		return errors.New(consts.InputParamsError)
	}
	return nil
}

func (loginController *loginControllerImpl) UserLogin(c *gin.Context) (err error) {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
	}
	if err = utils.GetDB().Where(&user).Take(&user).Error; err == gorm.ErrRecordNotFound {
		return err
	}

	c.SetSameSite(http.SameSiteLaxMode)
	cookie := fmt.Sprintf("%x", md5.Sum([]byte(user.Account+time.Now().String()))) // cookie值
	c.SetCookie(consts.CookieNameOfUser, cookie, consts.CookieContinueTime, consts.CookieValidationRange,
		consts.CookieValidationDomain, false, true)
	if err = loginController.rc.Set(consts.RedisCookieHashPrefix+cookie, user.Account,
		consts.CookieContinueTime); err != nil {
		return err
	}

	if err = loginController.rc.SetUser(user); err != nil { // 保存到 redis 缓存中
		log.Printf("user %s save into redis fail: %v", user.Account, err)
	}

	return nil
}

func (loginController *loginControllerImpl) UserRegedit(c *gin.Context) (err error) {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
		Email:    c.Query("email"), // todo 检测是否已被注册
		NickName: c.Query("nick_name"),
	}
	// todo 邮箱验证码
	if err = utils.GetDB().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (loginController *loginControllerImpl) SendAuthCode(c *gin.Context) (err error) {
	email := c.Query("email")
	if err = utils.GetEACC().SendAuthCodeByEmail(email); err != nil {
		return err
	}
	return nil
}

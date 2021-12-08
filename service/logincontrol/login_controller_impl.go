package logincontrol

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"time"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/utils"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type loginControllerImpl struct {
	rc utils.RedisController
}

func (loginController *loginControllerImpl) CheckAuthCode(c *gin.Context, email string, authCode int) (err error) {
	if !utils.GetEACC().CheckAuthCodeByEmail(email, authCode) {
		return errors.Wrap(errors.New(consts.CheckAuthCodeFail), utils.RunFuncNameWithFail())
	}
	return nil
}

func (loginController *loginControllerImpl) CheckEmailAvailable(c *gin.Context, email string) (err error) {
	user := model.User{Email: email}
	if err := utils.GetDB().Where(&user).Take(&user).Error; err != gorm.ErrRecordNotFound {
		return errors.Wrap(errors.New(consts.EmailUnavailable), utils.RunFuncNameWithFail())
	}
	return nil
}

func (loginController *loginControllerImpl) UserLogin(c *gin.Context, account string, password string) (err error) {
	user := model.User{
		Account:  account,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(password))),
	}
	if err = utils.GetDB().Where(&user).Take(&user).Error; err == gorm.ErrRecordNotFound {
		return errors.Wrap(err, utils.RunFuncNameWithFail())
	}

	c.SetSameSite(http.SameSiteLaxMode)
	cookie := fmt.Sprintf("%x", md5.Sum([]byte(user.Account+time.Now().String()))) // cookie值
	c.SetCookie(consts.CookieNameOfUser, cookie, consts.CookieContinueTime, consts.CookieValidationRange,
		consts.CookieValidationDomain, false, true)
	if err = loginController.rc.Set(consts.RedisCookieHashPrefix+cookie, user.Account,
		consts.CookieContinueTime); err != nil {
		return errors.WithMessage(err, utils.RunFuncNameWithFail())
	}

	if err = loginController.rc.SetUser(user); err != nil { // 保存到 redis 缓存中 失败也不必停止
		log.Printf("user %s save into redis fail: %v", user.Account, err)
	}

	return nil
}

func (loginController *loginControllerImpl) UserRegister(c *gin.Context, account string, password string, email string, nickName string) (err error) {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
		Email:    c.Query("email"), // todo 检测是否已被注册
		NickName: c.Query("nick_name"),
	}
	// todo 邮箱验证码
	if err = utils.GetDB().Create(&user).Error; err != nil {
		return errors.Wrap(err, utils.RunFuncNameWithFail())
	}
	return nil
}

func (loginController *loginControllerImpl) SendAuthCode(c *gin.Context, email string) (err error) {
	if err = utils.GetEACC().SendAuthCodeByEmail(email); err != nil {
		return errors.WithMessage(err, utils.RunFuncNameWithFail())
	}
	return nil
}

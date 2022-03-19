package login_control

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
		return utils.ErrWrapOrWithMessage(true, errors.New(""))
	}
	return nil
}

func (loginController *loginControllerImpl) CheckEmailAvailable(c *gin.Context, email string) (err error) {
	user := model.User{Email: email}
	if err := utils.GetDB().Where(&user).Take(&user).Error; err != gorm.ErrRecordNotFound {
		return utils.ErrWrapOrWithMessage(true, errors.New(consts.EmailUnavailable))
	}
	return nil
}

func (loginController *loginControllerImpl) UserLogin(c *gin.Context, account string, password string) (Err model.Err) {
	var err error
	user := model.User{
		Account:  account,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(password))),
	}
	if err = utils.GetDB().Where(&user).Take(&user).Error; err == gorm.ErrRecordNotFound {
		return model.Err{
			Code: consts.AccountOrPasswordWrong,
			Msg:  utils.ErrWrapOrWithMessage(true, err).Error(),
		}
	}

	c.SetSameSite(http.SameSiteLaxMode)
	cookie := fmt.Sprintf("%x", md5.Sum([]byte(user.Account+time.Now().String()))) // cookie值
	c.SetCookie(consts.CookieNameOfUser, cookie, consts.CookieContinueTime, consts.CookieValidationRange,
		consts.CookieValidationDomainIP, false, true)
	c.SetCookie(consts.CookieNameOfUser, cookie, consts.CookieContinueTime, consts.CookieValidationRange,
		consts.CookieValidationDomainLocal, false, true)
	if err = loginController.rc.Set(consts.RedisCookieHashPrefix+cookie, user.Account, consts.CookieContinueTime); err != nil {
		return model.Err{
			Code: consts.DatabaseWrong,
			Msg:  utils.ErrWrapOrWithMessage(false, err).Error(),
		}
	}

	if err := loginController.rc.SetUser(user); err != nil { // 保存到 redis 缓存中 失败也不必停止
		log.Printf("user %s save into redis fail: %v", user.Account, err)
	}

	Err.Code = consts.OperateSuccess
	return Err
}

func (loginController *loginControllerImpl) UserRegister(c *gin.Context, account string, password string, email string, authCode int, nickName string) (Err model.Err) {
	user := model.User{
		Account:  account,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(password))),
		Email:    email,
		NickName: nickName,
	}
	var err error
	if err = loginController.CheckAuthCode(c, email, authCode); err != nil {
		Err.Code = consts.CheckAuthCodeFail
		Err.Msg = utils.ErrWrapOrWithMessage(false, err).Error()
		return Err
	}
	if err = utils.GetDB().Create(&user).Error; err != nil {
		Err.Code = consts.DatabaseWrong
		Err.Msg = utils.ErrWrapOrWithMessage(true, err).Error()
		return Err
	}
	Err.Code = consts.OperateSuccess
	return Err
}

func (loginController *loginControllerImpl) SendAuthCode(c *gin.Context, email string) (Err model.Err) {
	var err error
	if err = utils.GetEACC().SendAuthCodeByEmail(email); err != nil {
		Err.Code = consts.SendAuthCodeByEmailFail
		Err.Msg = utils.ErrWrapOrWithMessage(false, err).Error()
		return Err
	}
	Err.Code = consts.OperateSuccess
	return Err
}

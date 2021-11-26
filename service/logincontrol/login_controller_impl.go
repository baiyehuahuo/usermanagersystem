package logincontrol

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/utils/databasecontrol"
	"usermanagersystem/utils/rediscontrol"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type loginControllerImpl struct {
	rc rediscontrol.RedisController
}

func (loginController *loginControllerImpl) UserLogin(c *gin.Context) error {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
	}
	if err := databasecontrol.GetDB().Where(&user).Take(&user).Error; err == gorm.ErrRecordNotFound {
		return err
	}

	c.SetSameSite(http.SameSiteLaxMode)
	cookie := fmt.Sprintf("%x", md5.Sum([]byte(user.Account+time.Now().String()))) // cookie值
	c.SetCookie(consts.UserCookieName, cookie, consts.CookieContinueTime, consts.CookieValidationRange,
		consts.CookieValidationDomain, false, true)
	if err := loginController.rc.Set(consts.RedisCookieHashPrefix+cookie, user.Account,
		time.Second*consts.CookieContinueTime); err != nil {
		return err
	}

	return nil
}

func (loginController *loginControllerImpl) UserRegedit(c *gin.Context) error {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
		Email:    c.Query("email"), // todo 检测是否已被注册
		NickName: c.Query("nickname"),
	}
	if err := databasecontrol.GetDB().Create(&user).Error; err != nil {
		return err
	}
	return nil
}

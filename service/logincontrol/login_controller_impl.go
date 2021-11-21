package logincontrol

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/utils/databasecontrol"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type loginControllerImpl struct {
}

func (loginController *loginControllerImpl) UserLogin(c *gin.Context) error {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
	}
	if err := databasecontrol.DB.Where(&user).Take(&user).Error; err == gorm.ErrRecordNotFound {
		return err
	}

	c.SetSameSite(http.SameSiteLaxMode)
	cookie := fmt.Sprintf("%x", md5.Sum([]byte(user.Account+time.Now().String()))) // cookie值
	c.SetCookie(consts.UserCookieName, cookie, consts.CookieContinueTime, consts.CookieValidationRange, consts.CookieValidationDomain, false, true)

	return nil
}

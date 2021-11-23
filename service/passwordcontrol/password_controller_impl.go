package passwordcontrol

import (
	"crypto/md5"
	"errors"
	"fmt"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/utils/databasecontrol"
	"usermanagersystem/utils/rediscontrol"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type passwordControllerImpl struct {
	rc rediscontrol.RedisController
}

func (p *passwordControllerImpl) ModifyPassword(c *gin.Context) error {
	var username string
	if cookie, err := c.Cookie(consts.UserCookieName); err == nil {
		username, _ = p.rc.Get(consts.RedisCookieHash + cookie)
	}
	if username == "" {
		return errors.New("无效cookie")
	}

	user := model.User{
		Account:  username,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.PostForm("oldPassword")))),
	}
	if err := databasecontrol.GetDB().Where(&user).Take(&user).Error; err == gorm.ErrRecordNotFound {
		return err
	}
	modify := model.User{
		Account:  username,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.PostForm("newPassword")))),
	}
	if err := databasecontrol.GetDB().Model(&user).Updates(modify).Error; err != nil {
		return err
	}
	return nil
}

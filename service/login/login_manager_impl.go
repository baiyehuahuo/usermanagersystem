package login

import (
	"crypto/md5"
	"fmt"
	"usermanagersystem/model"
	"usermanagersystem/utils/database"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type loginManagerImpl struct {
}

func (login *loginManagerImpl) UserLogin(c *gin.Context) error {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
	}

	if err := database.DB.Where(&user).Take(&user).Error; err == gorm.ErrRecordNotFound {
		return err
	}

	return nil
}

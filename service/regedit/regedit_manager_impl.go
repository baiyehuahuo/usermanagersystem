package regedit

import (
	"crypto/md5"
	"fmt"
	"usermanagersystem/model"
	"usermanagersystem/utils/database"

	"github.com/gin-gonic/gin"
)

type regeditManagerImpl struct {
}

func (regeditManager *regeditManagerImpl) UserRegedit(c *gin.Context) error {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

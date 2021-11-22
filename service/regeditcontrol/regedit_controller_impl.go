package regeditcontrol

import (
	"crypto/md5"
	"fmt"
	"usermanagersystem/model"
	"usermanagersystem/utils/databasecontrol"

	"github.com/gin-gonic/gin"
)

type regeditControllerImpl struct {
}

func (regeditController *regeditControllerImpl) UserRegedit(c *gin.Context) error {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
	}
	if err := databasecontrol.GetDB().Create(&user).Error; err != nil {
		return err
	}
	return nil
}
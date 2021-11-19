package login

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/utils/database"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type loginManagerImpl struct {
}

func (login *loginManagerImpl) UserLogin(c *gin.Context) {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
	}

	if err := database.DB.Where(&user).Take(&user).Error; err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, consts.LoginFail)
		return
	}

	c.JSON(http.StatusOK, consts.LoginSuccess)
}

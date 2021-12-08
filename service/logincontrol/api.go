package logincontrol

import (
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	UserLogin(c *gin.Context, account string, password string) error
	UserRegister(c *gin.Context, account string, password string, email string, nickName string) error
	SendAuthCode(c *gin.Context, email string) error
	CheckAuthCode(c *gin.Context, email string, authCode int) error
	CheckEmailAvailable(c *gin.Context, email string) error
}

func New() LoginController {
	return &loginControllerImpl{
		rc: utils.RedisNew(),
	}
}

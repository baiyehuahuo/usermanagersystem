package logincontrol

import (
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	UserLogin(c *gin.Context) error
	UserRegedit(c *gin.Context) error
	SendAuthCode(c *gin.Context) error
	CheckAuthCode(c *gin.Context) error
	CheckEmailAvaiable(c *gin.Context) error
}

func New() LoginController {
	return &loginControllerImpl{
		rc: utils.RedisNew(),
	}
}

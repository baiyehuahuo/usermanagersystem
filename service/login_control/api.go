package login_control

import (
	"usermanagersystem/model"
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	UserLogin(c *gin.Context, account string, password string) (err model.Err)
	UserRegister(c *gin.Context, account string, password string, email string, authCode int, nickName string) (err model.Err)
	SendAuthCode(c *gin.Context, email string) (err model.Err)
	CheckAuthCode(c *gin.Context, email string, authCode int) (err model.Err)
	CheckEmailAvailable(c *gin.Context, email string) (err model.Err)
}

func New() LoginController {
	return &loginControllerImpl{
		rc: utils.RedisNew(),
	}
}

package logincontrol

import (
	"usermanagersystem/utils/rediscontrol"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	UserLogin(c *gin.Context) error
	UserRegedit(c *gin.Context) error
}

func New() LoginController {
	return &loginControllerImpl{
		rc: rediscontrol.New(),
	}
}

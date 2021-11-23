package passwordcontrol

import (
	"usermanagersystem/utils/rediscontrol"

	"github.com/gin-gonic/gin"
)

type PasswordController interface {
	ModifyPassword(c *gin.Context) error
}

func New() PasswordController {
	return &passwordControllerImpl{
		rc: rediscontrol.New(),
	}
}

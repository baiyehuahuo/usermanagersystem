package usercontrol

import (
	"usermanagersystem/utils/rediscontrol"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	ModifyPassword(c *gin.Context) error
	FileUpload(c *gin.Context) error
}

func New() UserController {
	return &userControllerImpl{
		rc: rediscontrol.New(),
	}
}

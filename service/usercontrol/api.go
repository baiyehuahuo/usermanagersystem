package usercontrol

import (
	"usermanagersystem/model"
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	getAccountByCookie(c *gin.Context) (string, error)
	getUserByAccount(account string) (*model.User, error)
	GetUserMessageByCookie(c *gin.Context) (*model.User, error)
	ModifyPassword(c *gin.Context) error
	UploadFile(c *gin.Context) error
	UploadAvatar(c *gin.Context) error
}

func New() UserController {
	return &userControllerImpl{
		db: utils.GetDB(),
		rc: utils.RedisNew(),
	}
}

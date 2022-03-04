package user_control

import (
	"usermanagersystem/model"
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	SetPassword(c *gin.Context, email string, password string) error
	getUserByAccount(account string) (user *model.User, err error)
	GetAccountByCookie(c *gin.Context) (account string, err error)
	GetUserMessageByCookie(c *gin.Context) (user *model.User, err error)
	ModifyPassword(c *gin.Context, account string, oldPassword string, newPassword string) (err error)
	UploadFile(c *gin.Context) (err error)
	UploadAvatar(c *gin.Context) (err error)
}

func New() UserController {
	return &userControllerImpl{
		db: utils.GetDB(),
		rc: utils.RedisNew(),
	}
}

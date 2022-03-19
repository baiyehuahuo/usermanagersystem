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
	GetUserFilesPath(c *gin.Context, account string) ([]string, error)
	GetUserMessageByCookie(c *gin.Context) (user *model.User, err error)
	ModifyPassword(c *gin.Context, account string, oldPassword string, newPassword string) (err error)
	UploadPng(c *gin.Context, account string) (err error)
	UploadAvatar(c *gin.Context) (err error)
	PredictPng(c *gin.Context, account string, pngName string) (predictPath string, err error)
	DeletePng(c *gin.Context, account string, pngName string) (err error)
}

func New() UserController {
	return &userControllerImpl{
		db: utils.GetDB(),
		rc: utils.RedisNew(),
	}
}

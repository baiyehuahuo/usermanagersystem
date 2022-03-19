package user_control

import (
	"usermanagersystem/model"
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	SetPassword(c *gin.Context, email string, password string) (Err model.Err)
	getUserByAccount(account string) (user *model.User, err error)
	GetAccountByCookie(c *gin.Context) (account string, Err model.Err)
	GetUserFilesPath(c *gin.Context, account string) ([]string, model.Err)
	GetUserMessageByCookie(c *gin.Context) (user *model.User, Err model.Err)
	ModifyPassword(c *gin.Context, account string, oldPassword string, newPassword string) (Err model.Err)
	UploadPng(c *gin.Context, account string) (Err model.Err)
	UploadAvatar(c *gin.Context) (Err model.Err)
	PredictPng(c *gin.Context, account string, pngName string) (predictPath string, err error)
	DeletePng(c *gin.Context, account string, pngName string) (err error)
}

func New() UserController {
	return &userControllerImpl{
		db: utils.GetDB(),
		rc: utils.RedisNew(),
	}
}

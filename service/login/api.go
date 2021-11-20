package login

import "github.com/gin-gonic/gin"

type LoginManager interface {
	UserLogin(c *gin.Context) error
}

func New() LoginManager {
	return &loginManagerImpl{}
}

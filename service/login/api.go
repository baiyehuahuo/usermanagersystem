package login

import "github.com/gin-gonic/gin"

type LoginManager interface {
	UserLogin(c *gin.Context)
}

func New() LoginManager {
	return &loginManagerImpl{}
}

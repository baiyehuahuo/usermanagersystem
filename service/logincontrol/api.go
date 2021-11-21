package logincontrol

import "github.com/gin-gonic/gin"

type LoginController interface {
	UserLogin(c *gin.Context) error
}

func New() LoginController {
	return &loginControllerImpl{}
}

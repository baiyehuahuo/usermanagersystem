package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginManagerImpl struct {
}

func (login *loginManagerImpl) UserLogin(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello World")
}

package htmlcontrol

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type htmlControllerImpl struct {
}

func (htmlController *htmlControllerImpl) ToLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func (htmlController *htmlControllerImpl) ToUserManage(c *gin.Context) {
	c.HTML(http.StatusOK, "UserManage.html", "")
}

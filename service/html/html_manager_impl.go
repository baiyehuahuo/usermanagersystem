package html

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type htmlManagerImpl struct {
}

func (manager *htmlManagerImpl) ToLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func (manager *htmlManagerImpl) ToUserManage(c *gin.Context) {
	c.HTML(http.StatusOK, "UserManage.html", "")
}

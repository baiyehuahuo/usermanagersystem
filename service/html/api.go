package html

import "github.com/gin-gonic/gin"

type HtmlManager interface {
	ToLogin(c *gin.Context)
	ToUserManage(c *gin.Context)
}

func New() HtmlManager {
	return &htmlManagerImpl{}
}

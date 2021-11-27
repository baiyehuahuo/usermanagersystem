package htmlcontrol

import (
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

type HtmlController interface {
	ToLogin(c *gin.Context)
	ToUserManage(c *gin.Context)
}

func New() HtmlController {
	return &htmlControllerImpl{
		rc: utils.RedisNew(),
	}
}

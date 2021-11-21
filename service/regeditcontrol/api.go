package regeditcontrol

import "github.com/gin-gonic/gin"

type RegeditController interface {
	UserRegedit(c *gin.Context) error
}

func New() RegeditController {
	return &regeditControllerImpl{}
}

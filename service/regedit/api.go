package regedit

import "github.com/gin-gonic/gin"

type RegeditManager interface {
	UserRegedit(c *gin.Context)
}

func New() RegeditManager {
	return &regeditManagerImpl{}
}

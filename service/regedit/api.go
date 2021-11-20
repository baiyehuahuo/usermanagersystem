package regedit

import "github.com/gin-gonic/gin"

type RegeditManager interface {
	UserRegedit(c *gin.Context) error
}

func New() RegeditManager {
	return &regeditManagerImpl{}
}

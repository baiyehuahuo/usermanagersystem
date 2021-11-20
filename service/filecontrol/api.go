package filecontrol

import "github.com/gin-gonic/gin"

type FileControlManager interface {
	FileUpload(c *gin.Context) error
}

func New() FileControlManager {
	return &fileControlManagerImpl{}
}

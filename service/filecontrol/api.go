package filecontrol

import "github.com/gin-gonic/gin"

type FileController interface {
	FileUpload(c *gin.Context) error
}

func New() FileController {
	return &fileControllerImpl{}
}

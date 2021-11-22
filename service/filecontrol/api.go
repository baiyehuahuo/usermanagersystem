package filecontrol

import (
	"usermanagersystem/utils/rediscontrol"

	"github.com/gin-gonic/gin"
)

type FileController interface {
	FileUpload(c *gin.Context) error
}

func New() FileController {
	return &fileControllerImpl{
		rc: rediscontrol.New(),
	}
}

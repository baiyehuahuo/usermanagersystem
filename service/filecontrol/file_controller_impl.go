package filecontrol

import (
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"usermanagersystem/consts"

	"github.com/gin-gonic/gin"
)

type fileControllerImpl struct {
}

func (fileController *fileControllerImpl) FileUpload(c *gin.Context) error {
	var file *multipart.FileHeader
	var err error
	userName := c.PostForm("name") // todo 判断用户名是否可用
	if file, err = c.FormFile("file"); err != nil || userName == "" {
		return err
	}

	var filePath string
	if filePath, err = mkdir(userName); err != nil {
		return err
	}
	filePath = filepath.Join(filePath, file.Filename)

	if err = c.SaveUploadedFile(file, filePath); err != nil {
		return err
	}

	cookie, _ := c.Cookie(consts.UserCookieName)
	log.Print("cookie:", cookie)

	return nil
}

func mkdir(userName string) (string, error) {
	filePath := filepath.Join(consts.DefaultFileRootPath, userName, time.Now().Format("2006/01/02"))
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Print("目录创建失败", err)
		return "", err
	}
	return filePath, nil
}

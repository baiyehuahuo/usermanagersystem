package usercontrol

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"time"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/utils/databasecontrol"
	"usermanagersystem/utils/rediscontrol"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userControllerImpl struct {
	rc rediscontrol.RedisController
}

func (uc *userControllerImpl) ModifyPassword(c *gin.Context) error {
	var username string
	if cookie, err := c.Cookie(consts.UserCookieName); err == nil {
		username, _ = uc.rc.Get(consts.RedisCookieHashPrefix + cookie)
	}
	if username == "" {
		return errors.New("无效cookie")
	}

	user := model.User{
		Account:  username,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.PostForm("oldPassword")))),
	}
	if err := databasecontrol.GetDB().Where(&user).Take(&user).Error; err == gorm.ErrRecordNotFound {
		return err
	}
	modify := model.User{
		Account:  username,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.PostForm("newPassword")))),
	}
	if err := databasecontrol.GetDB().Model(&user).Updates(modify).Error; err != nil {
		return err
	}
	return nil
}

func (uc *userControllerImpl) FileUpload(c *gin.Context) error {
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

func (uc *userControllerImpl) AvatarUpload(c *gin.Context) error {
	var file *multipart.FileHeader
	var err error
	userName := c.PostForm("name") // todo 判断用户名是否可用
	if file, err = c.FormFile("avatar"); err != nil || userName == "" {
		return err
	}

	var filePath string
	if filePath, err = mkdir(userName); err != nil {
		return err
	}
	filePath = filepath.Join(consts.DefaultAvatarPath, fmt.Sprintf("%s_%s.%s", userName, consts.DefaultAvatarSuffix,
		path.Ext(file.Filename)))
	// todo 保存ext到数据库中

	if err = c.SaveUploadedFile(file, filePath); err != nil { // todo 删除旧头像
		return err
	}

	cookie, _ := c.Cookie(consts.UserCookieName)
	log.Print("cookie:", cookie)

	return nil
}

func mkdir(userName string) (string, error) {
	filePath := filepath.Join(consts.DefaultUserFilePath, userName, time.Now().Format("2006/01/02"))
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Print("目录创建失败", err)
		return "", err
	}
	return filePath, nil
}

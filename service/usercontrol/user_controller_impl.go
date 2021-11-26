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
	"usermanagersystem/utils/rediscontrol"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userControllerImpl struct {
	db *gorm.DB
	rc rediscontrol.RedisController
}

// GetUserMessageByCookie 通过Cookie获取用户信息
func (uc *userControllerImpl) GetUserMessageByCookie(c *gin.Context) (*model.User, error) {
	var account string
	var err error

	if account, err = uc.getAccount(c); err != nil {
		return nil, err
	}

	user := &model.User{Account: account}
	if err := uc.db.Where(user).Take(user).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return user, nil
}

// ModifyPassword 修改密码
func (uc *userControllerImpl) ModifyPassword(c *gin.Context) error {
	var account string
	var err error

	if account, err = uc.getAccount(c); err != nil {
		return err
	}
	// todo 合并检测存在于更新
	// 检测存在
	user := model.User{
		Account:  account,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.PostForm("oldPassword")))),
	}
	if err := uc.db.Where(&user).Take(&user).Error; err == gorm.ErrRecordNotFound {
		return err
	}

	// 更新
	modify := model.User{
		Account:  account,
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.PostForm("newPassword")))),
	}
	if err := uc.db.Model(&user).Updates(modify).Error; err != nil {
		return err
	}

	return nil
}

// UploadAvatar 上传头像
func (uc *userControllerImpl) UploadAvatar(c *gin.Context) error {
	var account string
	var file *multipart.FileHeader
	var err error

	if account, err = uc.getAccount(c); err != nil {
		return err
	}

	if file, err = c.FormFile("avatar"); err != nil {
		return err
	}

	filePath := filepath.Join(consts.DefaultAvatarPath, fmt.Sprintf("%s_%s%s", account, consts.DefaultAvatarSuffix,
		path.Ext(file.Filename)))
	// todo 保存ext到数据库中

	if err = c.SaveUploadedFile(file, filePath); err != nil { // todo 删除旧头像
		return err
	}

	return nil
}

// UploadFile 上传文件
func (uc *userControllerImpl) UploadFile(c *gin.Context) error {
	var account string
	var file *multipart.FileHeader
	var err error
	if account, err = uc.getAccount(c); err != nil {
		return err
	}
	if file, err = c.FormFile("file"); err != nil {
		return err
	}

	var filePath string
	if filePath, err = mkdir(account); err != nil {
		return err
	}
	filePath = filepath.Join(filePath, file.Filename)

	if err = c.SaveUploadedFile(file, filePath); err != nil {
		return err
	}

	return nil
}

// getAccount 通过cookie获取账户
func (uc *userControllerImpl) getAccount(c *gin.Context) (string, error) {
	var account string
	if cookie, err := c.Cookie(consts.UserCookieName); err == nil {
		account, _ = uc.rc.Get(consts.RedisCookieHashPrefix + cookie)
	}
	if account == "" {
		return "", errors.New(consts.CookieTimeOutError)
	}
	return account, nil
}

func mkdir(userName string) (string, error) {
	filePath := filepath.Join(consts.DefaultUserFilePath, userName, time.Now().Format("2006/01/02"))
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Print("目录创建失败", err)
		return "", err
	}
	return filePath, nil
}

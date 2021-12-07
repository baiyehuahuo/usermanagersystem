package usercontrol

import (
	"bytes"
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
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userControllerImpl struct {
	db *gorm.DB
	rc utils.RedisController
}

// GetUserMessageByCookie 通过Cookie获取用户信息
func (uc *userControllerImpl) GetUserMessageByCookie(c *gin.Context) (user *model.User, err error) {
	var account string

	if account, err = uc.getAccountByCookie(c); err != nil {
		return nil, err
	}

	return uc.getUserByAccount(account)
}

// ModifyPassword 修改密码
func (uc *userControllerImpl) ModifyPassword(c *gin.Context, oldPassword string, newPassword string) (err error) {
	var account string

	if account, err = uc.getAccountByCookie(c); err != nil {
		return err
	}
	oldPasswordMD5 := fmt.Sprintf("%x", md5.Sum([]byte(oldPassword)))
	newPasswordMD5 := fmt.Sprintf("%x", md5.Sum([]byte(newPassword)))
	user := model.User{
		Account:  account,
		Password: oldPasswordMD5,
	}
	if rows := uc.db.Where(&user).Updates(&model.User{Password: newPasswordMD5}).RowsAffected; rows == 0 {
		return errors.New(consts.UpdatePasswordFail)
	}

	return nil
}

// UploadAvatar 上传头像
func (uc *userControllerImpl) UploadAvatar(c *gin.Context) (err error) {
	var account string
	var file *multipart.FileHeader

	if account, err = uc.getAccountByCookie(c); err != nil {
		return err
	}

	if file, err = c.FormFile("avatar"); err != nil {
		return err
	}

	filePath := filepath.Join(consts.DefaultAvatarPath, fmt.Sprintf("%s_%s%s", account, consts.DefaultAvatarSuffix, path.Ext(file.Filename)))

	user := model.User{Account: account}
	if err = utils.GetDB().Where(&user).Take(&user).Error; err != nil {
		return err
	}
	if user.AvatarExt != "" {
		if err = uc.rc.DeleteUser(account); err != nil {
			return err
		}
		buffer := bytes.Buffer{}
		buffer.WriteString(consts.DefaultAvatarPath)
		buffer.WriteByte('/')
		buffer.WriteString(user.Account)
		buffer.WriteByte('_')
		buffer.WriteString(consts.DefaultAvatarSuffix)
		buffer.WriteString(user.AvatarExt)
		if err = os.Remove(buffer.String()); err != nil {
			return err
		}
	}
	if err = utils.GetDB().Where(&user).Updates(&model.User{AvatarExt: path.Ext(file.Filename)}).Error; err != nil {
		return err
	}

	// todo 如果保存文件失败 那数据库里的数据怎么办？
	// fmt.Println(filePath)
	if err = c.SaveUploadedFile(file, filePath); err != nil { // todo 删除旧头像
		return err
	}

	if err = uc.rc.DeleteUser(account); err != nil {
		return err
	}
	// if err = chmodFile(filePath, 0444); err != nil {
	// 	return err
	// }

	return nil
}

// UploadFile 上传文件
func (uc *userControllerImpl) UploadFile(c *gin.Context) (err error) {
	var account string
	var file *multipart.FileHeader
	if account, err = uc.getAccountByCookie(c); err != nil {
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
func (uc *userControllerImpl) getAccountByCookie(c *gin.Context) (account string, err error) {
	if cookie, err := c.Cookie(consts.CookieNameOfUser); err == nil {
		account, _ = uc.rc.Get(consts.RedisCookieHashPrefix + cookie)
	}
	if account == "" {
		return "", errors.New(consts.CookieTimeOutError)
	}
	return account, nil
}

func (uc *userControllerImpl) getUserByAccount(account string) (user *model.User, err error) {
	user, err = uc.rc.GetUser(account)
	if user != nil && err == nil {
		// log.Printf("get user %s from redis", account)
		return user, nil
	}
	// log.Printf("not get user %s from redis %v: %v", account, user, err)
	user = &model.User{Account: account}
	if err = uc.db.Where(user).Take(user).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	}
	if err = uc.rc.SetUser(*user); err != nil {
		return nil, err
	}
	return user, nil
}

func mkdir(userName string) (filePath string, err error) {
	filePath = filepath.Join(consts.DefaultUserFilePath, userName, time.Now().Format("2006/01/02"))
	if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Print("目录创建失败", err)
		return "", err
	}
	return filePath, nil
}

func chmodFile(filePath string, mod os.FileMode) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	if err = f.Chmod(mod); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

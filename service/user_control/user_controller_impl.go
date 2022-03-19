package user_control

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/utils"

	"github.com/pkg/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userControllerImpl struct {
	db *gorm.DB
	rc utils.RedisController
}

// GetUserFilesPath 获取该用户的所有png文件
func (uc *userControllerImpl) GetUserFilesPath(c *gin.Context, account string) (result []string, err error) {
	var filesDirPath string
	if filesDirPath, err = getUploadPngDirPath(account); err != nil {
		return nil, utils.ErrWrapOrWithMessage(false, err)
	}
	var fileListInfo []os.FileInfo
	if fileListInfo, err = ioutil.ReadDir(filesDirPath); err != nil {
		return nil, utils.ErrWrapOrWithMessage(true, err)
	}
	result = make([]string, 0, len(fileListInfo))
	for i := range fileListInfo {
		result = append(result, utils.GetNetUploadFilePath(account, fileListInfo[i].Name()))
	}
	return result, nil
}

// GetUserMessageByCookie 通过Cookie获取用户信息
func (uc *userControllerImpl) GetUserMessageByCookie(c *gin.Context) (user *model.User, Err model.Err) {
	var account string
	var err error
	if account, Err = uc.GetAccountByCookie(c); Err.Code != consts.OperateSuccess {
		// return nil, err
		return nil, Err
	}
	if user, err = uc.getUserByAccount(account); err != nil {
		Err.Code = consts.UserNotFound
		Err.Msg = utils.ErrWrapOrWithMessage(false, err).Error()
		return nil, Err
	}
	Err.Code = consts.OperateSuccess
	return user, Err
}

// ModifyPassword 修改密码
func (uc *userControllerImpl) ModifyPassword(c *gin.Context, account, oldPassword string, newPassword string) (err error) {
	oldPasswordMD5 := fmt.Sprintf("%x", md5.Sum([]byte(oldPassword)))
	newPasswordMD5 := fmt.Sprintf("%x", md5.Sum([]byte(newPassword)))
	user := model.User{
		Account:  account,
		Password: oldPasswordMD5,
	}
	if rows := uc.db.Where(&user).Updates(&model.User{Password: newPasswordMD5}).RowsAffected; rows == 0 {
		// return utils.ErrWrapOrWithMessage(true, errors.New(consts.UpdatePasswordFail))
		return errors.New("")
	}

	return nil
}

// PredictPng 预测的图片
func (uc *userControllerImpl) PredictPng(c *gin.Context, account string, pngName string) (predictPath string, err error) {
	var filePath string
	if filePath, err = getUploadPngDirPath(account); err != nil {
		return "", utils.ErrWrapOrWithMessage(false, err)
	}
	filePath = filepath.Join(filePath, pngName)
	cmd := exec.Command("main.exe", filePath)
	if err := cmd.Run(); err != nil {
		return "", utils.ErrWrapOrWithMessage(true, err)
	}

	return utils.GetNetUploadFilePath(account, pngName[:len(pngName)-len(path.Ext(pngName))]+"_predict.png"), nil
}

func (uc *userControllerImpl) DeletePng(c *gin.Context, account string, pngName string) (err error) {
	var filePath string
	if filePath, err = getUploadPngDirPath(account); err != nil {
		return utils.ErrWrapOrWithMessage(false, err)
	}
	filePath = filepath.Join(filePath, pngName)
	if err = os.Remove(filePath); err != nil {
		return utils.ErrWrapOrWithMessage(true, err)
	}
	return nil
}

// SetPassword 按照邮箱设置密码
func (uc *userControllerImpl) SetPassword(c *gin.Context, email string, password string) (Err model.Err) {
	newPasswordMD5 := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	user := model.User{
		Email: email,
	}
	if rows := uc.db.Where(&user).Updates(&model.User{Password: newPasswordMD5}).RowsAffected; rows == 0 {
		Err.Code = consts.UpdatePasswordFail
		Err.Msg = utils.ErrWrapOrWithMessage(true, errors.New(consts.ErrCodeMessage[Err.Code])).Error()
		return Err
	}
	Err.Code = consts.OperateSuccess
	return Err
}

// UploadAvatar 上传头像
func (uc *userControllerImpl) UploadAvatar(c *gin.Context) (Err model.Err) {
	var account string
	var file *multipart.FileHeader
	var err error
	if account, Err = uc.GetAccountByCookie(c); Err.Code != consts.OperateSuccess {
		return Err
	}

	if file, err = c.FormFile("avatar"); err != nil {
		Err.Code = consts.InputParamsWrong
		Err.Msg = utils.ErrWrapOrWithMessage(true, err).Error()
		return Err
	}

	filePath := filepath.Join(consts.DefaultAvatarPath, fmt.Sprintf("%s_%s%s", account, consts.DefaultAvatarSuffix, path.Ext(file.Filename)))

	user := model.User{Account: account}
	if err = utils.GetDB().Where(&user).Take(&user).Error; err != nil {
		Err.Code = consts.DatabaseWrong
		Err.Msg = utils.ErrWrapOrWithMessage(true, err).Error()
		return Err
	}
	if user.AvatarExt != "" {
		if err = uc.rc.DeleteUser(account); err != nil {
			Err.Code = consts.DatabaseWrong
			Err.Msg = utils.ErrWrapOrWithMessage(false, err).Error()
			return Err
		}
		if err = os.Remove(utils.GetLocalAvatarPath(user.Account, user.AvatarExt)); err != nil {
			Err.Code = consts.SystemError
			Err.Msg = utils.ErrWrapOrWithMessage(true, err).Error()
			return Err
		}
	}
	if err = utils.GetDB().Where(&user).Updates(&model.User{AvatarExt: path.Ext(file.Filename)}).Error; err != nil {
		Err.Code = consts.DatabaseWrong
		Err.Msg = utils.ErrWrapOrWithMessage(true, err).Error()
		return Err
	}

	// todo 如果保存文件失败 那数据库里的数据怎么办？
	// fmt.Println(filePath)
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		Err.Code = consts.SystemError
		Err.Msg = utils.ErrWrapOrWithMessage(true, err).Error()
		return Err
	}

	if err = uc.rc.DeleteUser(account); err != nil {
		Err.Code = consts.DatabaseWrong
		Err.Msg = utils.ErrWrapOrWithMessage(false, err).Error()
		return Err
	}
	// if err = chmodFile(filePath, 0444); err != nil {
	// 	return errors.WithMessage(err, utils.RunFuncNameWithFail())
	// }
	Err.Code = consts.OperateSuccess
	return Err
}

// UploadFile 上传文件
func (uc *userControllerImpl) UploadPng(c *gin.Context, account string) (err error) {
	var file *multipart.FileHeader
	if file, err = c.FormFile("file"); err != nil {
		return utils.ErrWrapOrWithMessage(true, err)
	}

	var filePath string
	if filePath, err = getUploadPngDirPath(account); err != nil {
		return utils.ErrWrapOrWithMessage(false, err)
	}
	filePath = filepath.Join(filePath, file.Filename)

	if err = c.SaveUploadedFile(file, filePath); err != nil {
		return utils.ErrWrapOrWithMessage(true, err)
	}

	return nil
}

// getAccount 通过cookie获取账户
func (uc *userControllerImpl) GetAccountByCookie(c *gin.Context) (account string, Err model.Err) {
	if cookie, err := c.Cookie(consts.CookieNameOfUser); err == nil {
		account, _ = uc.rc.Get(consts.RedisCookieHashPrefix + cookie)
	}
	if account == "" {
		Err.Code = consts.InputParamsWrong
		Err.Msg = utils.ErrWrapOrWithMessage(true, errors.New(consts.CookieTimeOutError)).Error()
		return "", Err
	}
	Err.Code = consts.OperateSuccess
	return account, Err
}

func (uc *userControllerImpl) getUserByAccount(account string) (user *model.User, err error) {
	user, err = uc.rc.GetUser(account)
	if user != nil && err == nil {
		return user, nil
	}
	user = &model.User{Account: account}
	if err = uc.db.Where(user).Take(user).Error; err == gorm.ErrRecordNotFound {
		return nil, utils.ErrWrapOrWithMessage(true, err)
	}
	if err = uc.rc.SetUser(*user); err != nil {
		return nil, utils.ErrWrapOrWithMessage(false, err)
	}
	return user, nil
}

func getUploadPngDirPath(userName string) (filePath string, err error) {
	filePath = filepath.Join(consts.DefaultUserPngPath, userName)
	if err = os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Print("目录创建失败", err)
		// return "", err
		return "", utils.ErrWrapOrWithMessage(true, err)
	}
	return filePath, nil
}

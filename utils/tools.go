package utils

import (
	"bytes"
	"runtime"
	"usermanagersystem/consts"

	"github.com/pkg/errors"
)

func netPrefixPath() string {
	buffer := bytes.Buffer{}
	buffer.WriteString(consts.HttpDomain)
	buffer.WriteByte('/')
	return buffer.String()
}

// GetNetAvatarPath 获取头像的网络路径
func GetNetAvatarPath(account string, avatarExt string) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(netPrefixPath())
	if avatarExt == "" {
		buffer.WriteString(consts.DefaultStaticPath)
		buffer.WriteString("/default_avatar.jpg")
		return buffer.String()
	}
	buffer.WriteString(GetLocalAvatarPath(account, avatarExt))
	return buffer.String()
}

// GetNetUploadFilePath 获取文件的网络路径
func GetNetUploadFilePath(account string, filePath string) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(netPrefixPath())
	buffer.WriteString(GetLocalUploadFilePath(account, filePath))
	return buffer.String()
}

// GetLocalAvatarPath 获取头像的本地路径
func GetLocalAvatarPath(account string, avatarExt string) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(consts.DefaultAvatarPath)
	buffer.WriteByte('/')
	buffer.WriteString(account)
	buffer.WriteByte('_')
	buffer.WriteString(consts.DefaultAvatarSuffix)
	buffer.WriteString(avatarExt)
	return buffer.String()
}

// GetNetUploadFilePath 获取文件的本地路径
func GetLocalUploadFilePath(account string, fileName string) string {
	buffer := bytes.Buffer{}
	buffer.WriteString(consts.DefaultUserFilePath)
	buffer.WriteByte('/')
	buffer.WriteString(account)
	buffer.WriteByte('/')
	buffer.WriteString(fileName)
	return buffer.String()
}

// ErrWrapOrWithMessage 给错误附加当前运行的函数名
func ErrWrapOrWithMessage(wrap bool, err error) error {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	message := f.Name() + " fail"
	if wrap {
		return errors.Wrap(err, message)
	}
	return errors.WithMessage(err, message)
}

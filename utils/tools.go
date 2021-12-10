package utils

import (
	"bytes"
	"runtime"
	"usermanagersystem/consts"

	"github.com/pkg/errors"
)

// GetNetAvatarPath 获取头像的网络路径
func GetNetAvatarPath(account string, avatarExt string) string {
	buffer := bytes.Buffer{}
	if avatarExt == "" {
		buffer.WriteString(consts.HttpDomain)
		buffer.WriteByte('/')
		buffer.WriteString(consts.DefaultStaticPath)
		buffer.WriteString("/default_avatar.jpg")
		return buffer.String()
	}
	buffer.WriteString(consts.HttpDomain)
	buffer.WriteByte('/')
	buffer.WriteString(GetLocalAvatarPath(account, avatarExt))
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

// ErrWrapOrWithMessage 给错误附加当前运行的函数名
func ErrWrapOrWithMessage(wrap bool, err error) error {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	message := f.Name() + " fail\n"
	if wrap {
		return errors.Wrap(err, message)
	}
	return errors.WithMessage(err, message)
}

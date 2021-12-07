package utils

import (
	"bytes"
	"usermanagersystem/consts"
)

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

package consts

import "time"

const (
	AuthEmailUser          = "测试"
	AuthEmailSubject       = "验证码"
	AuthCodeRandRange      = 1e6
	AuthCodeContinueTime   = 30 * time.Second
	AuthCodeCacheFlushTime = 300 * time.Second

	DefaultFileRootPath = "upload_files"
	DefaultUserFilePath = DefaultFileRootPath + "/user_files"
	DefaultAvatarPath   = "avatars"
	DefaultMoviePath    = DefaultFileRootPath + "/movies"
	DefaultStaticPath   = "static"
	DefaultAvatarSuffix = "avatar"

	CookieNameOfUser       = "UserCookie"
	CookieContinueTime     = 300 // 单位是秒 3600 就是一小时
	CookieValidationRange  = "/"
	CookieValidationDomain = "127.0.0.1"
	CookieTimeOutError     = "cookie is timeout"

	RedisCookieHashPrefix        = "UserCookie_"
	RedisAuthCodePrefix          = "AuthCode_" // 假如保存在redis的前缀名
	RedisUserMessagePrefix       = "UserMessage_"
	RedisUserMessageContinueTime = 30

	UpdatePasswordFail = "update password fail"
)

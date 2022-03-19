package consts

import "time"

const (
	AuthEmailUser          = "测试"
	AuthEmailSubject       = "验证码"
	AuthCodeRandRange      = 1e6
	AuthCodeContinueTime   = 30 * time.Second
	AuthCodeCacheFlushTime = 300 * time.Second

	DefaultUserPngRootPath = "user_files"
	DefaultUserPngPath     = DefaultUserPngRootPath + "/user_png"
	DefaultAvatarPath      = "avatars"
	DefaultStaticPath      = "static"
	DefaultAvatarSuffix    = "avatar"

	CookieNameOfUser            = "UserCookie"
	CookieContinueTime          = 300 // 单位是秒 3600 就是一小时
	CookieValidationRange       = "/"
	CookieValidationDomainIP    = "127.0.0.1"
	CookieValidationDomainLocal = "localhost"
	CookieTimeOutError          = "cookie is timeout"

	PredictSuccess = "PredictPng success"
	PredictFail    = "PredictPng fail"

	RedisCookieHashPrefix        = "UserCookie_"
	RedisAuthCodePrefix          = "AuthCode_" // 假如保存在redis的前缀名
	RedisUserMessagePrefix       = "UserMessage_"
	RedisUserMessageContinueTime = 30

	SystemLogPath   = "systemLogs"
	LogFilePath     = SystemLogPath + "/log.txt"
	MySQLBackUpPath = SystemLogPath + "/mysql_backup.sql"
)

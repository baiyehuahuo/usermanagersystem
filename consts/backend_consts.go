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
	DefaultAvatarSuffix    = "avatar"

	CookieNameOfUser            = "UserCookie"
	CookieContinueTime          = 300 // 单位是秒 3600 就是一小时
	CookieValidationRange       = "/"
	CookieValidationDomainIP    = "127.0.0.1"
	CookieValidationDomainLocal = "localhost"

	RedisCookieHashPrefix        = "UserCookie_"
	RedisUserMessagePrefix       = "UserMessage_"
	RedisUserMessageContinueTime = 30

	SystemLogPath   = "systemLogs"
	LogFilePath     = SystemLogPath + "/log.txt"
	MySQLBackUpPath = SystemLogPath + "/mysql_backup.sql"

	PredictQueueName = "PredictQueue"
	ExchangeName     = "ResultExchange"
	RouteType        = "fanout"
	AutoDelete       = true
	MaxQueueLength   = 10
)

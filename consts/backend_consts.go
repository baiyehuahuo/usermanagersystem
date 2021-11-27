package consts

const (
	DefaultFileRootPath = "uploadfiles"
	DefaultUserFilePath = DefaultFileRootPath + "/userfiles"
	DefaultAvatarPath   = DefaultFileRootPath + "/avatars"
	DefaultAvatarSuffix = "avatar"

	UserCookieName         = "UserCookie"
	CookieContinueTime     = 30 // 单位是秒 3600 就是一小时
	CookieValidationRange  = "/"
	CookieValidationDomain = "127.0.0.1:8080"
	CookieTimeOutError     = "cookie is timeout"

	RedisCookieHashPrefix        = "UserCookie_"
	RedisAuthCodePrefix          = "AuthCode_"
	RedisUserMessagePrefix       = "UserMessage_"
	RedisUserMessageContinueTime = 30

	AuthEmailUser          = "测试"
	AuthEmailSubject       = "验证码"
	AuthCodeRandRange      = 1e6
	AuthCodeContinueTime   = 30  // 单位是秒
	AuthCodeCacheFlushTime = 300 // 单位是秒
)

package consts

const (
	DefaultFileRootPath = "uploadfiles"

	UserCookieName         = "UserCookie"
	CookieContinueTime     = 30 // 单位是秒 3600 就是一小时
	CookieValidationRange  = "/"
	CookieValidationDomain = "127.0.0.1:8080"

	RedisCookieHashPrefix = "UserCookie_"
	RedisAuthCodePrefix   = "AuthCode_"

	AuthEmailUser        = "测试"
	AuthEmailSubject     = "验证码"
	AuthCodeRandRange    = 1e6
	AuthCodeContinueTime = 30 // 单位是秒
)

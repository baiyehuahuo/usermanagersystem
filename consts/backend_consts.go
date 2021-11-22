package consts

const (
	DefaultFileRootPath = "uploadfiles"

	UserCookieName         = "UserCookie"
	CookieContinueTime     = 30 // 单位是秒 3600 就是一小时
	CookieValidationRange  = "/"
	CookieValidationDomain = "127.0.0.1:8080"

	RedisCookieHash = "UserCookie_"
)

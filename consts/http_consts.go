package consts

const (
	LoginSuccess = "Login Success"
	LoginFail    = "Login Fail"

	RegeditSuccess = "Regedit Success"
	RegeditFail    = "Regedit Fail"

	UploadSuccess = "Upload Success"
	UploadFail    = "Upload Fail"

	DefaultFileRootPath = "uploadfiles"

	UserCookieName         = "UserCookie"
	CookieContinueTime     = 3600 // 单位是秒 3600 就是一小时
	CookieValidationRange  = "/"
	CookieValidationDomain = "127.0.0.1:8080"
)

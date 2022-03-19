package consts

const (
	OperateSuccess = 200

	InputParamsWrong        = 100000
	UserNotFound            = 100001
	AccountOrPasswordWrong  = 100002
	CheckAuthCodeFail       = 100003
	EmailIsRegistered       = 100004
	UpdatePasswordFail      = 100005
	CookieTimeOut           = 100006
	SendAuthCodeByEmailFail = 100011
	SystemError             = 999998
	DatabaseWrong           = 999999
)

var (
	ErrCodeMessage = map[int]string{
		OperateSuccess:          "操作成功",
		InputParamsWrong:        "输入参数格式错误",
		UserNotFound:            "用户不存在",
		AccountOrPasswordWrong:  "账号或密码错误",
		CheckAuthCodeFail:       "验证码错误",
		EmailIsRegistered:       "邮箱已被注册",
		UpdatePasswordFail:      "更新密码失败",
		CookieTimeOut:           "Cookie过期",
		SendAuthCodeByEmailFail: "邮箱验证码发送失败",
		SystemError:             "系统错误",
		DatabaseWrong:           "数据库错误",
	}
)

package ecode

var MessageMap = map[int]string{
	Success: "success",
	Error:   "服务端异常",

	AccountException:       "账号异常",
	Unauthorized:           "未授权：请登录",
	LoginFailed:            "用户名或密码错误",
	PermissionDenied:       "权限不足",
	PageNotFound:           "404 Not Found",
	ValidatesRequestsError: "表单验证错误",

	ResultEmpty:    "数据为空",
	OrmUpdateError: "参数错误，数据无法更新",
	OrmCreateError: "参数错误，数据无法创建",
	ParamsError:    "参数错误",
}

func GetMsg(code int) string {
	msg, ok := MessageMap[code]

	if ok {
		return msg
	}

	return MessageMap[Error]
}

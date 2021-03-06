package errors

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_USER_EXIST     = 1001
	ERROR_USER_NOT_EXIST = 1002
	ERROR_PASSWORD_WRONG = 1003

	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_EXPIRED    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008

	ERROR_FORBIDDEN_EDIT_ADMIN = 1009

	ERROR_USER_DISABLE = 1010
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ERROR_USER_EXIST:     "用户已存在",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_PASSWORD_WRONG: "密码错误",

	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_EXPIRED:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_FORBIDDEN_EDIT_ADMIN: "禁止编辑管理员",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

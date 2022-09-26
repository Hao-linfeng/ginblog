package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	ERROR_TOKEN_WRONG    = 1006
	ERROR_TOKEN_TYPE     = 1007
	ERROR_USER_NO_RIGHT  = 1008

	ERROR_ART_NOT_EXIST = 2001

	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

var codemsg map[int]string = map[int]string{
	SUCCSE:               "OK",
	ERROR:                "FAIL",
	ERROR_USERNAME_USED:  "用户名已存在",
	ERROR_PASSWORD_WRONG: "密码错误",
	ERROR_USER_NOT_EXIST: "用户不存在",
	ERROR_TOKEN_EXIST:    "token 不存在",
	ERROR_TOKEN_RUNTIME:  "token 已过期",
	ERROR_TOKEN_WRONG:    "token 不正确",
	ERROR_TOKEN_TYPE:     "token格式错误",
	ERROR_CATENAME_USED:  "文章不存在",
	ERROR_ART_NOT_EXIST:  "该文章不存在",
	ERROR_CATE_NOT_EXIST: "分类已经存在",
	ERROR_USER_NO_RIGHT:  "该用户无权限",
}

func GetErrMsg(code int) string {
	return codemsg[code]

}

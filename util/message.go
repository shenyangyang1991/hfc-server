package util

const (
	SUCCESS       = 1
	ERROR         = 500
	BAD_PARAM     = 400
	TOKEN_INVALID = 403
	NOT_FOUND     = 404

	LOGIN_ERROR         = 1000
	LOGIN_BAD_PARAM     = 1001
	LOGIN_INVALID_PARAM = 1002
	LOGIN_INVALID_SMS   = 1003
	LOGIN_INVALID_USER  = 1004
	REGISTER_ERROR      = 1005

	USER_INVALID_NAME = 2001
	USER_ERROR_UPDATE = 2002
	USER_BAD_NAME     = 2003
)

var msg = map[int]string{
	SUCCESS:       "请求成功",
	ERROR:         "请求失败",
	BAD_PARAM:     "请求参数错误",
	TOKEN_INVALID: "未登录或登录已过期，请重新登录",
	NOT_FOUND:     "未找到",

	LOGIN_ERROR:         "登录失败",
	LOGIN_BAD_PARAM:     "请输入手机号码与短信验证码后，点击登录",
	LOGIN_INVALID_PARAM: "请检查手机号码与短信验证码是否正确",
	LOGIN_INVALID_SMS:   "短信验证码错误",
	LOGIN_INVALID_USER:  "用户严重违规，已经被永久封号",
	REGISTER_ERROR:      "新帐号创建失败，请重试",

	USER_INVALID_NAME: "昵称已存在或昵称含有敏感字",
	USER_BAD_NAME:     "请输入昵称或上传头像",
	USER_ERROR_UPDATE: "更新信息失败",
}

func GetMsg(code int) string {
	return msg[code]
}

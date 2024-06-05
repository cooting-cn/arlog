package er

var codeMsg = map[int]string{
	200: "OK",
	201: "FAIL",
	202: "用户名已存在！",
	203: "密码错误",
	204: "用户不存在",
	205: "TOKEN不存在,请重新登陆",
	206: "需要otp 2步验证",
	207: "TOKEN不正确,请重新登陆",
	208: "otp验证错误!",
	403: "该用户无权限",
	210: "文章不存在",
	211: "该分类已存在",
	212: "该分类不存在",
	213: "书库链接失败",
	214: "文章添加失败",
	215: "该tag已存在",
	216: "该tag不存在",
}

// GetErr 统一的web访问报错
func GetErr(code int) string {
	return codeMsg[code]
}

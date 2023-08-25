package ttkerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	commonErrorInit()
	userServiceErrorInit()
}

func commonErrorInit() {
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[RequestParamError] = "参数错误"
	message[TokenExpireError] = "token失效，请重新登陆"
	message[TokenGenerateError] = "生成token失败"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[DbUpdateAffectedZeroError] = "更新数据影响行数为0"
	message[DataNoExistError] = "数据不存在"
	message[PhoneValidError] = "手机格式错误"
	message[EmailValidError] = "邮箱格式错误"
}

func userServiceErrorInit() {
	message[UserNotFountError] = "用户不存在"
	message[PassportError] = "用户登录失败"
	message[VerifyCodeNotFoundError] = "验证码错误"
	message[VerifyCodeJudgeError] = "验证码错误"
	message[SendVerifyCodeFrequentError] = "发送验证码过于频繁"
	message[TTIdValidError] = "TTKId格式不合法"
}

func MapErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}

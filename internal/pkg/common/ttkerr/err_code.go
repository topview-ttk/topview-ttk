package ttkerr

// OK 成功返回
const OK uint32 = 0

/**全局错误码*/
const (
	// Unauthorized 未鉴权
	Unauthorized uint32 = 100000 + iota

	// ServerCommonError 服务器开小差
	ServerCommonError
	// RequestParamError 请求参数错误
	RequestParamError

	// TokenExpireError token过期
	TokenExpireError

	// TokenGenerateError 生成token失败
	TokenGenerateError

	// DbError 数据库繁忙,请稍后再试
	DbError

	// DbUpdateAffectedZeroError 更新数据影响行数为0
	DbUpdateAffectedZeroError

	// DataNoExistError 数据不存在
	DataNoExistError

	// PhoneValidError 手机正则验证失败
	PhoneValidError
	// EmailValidError 邮箱正则验证失败
	EmailValidError
	// PhoneRegisteredError 手机已被注册
	PhoneRegisteredError
	// EmailRegisteredError 邮箱已被注册
	EmailRegisteredError
)

/**(前3位代表业务,后三位代表具体功能)**/

// 120为用户服务
const (
	// UserNotFountError 无法查到用户
	UserNotFountError uint32 = 120000 + iota
	// PassportError TTKId/mobile/email or password错误
	PassportError
	// VerifyCodeNotFoundError 验证码找不到错误
	VerifyCodeNotFoundError
	// VerifyCodeJudgeError 验证码错误
	VerifyCodeJudgeError

	SendVerifyCodeFrequentError

	TTIdValidError
)

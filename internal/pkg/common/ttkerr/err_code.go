package ttkerr

// OK 成功返回
const OK uint32 = 0

/**全局错误码*/
const (
	// Unauthorized 未鉴权
	Unauthorized uint32 = 100000

	// ServerCommonError 服务器开小差
	ServerCommonError uint32 = 100001
	// RequestParamError 请求参数错误
	RequestParamError uint32 = 100002

	// TokenExpireError token过期
	TokenExpireError uint32 = 100003

	// TokenGenerateError 生成token失败
	TokenGenerateError uint32 = 100004

	// DbError 数据库繁忙,请稍后再试
	DbError uint32 = 100005

	// DbUpdateAffectedZeroError 更新数据影响行数为0
	DbUpdateAffectedZeroError uint32 = 100006

	// DataNoExistError 数据不存在
	DataNoExistError uint32 = 100007
)

/**(前3位代表业务,后三位代表具体功能)**/

// 120为用户服务
const (
	UserNotFountError uint32 = 120000
)

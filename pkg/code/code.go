package code

const (
	// 1000xx 通用模块

	// ErrParameterInvalid @message=参数校验失败: %s
	ErrParameterInvalid = 100001

	// 1001xx 用户模块错误
	// ErrUserNotFound @message=用户未找到: %s
	ErrUserNotFound = 100101

	// ErrUserAuthFaild @message=用户认证失败
	ErrUserAuthFaild = 100102

	// ErrUserExist @message=用户已存在: %s
	ErrUserExist = 100103
)

package code

import "fmt"

const (
	// ErrUnknown @message=未知错误
	ErrUnknownCode = 100001

	// ErrDBUnknown @message=数据库未知异常
	ErrDBUnknown = 100002
)

var (
	ErrDB      = &RespError{ErrDBUnknown, "数据库未知异常"}
	ErrUnknown = &RespError{ErrUnknownCode, "未知错误"}
)

type RespError struct {
	Code    int
	Message string
}

func (r *RespError) Error() string {
	return r.Message
}

var ErrMap map[int]*RespError

func GetErrWithCode(code int) (e error) {
	if v, ok := ErrMap[code]; ok {
		return v
	} else {
		return ErrUnknown
	}
}

func NewErrWithCode(code int, args ...any) error {
	if v, ok := ErrMap[code]; ok {
		v.Message = fmt.Sprintf(v.Message, args...)
		return v
	} else {
		return ErrUnknown
	}
}

package errcode

//go:generate stringer -type ErrCode -linecomment
const (
	ErrUnknown  ErrCode = 1 // 未知错误
	ErrNotFound ErrCode = 2 // 未找到
)

type ErrCode int

func (e ErrCode) Error() string {
	return e.String()
}

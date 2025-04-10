package errcode

// Error 定义错误结构体，用于Swagger文档生成
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	//TODO implement me
	panic("implement me")
}

// NewCustomErr 创建自定义错误
func NewCustomErr(msg string) *Error {
	return &Error{
		Code:    400,
		Message: msg,
	}
}

// 预定义错误
var (
	ErrInvalidParams = &Error{Code: 400, Message: "Invalid parameters"}
	ErrUnexpected    = &Error{Code: 500, Message: "Unexpected error"}
)

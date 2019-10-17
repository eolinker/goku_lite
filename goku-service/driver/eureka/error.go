package eureka

import "fmt"

const (
	nilPointCode = 20001
)

//Error error
type Error struct {
	Message string
	Code    int
}

func _error(innerCode int, innerMessage string, message string) *Error {
	return &Error{Message: innerMessage + ",cause by: " + message, Code: innerCode}
}

//NewError 创建Error
func NewError(message string, code int) *Error {
	return &Error{Message: message, Code: code}
}

//NilPointError 空指针错误
func NilPointError(message string) *Error {
	return _error(nilPointCode, "nil point error", message)
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

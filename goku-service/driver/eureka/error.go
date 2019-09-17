package eureka

import "fmt"

const (
	nilPointCode = 20001
)

type Error struct {
	Message string
	Code    int
}

func _error(innerCode int, innerMessage string, message string) *Error {
	return &Error{Message: innerMessage + ",cause by: " + message, Code: innerCode}
}

func NewError(message string, code int) *Error {
	return &Error{Message: message, Code: code}
}

func NilPointError(message string) *Error {
	return _error(nilPointCode, "nil point error", message)
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

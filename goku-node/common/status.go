package common

//StatusHandler 状态处理器
type StatusHandler struct {
	code   int
	status string
}

//SetStatus 设置状态信息
func (s *StatusHandler) SetStatus(code int, status string) {
	s.code, s.status = code, status
}

//StatusCode 获取状态码
func (s *StatusHandler) StatusCode() int {
	return s.code
}

//Status 获取状态
func (s *StatusHandler) Status() string {
	return s.status
}

//NewStatusHandler 状态处理器
func NewStatusHandler() *StatusHandler {
	return new(StatusHandler)
}

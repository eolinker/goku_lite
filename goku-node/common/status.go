package common

type StatusHandler struct {
	code   int
	status string
}

func (s *StatusHandler) SetStatus(code int, status string) {
	s.code, s.status = code, status
}

func (s *StatusHandler) StatusCode() int {
	return s.code
}

func (s *StatusHandler) Status() string {
	return s.status
}

func NewStatusHandler() *StatusHandler {
	return new(StatusHandler)
}

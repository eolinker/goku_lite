package node

import (
	"github.com/eolinker/goku-api-gateway/admin/cmd"
)

type Callback func(code cmd.Code, data []byte) error

func (c Callback) ServerCode(code cmd.Code, data []byte) error {
	return c(code, data)
}

type CodeHandler interface {
	ServerCode(code cmd.Code, data []byte) error
}

//Register cmd 回调注册器
type Register struct {
	callbacks map[cmd.Code][]CodeHandler
}

//NewRegister create register
func NewRegister() *Register {
	return &Register{
		callbacks: make(map[cmd.Code][]CodeHandler),
	}
}

//Register 注册回调
func (s *Register) Register(code cmd.Code, handler CodeHandler) {
	s.callbacks[code] = append(s.callbacks[code], handler)
}

//Register 注册回调
func (s *Register) RegisterFunc(code cmd.Code, callback func(code cmd.Code, data []byte) error) {
	s.callbacks[code] = append(s.callbacks[code], Callback(callback))
}

//Callback 调用回调
func (s *Register) Callback(code cmd.Code, data []byte) error {
	m := s.callbacks

	for _, handler := range m[code] {
		if e := handler.ServerCode(code, data); e != nil {
			return e
		}
	}
	return nil
}

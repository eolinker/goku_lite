package console

import (
	"github.com/eolinker/goku-api-gateway/admin/cmd"
	goku_log "github.com/eolinker/goku-api-gateway/goku-log"
)

type Callback func(code cmd.Code, data []byte,client *Client) error

func (c Callback) ServerCode(code cmd.Code, data []byte, client *Client) error{
	return c(code,data,client)
}

type CodeHandler interface {
	ServerCode(code cmd.Code, data []byte,client *Client)error
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
func (s *Register) Register(code cmd.Code,handler CodeHandler){
	s.callbacks[code] = append(s.callbacks[code],handler)
}
//Register 注册回调
func (s *Register) RegisterFunc(code cmd.Code,callback func(code cmd.Code, data []byte,client *Client) error){
	s.callbacks[code] = append(s.callbacks[code],Callback(callback))
}

//Callback 调用回调
func (s *Register)Callback(code cmd.Code,data []byte,client *Client)error  {
	m:=s.callbacks
	callbacks,has:=  m[code]
	if !has{
		goku_log.Info("not exists call for ",code)
		return nil
	}
	for _,handler:=range callbacks{
		if e:=handler.ServerCode(code,data,client);e!=nil{
			return e
		}
	}
	return nil
}
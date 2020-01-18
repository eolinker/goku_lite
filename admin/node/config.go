package node

import (
	"github.com/eolinker/goku-api-gateway/admin/cmd"
)


func (c *TcpConsole)OnConfigChange(code cmd.Code,data []byte) error {

	conf,err:= cmd.DecodeConfig(data)
	if err!=nil{
		return err
	}
	c.lastConfig.Set(conf)
	c.listener.Call(conf)
	return nil
}
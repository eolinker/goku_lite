package console

import (
	"github.com/eolinker/goku-api-gateway/admin/cmd"
	"github.com/eolinker/goku-api-gateway/console/module/versionConfig"
	log "github.com/eolinker/goku-api-gateway/goku-log"
)


var(
	callbacksInit = NewRegister()
)

func doRegister()*Register{

	r:=callbacksInit
	callbacksInit = nil
	versionConfig.AddCallback(OnConfigChange)
	return r
}
func AddRegisterHandler(code cmd.Code,handler CodeHandler)  {
	if callbacksInit == nil{
		log.Panic("not allow register now")
	}
	callbacksInit.Register(code,handler)
}

func AddRegisterFunc(code cmd.Code,handleFunc func(code cmd.Code, data []byte,client *Client) error)  {
	if callbacksInit == nil{
		log.Panic("not allow register now")
	}
	callbacksInit.RegisterFunc(code,handleFunc)
}
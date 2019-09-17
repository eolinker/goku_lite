package console

import (
	module "github.com/eolinker/goku/console/module/config-log"
	log "github.com/eolinker/goku/goku-log"
)

func InitLog()  {
	c, _ := module.Get(module.ConsoleLog)

	period,_:=log.ParsePeriod(c.Period)
	log.SetOutPut(c.Enable ,c.Dir,c.File,period,c.Expire)
	l,_:= log.ParseLevel(c.Level)
	log.SetLevel(l)
}
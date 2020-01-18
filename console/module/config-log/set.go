package config_log

import (
	"fmt"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	entity "github.com/eolinker/goku-api-gateway/server/entity/config-log"
)

//Set 设置参数
func Set(name string, param *Param) error {
	if _, has := logNames[name]; !has {
		return fmt.Errorf("not has that log config of %s", name)
	}
	c := new(entity.LogConfig)
	c.Name = name
	c.Level = param.Level

	c.Period = param.Period
	c.File = param.File
	c.Dir = param.Dir
	if param.Enable {
		c.Enable = 1
	} else {
		c.Enable = 0
	}
	c.Fields = param.Fields
	c.Expire = param.Expire
	err := configLogDao.Set(c)
	if err != nil {
		return err
	}
	if name == ConsoleLog {
		go reset(c)
	}
	return nil
}
func reset(c *entity.LogConfig) {

	period, _ := log.ParsePeriod(c.Period)
	log.SetOutPut(c.Enable == 1, c.Dir, c.File, period, c.Expire)
	l, _ := log.ParseLevel(c.Level)
	log.SetLevel(l)

}

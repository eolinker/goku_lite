package config_log

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/common/auto-form"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	dao "github.com/eolinker/goku-api-gateway/server/dao/config-log"
)

func Get(name string) (*LogConfig, error) {
	if _, has := logNames[name]; !has {
		return nil, fmt.Errorf("not has that log config of %s", name)
	}

	c := &LogConfig{}
	c.Levels = Levels
	c.Periods = Periods
	c.Expires = Expires
	config, e := dao.Get(name)

	if e != nil || config == nil {
		auto.SetDefaults(c)
		c.Name = name
		c.File = name
		c.Level = log.ErrorLevel.String()
		c.Period = log.PeriodHour.String()
		c.Expire = ExpireDefault
	} else {
		c.Read(config)
	}

	return c, nil
}

func GetAccess() (*AccessConfig, error) {
	config, e := dao.Get(AccessLog)
	c := new(AccessConfig)
	c.Periods = Periods
	c.Expires = Expires
	if e != nil || config == nil {
		auto.SetDefaults(c)
		c.Name = AccessLog

		c.Period = log.PeriodHour.String()
		c.Expire = ExpireDefault
		c.InitFields()
	} else {
		c.Read(config)
	}
	return c, nil
}

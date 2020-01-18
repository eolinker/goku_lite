package server

import (
	"github.com/eolinker/goku-api-gateway/config"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	access_log "github.com/eolinker/goku-api-gateway/goku-node/access-log"
	access_field "github.com/eolinker/goku-api-gateway/server/access-field"
)

//SetLog setLog
func SetLog(c *config.LogConfig) {

	if c == nil {

		c = defaultNodeAppLogConfig()
	}

	period, err := log.ParsePeriod(c.Period)
	if err != nil {
		period = log.PeriodDay
		log.Warn("manager/config unmarshal access log period failed for nod , use the default config:%s", err)
	}

	level, err := log.ParseLevel(c.Level)
	if err != nil {
		level = log.WarnLevel
		log.Warn("manager/config unmarshal access log level failed for nod , use the default config:%s", err)
	}

	enable := c.Enable == 1
	log.SetOutPut(enable, c.Dir, c.File, period, c.Expire)
	log.SetLevel(level)
}

//SetAccessLog setAccessLog
func SetAccessLog(c *config.AccessLogConfig) {
	if c == nil {
		c = defaultAccessLogConfig()
	}
	period, err := log.ParsePeriod(c.Period)
	if err != nil {
		period = log.PeriodDay
		log.Warn("manager/config unmarshal period failed for , use the default config:%s", err)
	}
	enable := c.Enable == 1

	fieldsConfig := c.Fields
	if err != nil || len(fieldsConfig) == 0 {
		log.Warn("manager/config unmarshal access log fields error:", err)

		access_log.SetFields(access_field.Default())
	} else {
		fields := make([]access_field.AccessFieldKey, 0, access_field.Size())
		for _, f := range fieldsConfig {
			fields = append(fields, access_field.Parse(f))
		}
		access_log.SetFields(fields)
	}

	access_log.SetOutput(enable, c.Dir, c.File, period, c.Expire)
}

func defaultAccessLogConfig() *config.AccessLogConfig {
	return &config.AccessLogConfig{
		Name:   "access",
		Enable: 1,
		Dir:    "work/logs/",
		File:   "access.log",
		Period: "hour",
		Expire: 3,
		Fields: nil,
	}
}

func defaultNodeAppLogConfig() *config.LogConfig {
	return &config.LogConfig{
		Name:   "node",
		Enable: 1,
		Dir:    "work/logs/",
		File:   "node.log",
		Level:  "error",
		Period: "hour",
	}
}

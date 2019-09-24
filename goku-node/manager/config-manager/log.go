package config_manager

import (
	"encoding/json"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	access_log "github.com/eolinker/goku-api-gateway/goku-node/access-log"
	"github.com/eolinker/goku-api-gateway/goku-node/manager/updater"
	access_field "github.com/eolinker/goku-api-gateway/server/access-field"
	dao "github.com/eolinker/goku-api-gateway/server/dao/config-log"
	entity "github.com/eolinker/goku-api-gateway/server/entity/config-log"
)

const (
	AccessLog = "access"
	NodeLog   = "node"
)

func init() {
	updater.Add(reloadLogConfig, 1, "goku_config_log")
}
func defaultAccessLogConfig() *entity.LogConfig {
	return &entity.LogConfig{
		Name:   AccessLog,
		Enable: 1,
		Dir:    "work/logs/",
		File:   "access.log",

		Period: "hour",
		Fields: "",
	}
}
func defaultNodeAppLogConfig() *entity.LogConfig {
	return &entity.LogConfig{
		Name:   AccessLog,
		Enable: 1,
		Dir:    "work/logs/",
		File:   "node.log",
		Level:  "error",
		Period: "hour",
	}
}
func InitLog() {
	reloadLogConfig()
}
func reloadLogConfig() {
	reloadAppLog()
	reloadAccessLog()

}
func reloadAppLog() {
	c, e := dao.Get(NodeLog)
	if e != nil {
		log.Warn("manager/config load goku_config_log fro node  error:", e)
		c = defaultNodeAppLogConfig()
	}
	period, err := log.ParsePeriod(c.Period)
	if err != nil {
		period = log.PeriodDay
		log.Warn("manager/config unmarshal access log period failed for nod , use the default config:%s", e)
	}

	level, err := log.ParseLevel(c.Level)
	if err != nil {
		level = log.WarnLevel
		log.Warn("manager/config unmarshal access log level failed for nod , use the default config:%s", e)
	}

	enable := c.Enable == 1
	log.SetOutPut(enable, c.Dir, c.File, period, c.Expire)
	log.SetLevel(level)

}
func reloadAccessLog() {
	c, e := dao.Get(AccessLog)
	if e != nil {
		log.Warn("manager/config load  goku_config_log for access log error:", e)
		c = defaultAccessLogConfig()
	}

	period, err := log.ParsePeriod(c.Period)
	if err != nil {
		period = log.PeriodDay
		log.Warn("manager/config unmarshal period failed for , use the default config:%s", e)
	}
	enable := c.Enable == 1

	fieldsConfig := make([]AccessField, 0, access_field.Size())
	err = json.Unmarshal([]byte(c.Fields), &fieldsConfig)

	if err != nil || len(fieldsConfig) == 0 {
		log.Warn("manager/config unmarshal access log fields error:", err)

		access_log.SetFields(access_field.Default())
	} else {
		fields := make([]access_field.AccessFieldKey, 0, access_field.Size())
		for _, f := range fieldsConfig {
			if f.Select {
				if access_field.Has(f.Name) {
					fields = append(fields, access_field.Parse(f.Name))
				}
			}
		}
		access_log.SetFields(fields)
	}

	access_log.SetOutput(enable, c.Dir, c.File, period, c.Expire)

}

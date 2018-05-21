package dao

import (
	"goku-ce/server/conf"
	"goku-ce/utils"
	"gopkg.in/yaml.v2"
)

// 编辑全局配置
func EditGlobalConfig(gatewayPort string) bool {
	conf.GlobalConf.Port = gatewayPort
	globalConf,_ := yaml.Marshal(conf.GlobalConf)
	conf.WriteConfigToFile(utils.ConfFilepath,globalConf)
	return true
}
package conf

import (
	_ "goku-ce/utils"
	"goku-ce/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type GlobalConfig struct {
	Host					string					`json:"host" yaml:"host"`
	Port					string					`json:"port" yaml:"port"`
	GatewayConfPath			string					`json:"gatewayConfPath" yaml:"gateway_conf_path"`
	LoginName				string					`json:"loginName" yaml:"login_name"`
	LoginPassword			string					`json:"loginPassword" yaml:"login_password"`
}


var GlobalConf GlobalConfig
func init() {
	ParseAdminConfig()
}

// 读入管理员配置信息
func ParseAdminConfig() {
	err := yaml.Unmarshal([]byte(conf.Configure),&GlobalConf)
	if err != nil {
		panic("Error Global Config!")
	}
}


// 将内容写入文件中
func WriteConfigToFile(path string,content []byte) bool {
	err := ioutil.WriteFile(path, content, 0666)
	if err != nil {
		panic(err)
	}
	return true
}


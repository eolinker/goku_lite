package dao

import (
	"gopkg.in/yaml.v2"
	"goku-ce/server/conf"
	// "gouku-ce/utils"
)

// 安装
func Install(port,loginName,loginPassword,path string) bool {
	config := &conf.GlobalConfig{
		Port:port,
		LoginName:loginName,
		LoginPassword: loginPassword,
		GatewayConfPath:path,
	}
	content, err := yaml.Marshal(config)
	if err != nil {
		return false
	}
	// 判断网关路径是否存在，若不存在，则创建该文件夹
	// utils.CheckFileis
	conf.WriteConfigToFile("./config/goku.conf",content)
	return true
}
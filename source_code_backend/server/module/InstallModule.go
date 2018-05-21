package module

import (
	"goku-ce/server/dao"
)

// 安装
func Install(port,loginName,loginPassword,gatewayConfPath string) bool {
	return dao.Install(port,loginName,loginPassword,gatewayConfPath)
}


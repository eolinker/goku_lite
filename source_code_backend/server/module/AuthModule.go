package module

import (
	"goku-ce/server/dao"
)

// 编辑鉴权信息
func EditAuth(gatewayAlias,strategyID,auth,basicUserName,basicUserPassword,apiKey string) bool {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.EditAuth(gateway["strategyConfPath"],strategyID,auth,basicUserName,basicUserPassword,apiKey)
	}else {
		return false
	}
}

// 获取鉴权信息
func GetAuthInfo(gatewayAlias,strategyID string) map[string]string {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetAuthInfo(gateway["strategyConfPath"],strategyID)
	}else {
		return make(map[string]string)
	}
}
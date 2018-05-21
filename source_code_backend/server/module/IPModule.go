package module

import (
	"goku-ce/server/dao"
)

// 修改网关黑白名单
func EditGatewayIPList(gatewayAlias,ipLimitType,ipWhiteList,ipBlackList string) bool {
	return dao.EditGatewayIPList(gatewayAlias,ipLimitType,ipWhiteList,ipBlackList)
}

// 修改策略组黑白名单
func EditStrategyIPList(gatewayAlias,strategyID,ipLimitType,ipWhiteList,ipBlackList string) bool {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.EditStrategyIPList(gateway["strategyConfPath"],strategyID,ipLimitType,ipWhiteList,ipBlackList)
	}else {
		return false
	}
}

// 获取网关黑白名单
func GetGatewayIPList(gatewayAlias string) map[string]string {
	return dao.GetGatewayIPList(gatewayAlias)
}

// 获取策略组黑白名单
func GetStrategyIPList(gatewayAlias,strategyID string) map[string]string {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetStrategyIPList(gateway["strategyConfPath"],strategyID)
	}else {
		return make(map[string]string)
	}
}

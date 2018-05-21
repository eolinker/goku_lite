package module

import (
	"goku-ce/server/dao"
)

// 新增策略组
func AddStrategy(gatewayAlias,strategyName string) (bool,string) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.AddStrategy(gateway["strategyConfPath"],strategyName)
	}else {
		return false,""
	}
}

// 修改策略组
func EditStrategy(gatewayAlias,strategyName,strategyID string) (bool) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.EditStrategy(gateway["strategyConfPath"],strategyName,strategyID)
	}else {
		return false
	}
}


// 删除策略组
func DeleteStrategy(gatewayAlias,strategyID string) (bool) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		
		return dao.DeleteStrategy(gateway["strategyConfPath"],strategyID)
	}else {
		return false
	}
}


// 获取策略组列表
func GetStrategyList(gatewayAlias string) []map[string]interface{} {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetStrategyList(gateway["strategyConfPath"])
	}else {
		return make([]map[string]interface{},0)
	}
}

// 获取策略组列表
func GetSimpleStrategyList(gatewayAlias string) []map[string]interface{} {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetSimpleStrategyList(gateway["strategyConfPath"])
	}else {
		return make([]map[string]interface{},0)
	}
}

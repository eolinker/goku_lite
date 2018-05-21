package module

import (
	"goku-ce/server/dao"
	"goku-ce/server/conf"
)

// 新增流量限制
func AddRateLimit(gatewayAlias,strategyID,period string,startTime,endTime,priority,limitCount int,allow bool) bool {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.AddRateLimit(gateway["strategyConfPath"],strategyID,period,startTime,endTime,priority,limitCount,allow)
	}else {
		return false
	}
}

// 修改流量限制
func EditRateLimit(gatewayAlias,strategyID,period string,rateLimitID,startTime,endTime,priority,limitCount int,allow bool) bool {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.EditRateLimit(gateway["strategyConfPath"],strategyID,period,rateLimitID,startTime,endTime,priority,limitCount,allow)
	}else {
		return false
	}
}

// 删除流量限制
func DeleteRateLimit(gatewayAlias,strategyID string,rateLimitID int) bool {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.DeleteRateLimit(gateway["strategyConfPath"],strategyID,rateLimitID)
	}else {
		return false
	}
}

// 获取流量限制列表
func GetRateLimitInfo(gatewayAlias,strategyID string,limitID int) (bool,*conf.RateLimitInfo) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetRateLimitInfo(gateway["strategyConfPath"],strategyID,limitID)
	}else {
		return false,&conf.RateLimitInfo{}
	}
}


// 获取流量限制列表
func GetRateLimitList(gatewayAlias,strategyID string) []map[string]interface{}{
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetRateLimitList(gateway["strategyConfPath"],strategyID)
	}else {
		return make([]map[string]interface{},0)
	}
}
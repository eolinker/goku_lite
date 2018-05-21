package module

import (
	"goku-ce/server/dao"
)

// 新增网关
func AddGateway(gatewayName,gatewayAlias string) (bool) {
	return dao.AddGateway(gatewayName,gatewayAlias)
}

// 修改网关信息
func EditGateway(gatewayName,gatewayAlias,oldGatewayAlias string) bool {
	return dao.EditGateway(gatewayName,gatewayAlias,oldGatewayAlias)
}

// 删除网关
func DeleteGateway(gatewayAlias string) bool {
	return dao.DeleteGateway(gatewayAlias)
}

//获取网关列表
func GetGatewayList() (bool,[]map[string]interface{}) {
	return dao.GetGatewayList()
}


//获取网关信息
func GetGatewayInfo(gatewayAlias string) (bool,map[string]interface{}) {
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		flag,result := dao.GetGatewayInfo(gatewayAlias)
		if flag {
			result["apiCount"] = dao.GetApiCount(gateway["apiConfPath"])
			result["strategyCount"] = dao.GetStrategyCount(gateway["strategyConfPath"])
			result["apiGroupCount"] = dao.GetApiGroupCount(gateway["apiGroupConfPath"])
			return true,result
		}
	}
	return false,make(map[string]interface{})
	
}

// 检查网关别名是否存在
func CheckGatewayAliasIsExist(gatewayAlias string) bool {
	return dao.CheckGatewayAliasIsExist(gatewayAlias)
}

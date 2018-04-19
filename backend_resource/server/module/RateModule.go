package module

import (
	"goku-ce-1.0/server/dao"
)

// 新增流量控制
func AddRateLimit(viewType,strategyID,intervalType,limitCount,priorityLevel int,gatewayHashKey,startTime,endTime string) (bool,int){
	if flag := dao.CheckStrategyIsExist(strategyID);flag{
		return dao.AddRateLimit(viewType,strategyID,intervalType,limitCount,priorityLevel,gatewayHashKey,startTime,endTime)
	}else{
		return false,0
	}
	
}

// 编辑流量控制
func EditRateLimit(strategyID,limitID,viewType,intervalType,limitCount,priorityLevel int,gatewayHashKey,startTime,endTime string) (bool){
	if flag := dao.CheckRateIsInStrategy(strategyID,limitID);flag{
		return dao.EditRateLimit(strategyID,limitID,viewType,intervalType,limitCount,priorityLevel,gatewayHashKey,startTime,endTime)
	}else{
		return false
	}
}

// 删除流量控制
func DeleteRateLimit(strategyID,limitID int,gatewayHashKey string) bool{
	if flag := dao.CheckRateIsInStrategy(strategyID,limitID);flag{
		return dao.DeleteRateLimit(strategyID,limitID,gatewayHashKey)
	}else{
		return false
	}
}

// 获取流量控制信息
func GetRateLimitInfo(limitID int) (bool,interface{}){
	return dao.GetRateLimitInfo(limitID)
}

// 获取流量控制列表
func GetRateLimitList(strategyID int) (bool,interface{}){
	return dao.GetRateLimitList(strategyID)
}
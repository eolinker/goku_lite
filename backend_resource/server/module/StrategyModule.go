package module

import (
	"goku-ce-1.0/server/dao"
)

// 新增策略
func AddStrategy(strategyName,strategyDesc string,gatewayID int) (bool,int,string){
	return dao.AddStrategy(strategyName,strategyDesc,gatewayID)
}

// 修改策略
func EditStrategy(strategyName,strategyDesc string,gatewayID,strategyID int) (bool){
	flag := dao.CheckStrategyPermission(gatewayID,strategyID)
	if flag{
		return dao.EditStrategy(strategyName,strategyDesc,strategyID)
	}else{
		return false
	}
}

// 删除策略
func DeleteStrategy(gatewayID,strategyID int) (bool){
	flag := dao.CheckStrategyPermission(gatewayID,strategyID)
	if flag{
		return dao.DeleteStrategy(strategyID)
	}else{
		return false
	}
}

// 获取策略列表
func GetStrategyList(gatewayID int) (bool,interface{}){
	return dao.GetStrategyList(gatewayID)
}

// 查询操作权限
func CheckStrategyPermission(gatewayID,strategyID int) bool{
	return dao.CheckStrategyPermission(gatewayID,strategyID)
}

// 获取简易策略组列表
func GetSimpleStrategyList(gatewayID int) (bool,interface{}){
	return dao.GetSimpleStrategyList(gatewayID)
}
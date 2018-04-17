package module

import (
	"goku-ce-1.0/server/dao"
)

func EditIPList(strategyID,ipType int,gatewayHashKey,ipList string) bool{
	if flag := dao.CheckStrategyIsExist(strategyID);flag{
		return dao.EditIPList(strategyID,ipType,gatewayHashKey,ipList)
	}else{
		return false
	}
	
}

// 获取IP名单列表
func GetIPList(strategyID int) (bool,interface{}){
	if flag := dao.CheckStrategyIsExist(strategyID);flag{
		return dao.GetIPList(strategyID)
	}else{
		return false,nil
	}
}


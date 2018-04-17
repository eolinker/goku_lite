package module

import (
	"goku-ce-1.0/server/dao"
)

func EditAuthMethod(authType,strategyID int,gatewayHashKey,apiKey,userName,userPassword string) bool{
	if flag := dao.CheckStrategyIsExist(strategyID);flag{
		flag = dao.CheckAuthIsExist(strategyID)
		if flag{
			return dao.EditAuthMethod(authType,strategyID,gatewayHashKey,apiKey,userName,userPassword)
		}else{
			return dao.AddAuthMethod(authType,strategyID,apiKey,userName,userPassword)
		}
	} else {
		return false
	}
}


func GetAuthInfo(strategyID int) (bool,interface{}){
	return dao.GetAuthInfo(strategyID)
}
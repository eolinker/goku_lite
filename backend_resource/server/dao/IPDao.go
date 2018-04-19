package dao

import (
	"goku-ce-1.0/dao/database"	
	"goku-ce-1.0/utils"
)

func EditIPList(strategyID,ipType int,gatewayHashKey,ipList string) bool{
	db := database.GetConnection()
	chooseType := "blackList"
	if ipType == 2{
		chooseType = "whiteList"
	}
	if ipType == 0{
		stmt,err := db.Prepare("UPDATE eo_gateway_strategy_group SET chooseType = 0 WHERE strategyID =?;")
		if err != nil{
			return false
		}
		_,err = stmt.Exec(strategyID)
		if err != nil{
			return false
		}
	}else{
		stmt,err := db.Prepare("UPDATE eo_gateway_strategy_group SET " + chooseType + " = ?,chooseType = ? WHERE strategyID =?;")
		if err != nil{
			return false
		}
		_,err = stmt.Exec(ipList,ipType,strategyID)
		if err != nil{
			return false
		}
	}

	// 获取strategyKey
	var strategyKey string
	err := db.QueryRow("SELECT strategyKey FROM eo_gateway_strategy_group WHERE strategyID = ?;",strategyID).Scan(&strategyKey)
	if err != nil{
		return false
	}

	redisConn,err := utils.GetRedisConnection()
	defer redisConn.Close()
	
	_, err = redisConn.Do("del", "IPList:"+ gatewayHashKey + ":" + strategyKey)  
	if err != nil{
		return false
	}
	return true
}

// 获取IP名单列表
func GetIPList(strategyID int) (bool,utils.IPList){
	db := database.GetConnection()
	var ipList utils.IPList
	sql := "SELECT IFNULL(blackList,''),IFNULL(whiteList,''),chooseType FROM eo_gateway_strategy_group WHERE strategyID = ?;"
	err := db.QueryRow(sql,strategyID).Scan(&ipList.BlackList,&ipList.WhiteList,&ipList.ChooseType)
	if err != nil {
		return false,ipList
	}else{
		return true,ipList
	}
}
 
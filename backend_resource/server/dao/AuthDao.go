package dao

import (
	"goku-ce-1.0/dao/database"	
	"goku-ce-1.0/utils"
)

func AddAuthMethod(authType,strategyID int,apiKey,userName,userPassword string) bool{
	db := database.GetConnection()
	sql := "INSERT INTO eo_gateway_auth (authType,strategyID,apiKey,userName,userPassword) VALUES(?,?,?,?,?);"
	stmt,err := db.Prepare(sql)
	if err != nil{
		return false
	}
	defer stmt.Close()
	_,err = stmt.Exec(authType,strategyID,apiKey,userName,userPassword)
	if err != nil{
		return false
	}else{
		return true
	}
}

func EditAuthMethod(authType,strategyID int,gatewayHashKey,apiKey,userName,userPassword string) bool{
	db := database.GetConnection()
	sql := "UPDATE eo_gateway_auth SET authType = ?,apiKey = ?,userName = ?,userPassword = ? WHERE strategyID = ?;"
	stmt,err := db.Prepare(sql)
	if err != nil{
		return false
	}
	defer stmt.Close()
	_,err = stmt.Exec(authType,apiKey,userName,userPassword,strategyID)
	if err != nil{
		return false
	}else{
		var strategyKey string
		err = db.QueryRow("SELECT strategyKey FROM eo_gateway_strategy_group WHERE strategyID = ?;",strategyID).Scan(&strategyKey)
		if err != nil{
			return false
		}
		redisConn,err := utils.GetRedisConnection()
		defer redisConn.Close()
		
		_, err = redisConn.Do("del", "authInfo:"+ gatewayHashKey + ":" + strategyKey)  
		if err != nil{
			return false
		}
		return true
	}
}

func DeleteAuth(strategyID int) (bool){
	db := database.GetConnection()
	sql := "DELETE FROM eo_gateway_auth WHERE strategyID = ?;"
	stmt,err := db.Prepare(sql)
	if err != nil{
		return false
	}
	defer stmt.Close()
	res,err := stmt.Exec(strategyID)
	if err != nil{
		return false
	}else{
		if rowAffect,_:=res.RowsAffected(); rowAffect > 0{
			return true
		}else{
			return false
		}
	}
}

func GetAuthInfo(strategyID int) (bool,utils.AuthInfo){
	db := database.GetConnection()
	var authInfo utils.AuthInfo
	sql := "SELECT authType,apiKey,userName,userPassword FROM eo_gateway_auth WHERE strategyID = ?;"
	err := db.QueryRow(sql,strategyID).Scan(&authInfo.AuthType,&authInfo.ApiKey,&authInfo.UserName,&authInfo.UserPassword)
	if err != nil{
		return false,authInfo
	}else{
		return true,authInfo
	}
}

func CheckAuthIsExist(strategyID int) (bool){
	db := database.GetConnection()
	sql := "SELECT strategyID FROM eo_gateway_auth WHERE strategyID = ?;"
	err := db.QueryRow(sql,strategyID).Scan(&strategyID)
	if err != nil{
		return false
	}else{
		return true
	}
}


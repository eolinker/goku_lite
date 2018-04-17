package dao

import (
	"goku-ce-1.0/dao/database"	
	"goku-ce-1.0/utils"
	"encoding/json"
)

// 获取环境列表
func GetBackendList(gatewayID int) (bool,[]*utils.BackendInfo){
	db := database.GetConnection()
	rows,err := db.Query(`SELECT eo_gateway_backend.backendID,eo_gateway_backend.backendName,eo_gateway_backend.backendURI FROM eo_gateway_backend WHERE eo_gateway_backend.gatewayID = ? ORDER BY eo_gateway_backend.backendID DESC;`,gatewayID)
	
	backendList := make([]*utils.BackendInfo,0)
	flag := true
	if err != nil {
		flag = false
	}
	num :=0
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
    	return false,backendList
	} else {
		for rows.Next(){
			var backend utils.BackendInfo

			err:= rows.Scan(&backend.BackendID,&backend.BackendName,&backend.BackendURI)
			if err!=nil{
				flag = false
				break
			}
			backendList = append(backendList,&backend)
			num +=1
		}
	}
	if num == 0{
		flag =false
	}
	
	return flag,backendList
}

// 添加环境
func AddBackend(gatewayID int,backendName ,backendURI string) (bool,int){
	db := database.GetConnection()
	stmt,err := db.Prepare(`INSERT INTO eo_gateway_backend (eo_gateway_backend.backendName,eo_gateway_backend.backendURI,eo_gateway_backend.gatewayID) VALUES (?,?,?);`)
	defer stmt.Close()
	if err != nil {
		return false,0
	} 
	
	res, err := stmt.Exec(backendName,backendURI,gatewayID)
	if err != nil {
		return false,0
	} else{
		id, _ := res.LastInsertId()
		return true,int(id)
	}

}

func DeleteBackend(gatewayID,backendID int) bool{
	db := database.GetConnection()
	stmt,err := db.Prepare(`DELETE FROM eo_gateway_backend WHERE eo_gateway_backend.backendID = ? AND eo_gateway_backend.gatewayID = ?;`)
	defer stmt.Close()
	if err != nil {
		return false
	} 
	
	res, err := stmt.Exec(backendID,gatewayID)
	if err != nil {
		return false
	} else{
		if rowAffect,_:=res.RowsAffected(); rowAffect > 0{
			return true
		}else{
			return false
		}
	}
}

func EditBackend(backendID,gatewayID int,backendName,backendURI,gatewayHashKey string) bool{
	db := database.GetConnection()
	Tx,_ := db.Begin()
	stmt,err := Tx.Prepare(`UPDATE eo_gateway_backend SET eo_gateway_backend.backendName = ?,eo_gateway_backend.backendURI = ? WHERE eo_gateway_backend.backendID = ?;`)
	if err != nil {
		Tx.Rollback()
		return false
	} 
	_, err = stmt.Exec(backendName,backendURI,backendID)
	if err != nil {
		Tx.Rollback()
		return false
	} else{
		Tx.Exec("UPDATE eo_gateway_api SET backendURI = ? WHERE backendID = ?",backendURI,backendID)
		Tx.Exec("UPDATE eo_gateway_api_cache SET path = ? WHERE backendID = ?",backendURI,backendID)
		
		redisConn,err := utils.GetRedisConnection()
		defer redisConn.Close()
		if err != nil{
			Tx.Rollback()
			return false
		}
		var queryJson utils.QueryJson
		var operationData utils.OperationData
		operationData.GatewayID = gatewayID
		operationData.GatewayHashKey = gatewayHashKey
		queryJson.OperationType = "backend"
		queryJson.Operation = "update"
		queryJson.Data = operationData
		redisString,_ := json.Marshal(queryJson)
		_, err = redisConn.Do("rpush", "gatewayQueue", string(redisString[:]))  
		if err != nil{
			Tx.Rollback()
			return false
		}
		Tx.Commit()
		return true
	}
}

// 获取网关信息
func GetBackendInfo(backendID int) (bool,utils.BackendInfo){
	db := database.GetConnection()
	var backendInfo utils.BackendInfo
	err := db.QueryRow("SELECT eo_gateway_backend.backendID,eo_gateway_backend.backendName,eo_gateway_backend.backendURI FROM eo_gateway_backend WHERE eo_gateway_backend.backendID = ?;",backendID).Scan(&backendInfo.BackendID,&backendInfo.BackendName,&backendInfo.BackendURI)
	if err != nil {
		return false,backendInfo
	} else{
		return true,backendInfo
	}
}
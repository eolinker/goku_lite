package dao

import (
	"goku-ce-1.0/dao/database"
	"goku-ce-1.0/utils"
	"goku-ce-1.0/conf"
	"time"
)

// 新增策略
func AddStrategy(strategyName,strategyDesc string,gatewayID int) (bool,int,string){
	db := database.GetConnection()
	// 随机生成字符串
	randomStr := utils.GetRandomString(6)
	sqlCode := "SELECT gatewayID FROM eo_gateway_strategy_group WHERE strategyKey = ?"
	var strategyKey string
	for i:=0;i<5;i++{
		err := db.QueryRow(sqlCode,randomStr).Scan(strategyKey)
		if err != nil{
			break
		}else{
			continue
		}
	}


	createTime := time.Now().Format("2006-01-02 15:04:05")
	sql := "INSERT INTO eo_gateway_strategy_group (strategyName,strategyDesc,updateTime,createTime,gatewayID,strategyKey) VALUES(?,?,?,?,?,?);"
	stmt,err := db.Prepare(sql)
	if err != nil {
		return false,0,""
	} 
	defer stmt.Close()
	res, err := stmt.Exec(strategyName,strategyDesc,createTime,createTime,gatewayID,randomStr)
	if err != nil {
		return false,0,""
	} else{
		id, _ := res.LastInsertId()
		var strategyKey string
		sql = "SELECT strategyKey FROM eo_gateway_strategy_group WHERE strategyID =?"
		err = db.QueryRow(sql,id).Scan(&strategyKey)
		if err != nil{
			return false,0,""
		}
		return true,int(id),strategyKey
	}
}

// 修改策略
func EditStrategy(strategyName,strategyDesc string,strategyID int) (bool){
	db := database.GetConnection()
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	sql := "UPDATE eo_gateway_strategy_group SET strategyName = ?,strategyDesc = ?,updateTime = ? WHERE strategyID = ?;"
	stmt,err := db.Prepare(sql)
	if err != nil {
		return false
	} 
	defer stmt.Close()
	_, err = stmt.Exec(strategyName,strategyDesc,updateTime,strategyID)
	if err != nil {
		return false
	} else{
		return true
	}
}

// 删除策略
func DeleteStrategy(strategyID int) (bool){
	db := database.GetConnection()
	stmt,err := db.Prepare(`DELETE FROM eo_gateway_strategy_group WHERE eo_gateway_strategy_group.strategyID = ?`)
	defer stmt.Close()
	if err != nil {
		return false
	} 
	
	_, err = stmt.Exec(strategyID)
	if err != nil {
		return false
	} else{
		return true
	}
}

// 获取策略列表
func GetStrategyList(gatewayID int) (bool,interface{}){
	db := database.GetConnection()
	
	sql := "SELECT strategyID,strategyName,strategyDesc,strategyKey,eo_gateway_strategy_group.updateTime,eo_gateway.gatewayAlias FROM eo_gateway_strategy_group INNER JOIN eo_gateway ON  eo_gateway.gatewayID = eo_gateway_strategy_group.gatewayID WHERE eo_gateway_strategy_group.gatewayID = ? ORDER BY updateTime DESC"
	rows,err := db.Query(sql,gatewayID)
	if err != nil {
		return false,nil
	}
	num :=0
	strategyList := make([]map[string]interface{},0)
	gatewayPort := conf.Configure["eotest_port"]
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
    	return false,nil
	} else {
		for rows.Next(){
			var strategyID int
			var strategyName,strategyDesc,updateTime,gatewayAlias,strategyKey string
			err = rows.Scan(&strategyID,&strategyName,&strategyDesc,&strategyKey,&updateTime,&gatewayAlias);
			if err!=nil{
				return false,nil
			}
			strategyList = append(strategyList,map[string]interface{}{"strategyID":strategyID,"strategyName":strategyName,"strategyDesc":strategyDesc,"strategyKey":strategyKey,"updateTime":updateTime,"gatewayAlias":gatewayAlias,"gatewayPort":gatewayPort})
			num +=1
		}
	}
	if num == 0{
		return false,nil
	}
	return true,strategyList
}

// 查询操作权限
func CheckStrategyPermission(gatewayID,strategyID int) bool{
	db := database.GetConnection()
	err := db.QueryRow("SELECT gatewayID FROM eo_gateway_strategy_group WHERE gatewayID = ? AND strategyID = ?;",gatewayID,strategyID).Scan(&gatewayID)
	if err != nil {
		return false
	} else{
		return true
	}
}

// 检查策略组是否存在
func CheckStrategyIsExist(strategyID int) bool {
	db := database.GetConnection()
	err := db.QueryRow("SELECT strategyID FROM eo_gateway_strategy_group WHERE strategyID = ?;",strategyID).Scan(&strategyID)
	if err != nil {
		return false
	} else{
		return true
	}
}

// 获取策略组数量
func GetStrategyCount(gatewayHashKey string) int{
	var count int
	db := database.GetConnection()
	sql := "SELECT COUNT(0) FROM eo_gateway_strategy_group INNER JOIN eo_gateway ON eo_gateway_strategy_group.gatewayID = eo_gateway.gatewayID WHERE eo_gateway.hashKey = ?;"
	db.QueryRow(sql,gatewayHashKey).Scan(&count)
	return count
}

// 获取简易策略组列表
func GetSimpleStrategyList(gatewayID int) (bool,interface{}){
	db := database.GetConnection()
	
	sql := "SELECT strategyID,strategyName,strategyKey FROM eo_gateway_strategy_group WHERE eo_gateway_strategy_group.gatewayID = ? ORDER BY updateTime DESC"
	rows,err := db.Query(sql,gatewayID)
	if err != nil {
		return false,nil
	}
	num :=0
	strategyList := make([]map[string]interface{},0)
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
    	return false,nil
	} else {
		for rows.Next(){
			var strategyID int
			var strategyName,strategyKey string
			err = rows.Scan(&strategyID,&strategyName,&strategyKey)
			if err!=nil{
				return false,nil
			}
			strategyList = append(strategyList,map[string]interface{}{"strategyID":strategyID,"strategyName":strategyName,"strategyKey":strategyKey})
			num +=1
		}
	}
	if num == 0{
		return false,nil
	}
	return true,strategyList
}
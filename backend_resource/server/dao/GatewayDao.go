package dao

import (
	"goku-ce-1.0/dao/database"	
	"goku-ce-1.0/utils"
	"encoding/json"
	"time"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"goku-ce-1.0/conf"
)

// 新增网关
func Addgateway(gatewayName,gatewayDesc,gatewayAlias,createTime,hashKey string,userID int) (bool,int){
	db := database.GetConnection()
	Tx,_ :=db.Begin()
	stmt,err := Tx.Prepare(`INSERT INTO eo_gateway (eo_gateway.gatewayName,eo_gateway.gatewayDesc,eo_gateway.gatewayAlias,eo_gateway.hashKey,eo_gateway.createTime,eo_gateway.updateTime) VALUES (?,?,?,?,?,?);`)
	if err != nil {
		Tx.Rollback()
		return false,0
	} 
	defer stmt.Close()
	
	res, err := stmt.Exec(gatewayName, gatewayDesc,gatewayAlias,hashKey,createTime,createTime)
	if err != nil {
		Tx.Rollback()
		return false,0
	} else{
		id, _ := res.LastInsertId()
		stmt ,err = Tx.Prepare("INSERT INTO eo_conn_gateway (eo_conn_gateway.gatewayID,eo_conn_gateway.userID) VALUES (?,?);")
		if err != nil {
			Tx.Rollback()
			return false,0
		} 
		_,err = stmt.Exec(int(id),userID)
		if err != nil {
			Tx.Rollback()
			return false,0
		} 
		redisConn,err := utils.GetRedisConnection()
		defer redisConn.Close()
		var queryJson utils.QueryJson
		var operationData utils.OperationData
		
		operationData.GatewayAlias = gatewayAlias
		operationData.GatewayID = int(id)
		operationData.GatewayHashKey = hashKey
		queryJson.OperationType = "gateway"
		queryJson.Operation = "add"
		queryJson.Data = operationData
		redisString,_ := json.Marshal(queryJson)
		_, err = redisConn.Do("rpush", "gatewayQueue", string(redisString[:]))  
		if err != nil{
			Tx.Rollback()
			return false,0
		}
		Tx.Commit()
		return true,int(id)
	}
}

// 修改网关
func EditGateway(gatewayName,gatewayAlias,gatewayDesc,gatewayHashKey string) bool{
	db := database.GetConnection()
	flag,oldAlias := GetGatewayAlias(gatewayHashKey)
	if flag{
		stmt,err := db.Prepare(`UPDATE eo_gateway SET gatewayName = ?,gatewayAlias = ?,gatewayDesc = ? WHERE hashKey = ?;`)
		defer stmt.Close()
		if err != nil {
			return false
		} 
		
		_,err = stmt.Exec(gatewayName,gatewayAlias,gatewayDesc,gatewayHashKey)
		if err != nil {
			return false
		} else{
			redisConn,err := utils.GetRedisConnection()
			defer redisConn.Close()
			var queryJson utils.QueryJson
			var operationData utils.OperationData
			operationData.GatewayHashKey = gatewayHashKey
			operationData.GatewayAlias = oldAlias
			queryJson.OperationType = "gateway"
			queryJson.Operation = "delete"
			queryJson.Data = operationData
			redisString,_ := json.Marshal(queryJson)
			_, err = redisConn.Do("rpush", "gatewayQueue", string(redisString[:]))  
			if err != nil{
				return false
			}
			return true
		}
	}else{
		return false
	}
}

// 删除网关
func DeleteGateway(gatewayHashkey string) bool{
	db := database.GetConnection()
	flag,id :=GetIDFromHashKey(gatewayHashkey)
	if flag{
		flag,gatewayAlias := GetGatewayAlias(gatewayHashkey)
		if flag{
			stmt,err := db.Prepare(`DELETE FROM eo_gateway WHERE hashKey = ?;`)
			defer stmt.Close()
			if err != nil {
				return false
			} 
			_,err = stmt.Exec(gatewayHashkey)
			if err != nil {
				return false
			} else{
				redisConn,err := utils.GetRedisConnection()
				defer redisConn.Close()
				var queryJson utils.QueryJson
				var operationData utils.OperationData
				operationData.GatewayID = int(id)
				operationData.GatewayHashKey = gatewayHashkey
				operationData.GatewayAlias = gatewayAlias
				queryJson.OperationType = "gateway"
				queryJson.Operation = "delete"
				queryJson.Data = operationData
				redisString,_ := json.Marshal(queryJson)
				_, err = redisConn.Do("rpush", "gatewayQueue", string(redisString[:]))  
				if err != nil{
					return false
				}
				return true
			}
		}else{
			return false
		}
	}else{
		return false
	}

}

// 从hashKey获取ID
func GetIDFromHashKey(gatewayHashKey string) (bool,int){
	db := database.GetConnection()
	gatewayID := 0
	sql := `SELECT eo_gateway.gatewayID FROM eo_gateway WHERE eo_gateway.hashKey = ?;`
	err := db.QueryRow(sql,gatewayHashKey).Scan(&gatewayID)
	flag := true
	if err != nil{
		flag = false
	}
	return flag,gatewayID
}

/**
 * 判断网关和用户是否匹配
 * @param $gateway_id 网关数字ID
 * @param $user_id 用户数字ID
 */
func CheckGatewayPermission(gatewayID ,userID int) bool{
	db :=database.GetConnection()
	sql := `SELECT eo_conn_gateway.gatewayID FROM eo_conn_gateway WHERE eo_conn_gateway.gatewayID = ? AND eo_conn_gateway.userID = ?;`
	err := db.QueryRow(sql,gatewayID,userID).Scan(&gatewayID)
	if err != nil{
		return true
	}else{
		return true
	}

}

// 获取网关列表
func GetGatewayList(userID int) (bool,[]*utils.GatewayInfo){
	db := database.GetConnection()
	var err error
	rows,err := db.Query(`SELECT eo_gateway.gatewayID,eo_gateway.gatewayName,eo_gateway.gatewayAlias,eo_gateway.gatewayStatus,eo_gateway.productType,eo_gateway.gatewayDesc,eo_gateway.updateTime,eo_gateway.hashKey AS gatewayHashKey FROM eo_gateway INNER JOIN eo_conn_gateway ON eo_gateway.gatewayID = eo_conn_gateway.gatewayID WHERE eo_conn_gateway.userID = ?;`,userID)
	
	gatewayList := make([]*utils.GatewayInfo,0)
	flag := true
	if err != nil {
		flag = false
	}
	num :=0
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
    	return false,gatewayList
	} else {
		for rows.Next(){
			var gateway utils.GatewayInfo

			err = rows.Scan(&gateway.GatewayID,&gateway.GatewayName,&gateway.GatewayAlias,&gateway.GatewayStatus,&gateway.ProductType,&gateway.GatewayDesc,&gateway.UpdateTime,&gateway.GatewayHashKey);
			if err!=nil{
				flag = false
				break
			}
			gatewayList = append(gatewayList,&gateway)
			num +=1
		}
	}
	if num == 0{
		flag =false
	}
	return flag,gatewayList

}

// 获取网关信息
func GetGatewayInfo(gatewayHashKey string) (bool,map[string]interface{}){
	db := database.GetConnection()
	var gatewayID int
	var gatewayName,gatewayDesc,gatewayStatus,updateTime,createTime,gatewayAlias string
	sql := `SELECT eo_gateway.gatewayID,eo_gateway.gatewayName,eo_gateway.gatewayDesc,eo_gateway.gatewayStatus,eo_gateway.updateTime,eo_gateway.createTime,eo_gateway.gatewayAlias FROM eo_gateway WHERE eo_gateway.hashKey = ?;`
	err := db.QueryRow(sql,gatewayHashKey).Scan(&gatewayID,&gatewayName,&gatewayDesc,&gatewayStatus,&updateTime,&createTime,&gatewayAlias)
	if err != nil{
		return false,nil
	}
	redisInfo := GetGatewayInfoByRedis(gatewayHashKey)
	now := time.Now()
	ct,err := time.ParseInLocation("2006-01-02 15:04:05",createTime,time.Local)
	if err != nil{
		return false,nil
	}
	
	subTime := now.Sub(ct)
	minute := int(subTime.Minutes())%60
	hour := int(subTime.Hours())%24
	day := int(subTime.Hours())/24
	
	monitorTime := strconv.Itoa(day) + " 天" + strconv.Itoa(hour) + " 时" + strconv.Itoa(minute) + " 分"
	gatewayPort := conf.Configure["eotest_port"]

	// 获取策略组数量
	apiGroupCount := GetApiGroupCount(gatewayHashKey)
	strategyGroupCount := GetStrategyCount(gatewayHashKey)

	return true,map[string]interface{}{"gatewayID":gatewayID,"gatewayName":gatewayName,"gatewayDesc":gatewayDesc,"gatewayStatus":gatewayStatus,"updateTime":updateTime,"createTime":createTime,"gatewayAlias":gatewayAlias,"gatewayPort":gatewayPort,"monitorTime":monitorTime,"apiGroupCount":apiGroupCount,"strategyGroupCount":strategyGroupCount,"redisInfo":redisInfo}
}

// 通过hashKey获取网关别名
func GetGatewayAlias(gatewayHashKey string) (bool,string){
	db := database.GetConnection()
	var gatewayAlias string
	sql := "SELECT gatewayAlias FROM eo_gateway WHERE hashKey = ?;"
	err := db.QueryRow(sql,gatewayHashKey).Scan(&gatewayAlias)
	if err != nil{
		return false,""
	}else{
		return true,gatewayAlias
	}
}

// 查询网关别名是否存在
func CheckGatewayAliasIsExist(gatewayAlias string) (bool,string){
	db := database.GetConnection()
	var hashKey string
	sql := "SELECT gatewayHashkey FROM eo_gateway WHERE gatewayAlias = ?;"
	err := db.QueryRow(sql,gatewayAlias).Scan(&hashKey)
	if err != nil{
		return false,""
	}else{
		return true,hashKey
	}
}

// 获取网关访问实时信息
func GetGatewayInfoByRedis(gatewayHashKey string) interface{}{
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	timeStr := dateStr + "-" + strconv.Itoa(hour) + "-" + strconv.Itoa(minute)

	redisConn,err := utils.GetRedisConnection()
	defer redisConn.Close()
	redisKey := "gatewayMinuteCount:" + gatewayHashKey + ":" + timeStr
	gatewayMinuteCount, err := redis.String(redisConn.Do("GET",redisKey))
	if err != nil {
        gatewayMinuteCount = "0"
	}
	redisKey = "gatewayDayCount:" + gatewayHashKey+ ":" + dateStr
	gatewayDayCount, err := redis.String(redisConn.Do("GET",redisKey))
	if err != nil {
        gatewayDayCount = "0"
	}
	redisKey = "gatewaySuccessCount:" + gatewayHashKey + ":" + dateStr
	gatewaySuccessCount, err := redis.String(redisConn.Do("GET",redisKey))
	if err != nil {
        gatewaySuccessCount = "0"
	}

	redisKey = "gatewayFailureCount:" + gatewayHashKey + ":" + dateStr
	gatewayFailureCount, err := redis.String(redisConn.Do("GET",redisKey))
	if err != nil {
        gatewayFailureCount = "0"
	}
	hourStr := ""
	if hour < 10 {
		hourStr += "0"
	}
	minuteStr := ""
	if minute < 10{
		minuteStr += "0"
	}
	secondStr := ""
	if second<10{
		secondStr += "0" 
	}
	lastUpdateTime := hourStr + strconv.Itoa(hour) + ":" + minuteStr + strconv.Itoa(minute) + ":" + secondStr + strconv.Itoa(second)
	return map[string]interface{}{"gatewayMinuteCount":gatewayMinuteCount,"gatewayDayCount":gatewayDayCount,"gatewaySuccessCount":gatewaySuccessCount,"gatewayFailureCount":gatewayFailureCount,"lastUpdateTime":lastUpdateTime}
}

// 获取简易网关信息
func GetSimpleGatewayInfo(gatewayHashKey string) (bool,interface{}) {
	db := database.GetConnection()
	var gatewayID int
	var gatewayName,gatewayAlias string
	sql := `SELECT eo_gateway.gatewayID,eo_gateway.gatewayName,eo_gateway.gatewayAlias FROM eo_gateway WHERE eo_gateway.hashKey = ?;`
	err := db.QueryRow(sql,gatewayHashKey).Scan(&gatewayID,&gatewayName,&gatewayAlias)
	if err != nil{
		return false,nil
	}
	gatewayPort := conf.Configure["eotest_port"]

	return true,map[string]interface{}{"gatewayID":gatewayID,"gatewayName":gatewayName,"gatewayAlias":gatewayAlias,"gatewayPort":gatewayPort}
}
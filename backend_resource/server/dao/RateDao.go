package dao

import (
	"goku-ce-1.0/dao/database"	
	"goku-ce-1.0/utils"
	"time"
	"strconv"
	"github.com/garyburd/redigo/redis"
)

// 新增流量控制
func AddRateLimit(viewType,strategyID,intervalType,limitCount,priorityLevel int,gatewayHashKey,startTime,endTime string) (bool,int){
	db := database.GetConnection()
	createTime := time.Now().Format("2006-01-02 15:04:05")
	sql := "INSERT INTO eo_gateway_rate_limit (intervalType,viewType,limitCount,strategyID,priorityLevel,startTime,endTime,createTime,updateTime) VALUES(?,?,?,?,?,?,?,?,?);"
	stmt,err := db.Prepare(sql)
	if err !=nil{
		return false,0
	}
	defer stmt.Close()
	res,err := stmt.Exec(intervalType,viewType,limitCount,strategyID,priorityLevel,startTime,endTime,createTime,createTime)
	if err != nil{
		return false,0
	}else{
		if rowAffect,_:=res.RowsAffected(); rowAffect > 0{
			id, _ := res.LastInsertId()
			var strategyKey string
			err = db.QueryRow("SELECT strategyKey FROM eo_gateway_strategy_group WHERE strategyID = ?;",strategyID).Scan(&strategyKey)
			if err != nil{
				return false,0
			}

			now := time.Now()
			hour := now.Hour()
			var nextHour int
			if hour == 23{
				nextHour = 0
			}else{
				nextHour = hour + 1
			}
			var valveKey string
			if intervalType == 0{
				valveKey = "gatewayStrategySecondValve:"
			}else if intervalType == 1{
				valveKey = "gatewayStrategyMinuteValve:"
			}else if intervalType == 2{
				valveKey = "gatewayStrategyHourValve:"
			}else if intervalType == 3{
				valveKey = "gatewayStrategyDayValve:"
			}
			var redisKey string = valveKey + gatewayHashKey + ":" + strategyKey + ":" + strconv.Itoa(hour) + "-" + strconv.Itoa(nextHour)
			redisConn,err := utils.GetRedisConnection()
			defer redisConn.Close()
			if err != nil{
				return false,0
			}
			redisConn.Do("del", redisKey)  
			return true,int(id)
		}else{
			return false,0
		}
	}
}

// 编辑流量控制
func EditRateLimit(strategyID,limitID,viewType,intervalType,limitCount,priorityLevel int,gatewayHashKey,startTime,endTime string) (bool){
	db := database.GetConnection()
	var oldType int
	err := db.QueryRow("SELECT intervalType FROM eo_gateway_rate_limit WHERE limitID = ?;",limitID).Scan(&oldType)
	if err != nil{
		return false
	}
	updateTime := time.Now().Format("2006-01-02 15:04:05")
	sql := "UPDATE eo_gateway_rate_limit SET intervalType = ?,viewType = ?,limitCount = ?,priorityLevel = ?,startTime = ?,endTime = ?,updateTime = ? WHERE limitID = ?;"
	stmt,err := db.Prepare(sql)
	if err !=nil{
		return false
	}
	defer stmt.Close()
	res,err := stmt.Exec(intervalType,viewType,limitCount,priorityLevel,startTime,endTime,updateTime,limitID)
	if err != nil{
		return false
	}else{
		if rowAffect,_:=res.RowsAffected(); rowAffect > 0{
			var strategyKey string
			err = db.QueryRow("SELECT strategyKey FROM eo_gateway_strategy_group WHERE strategyID = ?;",strategyID).Scan(&strategyKey)
			if err != nil{
				return false
			}

			now := time.Now()
			hour := now.Hour()
			var nextHour int
			if hour == 23{
				nextHour = 0
			}else{
				nextHour = hour + 1
			}
			year, month, day := now.Date()
			dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
			timeStr := dateStr + "-" + strconv.Itoa(now.Hour()) + "-" + strconv.Itoa(now.Minute())
			
			redisConn,err := utils.GetRedisConnection()
			defer redisConn.Close()
			if err != nil{
				return false
			}
			var valveKey string
			var countKey string
			if oldType == 0{
				valveKey = "gatewayStrategySecondValve:"
				countKey = "gatewayStrategy:" +  gatewayHashKey + ":" + strategyKey + ":IPSecondCount:"  + timeStr + "-" + strconv.Itoa(now.Second()) + "*"
			}else if oldType == 1{
				valveKey = "gatewayStrategyMinuteValve:"
				countKey = "gatewayStrategy:" + gatewayHashKey + ":" + strategyKey + ":IPMinuteCount:" + timeStr + "*"
			}else if oldType == 2{
				valveKey = "gatewayStrategyHourValve:"
				countKey = "gatewayStrategy:" + gatewayHashKey + ":" + strategyKey + ":IPHourCount:"  + dateStr + "-" +  strconv.Itoa(now.Hour()) + "*"
			}else if oldType == 3{
				valveKey = "gatewayStrategyDayValve:"
				countKey = "gatewayStrategy:" + gatewayHashKey + ":" + strategyKey + ":IPDayCount:" + dateStr + "*"
			}
			
			var redisKey string = valveKey + gatewayHashKey + ":" + strategyKey + ":" + strconv.Itoa(hour) + "-" + strconv.Itoa(nextHour)
			redisConn.Do("del", redisKey) 
			keys,err := redis.Strings(redisConn.Do("keys",countKey))
			if err != nil{
				panic(err)
			}
			if len(keys) > 0 {
				for _,key := range keys{
					_,err = redisConn.Do("del",key)
					if err != nil{
						panic(err)
					}
				}
			}
			if intervalType == 0{
				valveKey = "gatewayStrategySecondValve:"
				countKey = "gatewayStrategy:" +  gatewayHashKey + ":" + strategyKey + ":IPSecondCount:"  + timeStr + "-" + strconv.Itoa(now.Second()) + "*"
			}else if intervalType == 1{
				valveKey = "gatewayStrategyMinuteValve:"
				countKey = "gatewayStrategy:" + gatewayHashKey + ":" + strategyKey + ":IPMinuteCount:" + timeStr + "*"
			}else if intervalType == 2{
				valveKey = "gatewayStrategyHourValve:"
				countKey = "gatewayStrategy:" + gatewayHashKey + ":" + strategyKey + ":IPHourCount:"  + dateStr + "-" +  strconv.Itoa(now.Hour()) + "*"
			}else if intervalType == 3{
				valveKey = "gatewayStrategyDayValve:"
				countKey = "gatewayStrategy:" + gatewayHashKey + ":" + strategyKey + ":IPDayCount:" + dateStr + "*"
			}
			redisKey = valveKey + gatewayHashKey + ":" + strategyKey + ":" + strconv.Itoa(hour) + "-" + strconv.Itoa(nextHour)
			redisConn.Do("del", redisKey)
			keys,err = redis.Strings(redisConn.Do("keys",countKey))
			if err != nil{
				panic(err)
			}
			if len(keys) > 0 {
				for _,key := range keys{
					_,err = redisConn.Do("del",key)
					if err != nil{
						panic(err)
					}
				}
			}
			return true
		}else{
			return false
		}
	}
}

// 删除流量控制
func DeleteRateLimit(strategyID,limitID int,gatewayHashkey string) bool{
	db := database.GetConnection()
	var oldType int
	err := db.QueryRow("SELECT intervalType FROM eo_gateway_rate_limit WHERE limitID = ?;",limitID).Scan(&oldType)
	if err != nil{
		return false
	}
	sql := "DELETE FROM eo_gateway_rate_limit WHERE limitID = ?;"
	stmt,err := db.Prepare(sql)
	if err !=nil{
		return false
	}
	defer stmt.Close()
	res,err := stmt.Exec(limitID)
	if err != nil{
		return false
	}else{
		if rowAffect,_:=res.RowsAffected(); rowAffect > 0{
			var strategyKey string
			err = db.QueryRow("SELECT strategyKey FROM eo_gateway_strategy_group WHERE strategyID = ?;",strategyID).Scan(&strategyKey)
			if err != nil{
				return false
			}
			now := time.Now()
			hour := now.Hour()
			var nextHour int
			if hour == 23{
				nextHour = 0
			}else{
				nextHour = hour + 1
			}
			var valveKey string
			var countKey string
			year, month, day := now.Date()
			dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
			timeStr := dateStr + "-" + strconv.Itoa(now.Hour()) + "-" + strconv.Itoa(now.Minute())

			if oldType == 0{
				valveKey = "gatewayStrategySecondValve:"
				countKey = "gatewayStrategy:" +  gatewayHashkey + ":" + strategyKey + ":IPSecondCount:"  + timeStr + "-" + strconv.Itoa(now.Second()) + "*"
			}else if oldType == 1{
				valveKey = "gatewayStrategyMinuteValve:"
				countKey = "gatewayStrategy:" + gatewayHashkey + ":" + strategyKey + ":IPMinuteCount:" + timeStr + "*"
			}else if oldType == 2{
				valveKey = "gatewayStrategyHourValve:"
				countKey = "gatewayStrategy:" + gatewayHashkey + ":" + strategyKey + ":IPHourCount:"  + dateStr + "-" +  strconv.Itoa(now.Hour()) + "*"
			}else if oldType == 3{
				valveKey = "gatewayStrategyDayValve:"
				countKey = "gatewayStrategy:" + gatewayHashkey + ":" + strategyKey + ":IPDayCount:" + dateStr + "*"
			}
			
			var redisKey string = valveKey + gatewayHashkey + ":" + strategyKey + ":" + strconv.Itoa(hour) + "-" + strconv.Itoa(nextHour)
			
			redisConn,err := utils.GetRedisConnection()
			defer redisConn.Close()
			if err != nil{
				return false
			}
			keys,err := redis.Strings(redisConn.Do("keys",countKey))
			if err != nil{
				panic(err)
			}
			if len(keys) > 0 {
				for _,key := range keys{
					
					_,err = redisConn.Do("del",key)
					if err != nil{
						panic(err)
					}
				}
			}
			redisConn.Do("del", redisKey) 
			redisConn.Do("del",countKey)
			return true
		}else{
			return false
		}
	}
}

// 获取流量控制信息
func GetRateLimitInfo(limitID int) (bool,utils.RateLimitInfo){
	db := database.GetConnection()
	var rateLimitInfo utils.RateLimitInfo
	sql := "SELECT limitID,viewType,intervalType,limitCount,priorityLevel,startTime,endTime FROM eo_gateway_rate_limit WHERE limitID = ?;"
	err := db.QueryRow(sql,limitID).Scan(&rateLimitInfo.LimitID,&rateLimitInfo.ViewType,&rateLimitInfo.IntervalType,&rateLimitInfo.LimitCount,&rateLimitInfo.PriorityLevel,&rateLimitInfo.StartTime,&rateLimitInfo.EndTime)
	if err != nil {
		return false,rateLimitInfo
	}else{
		return true,rateLimitInfo
	}
}

// 获取流量控制列表
func GetRateLimitList(strategyID int) (bool,[]utils.RateLimitInfo){
	db := database.GetConnection()
	rateLimitList := make([]utils.RateLimitInfo,0)
	sql := "SELECT limitID,viewType,intervalType,limitCount,priorityLevel,startTime,endTime FROM eo_gateway_rate_limit WHERE strategyID = ? ORDER BY priorityLevel DESC,startTime ASC;"
	rows,err := db.Query(sql,strategyID)
	if err != nil {
		return false,rateLimitList
	}
	num := 0
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
    	return false,rateLimitList
	} else {
		for rows.Next(){
			var rateLimitInfo utils.RateLimitInfo

			err:= rows.Scan(&rateLimitInfo.LimitID,&rateLimitInfo.ViewType,&rateLimitInfo.IntervalType,&rateLimitInfo.LimitCount,&rateLimitInfo.PriorityLevel,&rateLimitInfo.StartTime,&rateLimitInfo.EndTime)
			if err!=nil{
				return false,rateLimitList
			}
			rateLimitList = append(rateLimitList,rateLimitInfo)
			num +=1
		}
	}
	if num == 0{
		return false,rateLimitList
	}
	return true,rateLimitList
}

// 检查流量限制是否属于该策略组
func CheckRateIsInStrategy(strategyID,limitID int) (bool){
	db := database.GetConnection()
	sql := "SELECT limitID FROM eo_gateway_rate_limit WHERE strategyID = ? AND limitID = ?;"
	err := db.QueryRow(sql,strategyID,limitID).Scan(&limitID)
	if err != nil {
		return false
	} else{
		return true
	}
}
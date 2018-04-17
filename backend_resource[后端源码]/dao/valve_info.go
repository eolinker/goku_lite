package dao

import (
	"goku-ce-1.0/dao/cache"
	"goku-ce-1.0/utils"
	"strconv"
	"time"
	"goku-ce-1.0/dao/database"
	"github.com/farseer810/yawf"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"fmt"
)

// 获取网关当日访问次数
func GetGatewayDayVisitCount(context yawf.Context, info *utils.MappingInfo) int {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)

	var redisKey string = "gatewayDayCount:" + info.GatewayHashKey + ":" + dateStr
	conn := cache.GetConnection(context)
	count, err := redis.Int(conn.Do("GET", redisKey))

	if err == redis.ErrNil {
		return 0
	} else if err != nil {
		panic(err)
	}
	return count
}

// 获取网关实时访问次数（精确到分钟）
func GetGatewayMinuteCount(context yawf.Context, info *utils.MappingInfo) int {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	timeStr := dateStr + "-" + strconv.Itoa(now.Hour()) + "-" + strconv.Itoa(now.Minute())

	var redisKey string = "gatewayMinuteCount:" + info.GatewayHashKey + ":" + timeStr
	conn := cache.GetConnection(context)
	count, err := redis.Int(conn.Do("GET", redisKey))

	if err == redis.ErrNil {
		return 0
	} else if err != nil {
		panic(err)
	}
	return count
}

// 获取当天网关策略组访问次数
func GetGatewayStrategyDayCount(context yawf.Context, info *utils.MappingInfo) int {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)

	var redisKey string = "gatewayStrategy:" + info.GatewayHashKey  + ":" + info.StrategyKey + ":" + dateStr
	conn := cache.GetConnection(context)
	count, err := redis.Int(conn.Do("GET", redisKey))

	if err == redis.ErrNil {
		return 0
	} else if err != nil {
		panic(err)
	}
	return count
}

// 获取实时网关策略组访问次数
func GetGatewayStrategyMinuteCount(context yawf.Context, info *utils.MappingInfo) int {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	timeStr := dateStr + "-" + strconv.Itoa(now.Hour()) + "-" + strconv.Itoa(now.Minute())
	var redisKey string = "gatewayStrategy:" + info.GatewayHashKey  + ":" + info.StrategyKey + ":" + timeStr
	conn := cache.GetConnection(context)
	count, err := redis.Int(conn.Do("GET", redisKey))

	if err == redis.ErrNil {
		return 0
	} else if err != nil {
		panic(err)
	}
	return count
}

// 获取当天IP访问网关策略次数
func GetGatewayStrategyIPDayCount(context yawf.Context,gatewayHashKey,strategyKey, remoteIP string) int {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)

	var redisKey string = "gatewayStrategy:" + gatewayHashKey  + ":" + strategyKey + ":IPDayCount:"  + dateStr + ":" + remoteIP
	conn := cache.GetConnection(context)
	count, err := redis.Int(conn.Do("GET", redisKey))
	
	if err == redis.ErrNil {
		return -1
	} else if err != nil {
		panic(err)
	}
	return count
}

// 获取当前IP访问网关策略次数（小时）
func GetGatewayStrategyIPHourCount(context yawf.Context,gatewayHashKey,strategyKey, remoteIP string) int {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	timeStr := dateStr + "-" + strconv.Itoa(now.Hour()) 
	var redisKey string = "gatewayStrategy:" + gatewayHashKey  + ":" + strategyKey + ":IPHourCount:" + timeStr + ":" + remoteIP
	conn := cache.GetConnection(context)
	count, err := redis.Int(conn.Do("GET", redisKey))

	if err == redis.ErrNil {
		return -1
	} else if err != nil {
		panic(err)
	}
	return count
}

// 获取当前IP访问网关策略次数（分钟）
func GetGatewayStrategyIPMinuteCount(context yawf.Context,gatewayHashKey,strategyKey, remoteIP string) int {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	timeStr := dateStr + "-" + strconv.Itoa(now.Hour()) + "-" + strconv.Itoa(now.Minute())

	var redisKey string = "gatewayStrategy:" + gatewayHashKey  + ":" + strategyKey + ":IPMinuteCount:" + timeStr + ":" + remoteIP
	conn := cache.GetConnection(context)
	count, err := redis.Int(conn.Do("GET", redisKey))

	if err == redis.ErrNil {
		return -1
	} else if err != nil {
		panic(err)
	}
	return count
}

// 获取当前IP访问网关策略次数（秒）
func GetGatewayStrategyIPSecondCount(context yawf.Context,gatewayHashKey,strategyKey, remoteIP string) int {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	timeStr := dateStr + "-" + strconv.Itoa(now.Hour()) + "-" + strconv.Itoa(now.Minute()) + "-" + strconv.Itoa(now.Second())

	var redisKey string = "gatewayStrategy:" + gatewayHashKey  + ":" + strategyKey + ":IPSecondCount:" + timeStr  + ":" + remoteIP
	conn := cache.GetConnection(context)
	count, err := redis.Int(conn.Do("GET", redisKey))

	if err == redis.ErrNil {
		return -1
	} else if err != nil {
		panic(err)
	}
	return count
}

// 加载策略组秒阀值
func loadStrategySecondValve(strategyKey string) utils.RateLimitInfo{
	now := time.Now()
	db := database.GetConnection()
	sql := "SELECT intervalType,viewType,limitCount,priorityLevel FROM eo_gateway_rate_limit INNER JOIN eo_gateway_strategy_group ON eo_gateway_rate_limit.strategyID = eo_gateway_strategy_group.strategyID  WHERE eo_gateway_strategy_group.strategyKey = ? AND intervalType = 0 AND (HOUR(startTime) <= ? AND HOUR(endTime) > ?) OR (HOUR(startTime) > HOUR(endTime) AND HOUR(startTime) <= ? AND HOUR(endTime) > ?) ORDER BY priorityLevel DESC"
	hour := now.Hour()
	var rateLimitInfo utils.RateLimitInfo
	err := db.QueryRow(sql,strategyKey,hour,hour,hour,hour+24).Scan(&rateLimitInfo.IntervalType,&rateLimitInfo.ViewType,&rateLimitInfo.LimitCount,&rateLimitInfo.PriorityLevel)
	if err != nil{
		rateLimitInfo.ViewType = 0
		rateLimitInfo.IntervalType = 0
		rateLimitInfo.LimitCount = -1
		rateLimitInfo.PriorityLevel = 1
	}
	return rateLimitInfo
}

// 加载策略组分阀值
func loadStrategyMinuteValve(strategyKey string) utils.RateLimitInfo {
	now := time.Now()
	db := database.GetConnection()
	sql := "SELECT intervalType,viewType,limitCount,priorityLevel FROM eo_gateway_rate_limit INNER JOIN eo_gateway_strategy_group ON eo_gateway_rate_limit.strategyID = eo_gateway_strategy_group.strategyID  WHERE eo_gateway_strategy_group.strategyKey = ? AND intervalType = 1 AND (HOUR(startTime) <= ? AND HOUR(endTime) > ?) OR (HOUR(startTime) > HOUR(endTime) AND HOUR(startTime) <= ? AND HOUR(endTime) > ?) ORDER BY priorityLevel DESC"
	hour := now.Hour()
	var rateLimitInfo utils.RateLimitInfo
	err := db.QueryRow(sql,strategyKey,hour,hour,hour,hour+24).Scan(&rateLimitInfo.IntervalType,&rateLimitInfo.ViewType,&rateLimitInfo.LimitCount,&rateLimitInfo.PriorityLevel)
	if err != nil{
		rateLimitInfo.ViewType = 0
		rateLimitInfo.IntervalType = 1
		rateLimitInfo.LimitCount = -1
		rateLimitInfo.PriorityLevel = 1
	}
	return rateLimitInfo
}

// 加载策略组小时阀值
func loadStrategyHourValve(strategyKey string) utils.RateLimitInfo {
	now := time.Now()
	db := database.GetConnection()
	sql := "SELECT intervalType,viewType,limitCount,priorityLevel FROM eo_gateway_rate_limit INNER JOIN eo_gateway_strategy_group ON eo_gateway_rate_limit.strategyID = eo_gateway_strategy_group.strategyID  WHERE eo_gateway_strategy_group.strategyKey = ? AND intervalType = 2 AND (HOUR(startTime) <= ? AND HOUR(endTime) > ?) OR (HOUR(startTime) > HOUR(endTime) AND HOUR(startTime) <= ? AND HOUR(endTime) > ?) ORDER BY priorityLevel DESC"
	hour := now.Hour()
	var rateLimitInfo utils.RateLimitInfo
	err := db.QueryRow(sql,strategyKey,hour,hour,hour,hour+24).Scan(&rateLimitInfo.IntervalType,&rateLimitInfo.ViewType,&rateLimitInfo.LimitCount,&rateLimitInfo.PriorityLevel)
	if err != nil{
		rateLimitInfo.ViewType = 0
		rateLimitInfo.IntervalType = 2
		rateLimitInfo.LimitCount = -1
		rateLimitInfo.PriorityLevel = 1
	}
	return rateLimitInfo
}

// 加载策略组天阀值
func loadStrategyDayValve(strategyKey string) utils.RateLimitInfo {
	now := time.Now()
	db := database.GetConnection()
	sql := "SELECT intervalType,viewType,limitCount,priorityLevel,startTime,endTime FROM eo_gateway_rate_limit INNER JOIN eo_gateway_strategy_group ON eo_gateway_rate_limit.strategyID = eo_gateway_strategy_group.strategyID  WHERE eo_gateway_strategy_group.strategyKey = ? AND intervalType = 3 AND (HOUR(startTime) <= ? AND HOUR(endTime) > ?) OR (HOUR(startTime) > HOUR(endTime) AND HOUR(startTime) <= ? AND HOUR(endTime) > ?) ORDER BY priorityLevel DESC"
	hour := now.Hour()
	var rateLimitInfo utils.RateLimitInfo
	err := db.QueryRow(sql,strategyKey,hour,hour,hour,hour+24).Scan(&rateLimitInfo.IntervalType,&rateLimitInfo.ViewType,&rateLimitInfo.LimitCount,&rateLimitInfo.PriorityLevel,&rateLimitInfo.StartTime,&rateLimitInfo.EndTime)
	if err != nil{
		rateLimitInfo.ViewType = 0
		rateLimitInfo.IntervalType = 3
		rateLimitInfo.LimitCount = -1
		rateLimitInfo.PriorityLevel = 1
		rateLimitInfo.StartTime = "00:00"
		rateLimitInfo.EndTime = "00:00"
	}
	return rateLimitInfo
}

// 获取策略组秒阀值
func GetStrategySecondValve(context yawf.Context,hashKey,strategyKey string) utils.RateLimitInfo {
	now := time.Now()
	hour := now.Hour()
	var nextHour int
	if hour == 23{
		nextHour = 0
	}else{
		nextHour = hour + 1
	}
	var redisKey string = "gatewayStrategySecondValve:" + hashKey + ":" + strategyKey + ":" + strconv.Itoa(hour) + "-" + strconv.Itoa(nextHour)
	var info utils.RateLimitInfo
	conn := cache.GetConnection(context)
	infoStr, err := redis.String(conn.Do("GET", redisKey))
	if err == redis.ErrNil {
		info := loadStrategySecondValve(strategyKey)
		result, err := json.Marshal(info)
		if err != nil {
			panic(err)
		}
		infoStr = string(result)
		conn.Do("SET", redisKey, infoStr, "EX", 3600)
		
	} else if err != nil {
		panic(err)
	} else {
		err = json.Unmarshal([]byte(infoStr), &info)
		if err != nil {
			panic(err)
		}
	}
	return info
}

// 获取策略组分阀值
func GetStrategyMinuteValve(context yawf.Context,hashKey,strategyKey string) utils.RateLimitInfo {
	now := time.Now()
	hour := now.Hour()
	var nextHour int
	if hour == 23{
		nextHour = 0
	}else{
		nextHour = hour + 1
	}
	var redisKey string = "gatewayStrategyMinuteValve:" + hashKey + ":" + strategyKey + ":" + strconv.Itoa(hour) + "-" + strconv.Itoa(nextHour)
	var info utils.RateLimitInfo
	conn := cache.GetConnection(context)
	infoStr, err := redis.String(conn.Do("GET", redisKey))
	if err == redis.ErrNil {
		info = loadStrategyMinuteValve(strategyKey)
		result, err := json.Marshal(info)
		if err != nil {
			panic(err)
		}
		infoStr = string(result)
		conn.Do("SET", redisKey, infoStr, "EX", 3600)
	} else if err != nil {
		panic(err)
	} else {
		err = json.Unmarshal([]byte(infoStr), &info)
		if err != nil {
			panic(err)
		}
	}
	return info
}

// 获取策略组时阀值
func GetStrategyHourValve(context yawf.Context,hashKey,strategyKey string) utils.RateLimitInfo {
	now := time.Now()
	hour := now.Hour()
	var nextHour int
	if hour == 23{
		nextHour = 0
	}else{
		nextHour = hour + 1
	}
	var redisKey string = "gatewayStrategyHourValve:" + hashKey + ":" + strategyKey + ":" + strconv.Itoa(hour) + "-" + strconv.Itoa(nextHour)
	var info utils.RateLimitInfo
	conn := cache.GetConnection(context)
	infoStr, err := redis.String(conn.Do("GET", redisKey))
	if err == redis.ErrNil {
		info = loadStrategyHourValve(strategyKey)
		result, err := json.Marshal(info)
		if err != nil {
			panic(err)
		}
		infoStr = string(result)
		conn.Do("SET", redisKey, infoStr, "EX", 3600)
	} else if err != nil {
		panic(err)
	} else {
		err = json.Unmarshal([]byte(infoStr), &info)
		if err != nil {
			panic(err)
		}
	}
	return info
}

// 获取策略组天阀值
func GetStrategyDayValve(context yawf.Context,hashKey,strategyKey string) utils.RateLimitInfo {
	now := time.Now()
	hour := now.Hour()
	var nextHour int
	if hour == 23{
		nextHour = 0
	}else{
		nextHour = hour + 1
	}
	var redisKey string = "gatewayStrategyDayValve:" + hashKey + ":" + strategyKey + ":" + strconv.Itoa(hour) + "-" + strconv.Itoa(nextHour)
	fmt.Println(redisKey)
	var info utils.RateLimitInfo
	conn := cache.GetConnection(context)
	infoStr, err := redis.String(conn.Do("GET", redisKey))
	if err == redis.ErrNil {
		info = loadStrategyDayValve(strategyKey)
		result, err := json.Marshal(info)
		if err != nil {
			panic(err)
		}
		infoStr = string(result)
		conn.Do("SET", redisKey, infoStr, "EX", 3600)
	} else if err != nil {
		panic(err)
	} else {
		err = json.Unmarshal([]byte(infoStr), &info)
		if err != nil {
			panic(err)
		}
	}

	// 如果开始时间不等于结束时间，开始计数
	if info.StartTime != info.EndTime{
		startTime,_ := time.ParseInLocation("15:06",info.StartTime,time.Local)
		endTime,_ := time.ParseInLocation("15:06",info.EndTime,time.Local)
		var startHour, endHour int
		startHour = startTime.Hour()
		if endTime.Minute() >= 1 && endTime.Hour() == 23 {
			endHour =  24
		}else{
			endHour = endTime.Hour()
		}
		key := "gatewayStrategyPeriodCount:" + hashKey + ":" + strategyKey + ":" + strconv.Itoa(startHour) + "-" +strconv.Itoa(endHour)
		conn.Do("INCR", key)
	}
	return info
}


// 获取当天某一时间段的访问次数
func GetStrategyPeriodCount(context yawf.Context,hashKey,strategyKey,start,end string) int {
	startTime,_ := time.ParseInLocation("15:06",start,time.Local)
	endTime,_ := time.ParseInLocation("15:06",end,time.Local)
	var startHour, endHour int
	startHour = startTime.Hour()
	if endTime.Minute() >= 1 && endTime.Hour() == 23 {
		endHour =  24
	}else{
		endHour = endTime.Hour()
	}
	var redisKey string = "gatewayStrategyPeriodCount:" + hashKey + ":" + strategyKey + ":" + strconv.Itoa(startHour) + "-" +strconv.Itoa(endHour)
	conn := cache.GetConnection(context)
	count, err := redis.Int(conn.Do("GET", redisKey))

	if err == redis.ErrNil {
		return -1
	} else if err != nil {
		panic(err)
	}
	return count
}
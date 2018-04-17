package dao

import (
	"goku-ce-1.0/utils"
	"strconv"
	"time"
	"github.com/farseer810/yawf"
)

// 更新访问次数
func UpdateVisitCount(context yawf.Context, info *utils.MappingInfo,remoteIP string) {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	var redisKey string = "gatewayDayCount:" + info.GatewayHashKey + ":" + dateStr
	redisConn,err := utils.GetRedisConnection()
	defer redisConn.Close()
	if err != nil{
		panic(err)
	}
	// 更新网关当日访问次数
	redisConn.Do("INCR", redisKey)

	// 更新网关实时访问次数
	timeStr := dateStr + "-" + strconv.Itoa(now.Hour()) + "-" + strconv.Itoa(now.Minute())
	redisKey = "gatewayMinuteCount:" + info.GatewayHashKey + ":" + timeStr
	redisConn.Do("INCR", redisKey)
	
	// 更新当日网关策略访问次数
	redisKey = "gatewayStrategyDayCount:" + info.GatewayHashKey  + ":" + info.StrategyKey + ":" + dateStr
	redisConn.Do("INCR", redisKey)

	// 更新实时访问次数
	redisKey = "gatewayStrategyMinuteCount:" + info.GatewayHashKey  + ":" + info.StrategyKey + ":" + timeStr
	redisConn.Do("INCR", redisKey)
	
	// 更新当日网关策略IP访问次数
	redisKey = "gatewayStrategy:" + info.GatewayHashKey  + ":" + info.StrategyKey + ":IPDayCount:" + dateStr + ":" + remoteIP
	redisConn.Do("INCR", redisKey)

	// 更新网关策略IP访问次数(小时)
	redisKey = "gatewayStrategy:" + info.GatewayHashKey  + ":" + info.StrategyKey + ":IPHourCount:" + dateStr + "-" + strconv.Itoa(now.Hour()) + ":" + remoteIP
	redisConn.Do("INCR", redisKey)
	
	// 更新网关策略IP访问次数(分钟)
	redisKey = "gatewayStrategy:" + info.GatewayHashKey  + ":" + info.StrategyKey + ":IPMinuteCount:" + timeStr + ":" + remoteIP
	redisConn.Do("INCR", redisKey)

	// 更新网关策略IP访问次数(秒)
	redisKey = "gatewayStrategy:" + info.GatewayHashKey  + ":" + info.StrategyKey + ":IPSecondCount:" + timeStr + "-" + strconv.Itoa(now.Second()) + ":" + remoteIP
	redisConn.Do("INCR", redisKey)
}

// 更新成功次数
func UpdateSuccessCount(context yawf.Context, gatewayHashKey string) {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	var redisKey string = "gatewaySuccessCount:" + gatewayHashKey + ":" + dateStr
	redisConn,err := utils.GetRedisConnection()
	defer redisConn.Close()
	if err != nil{
		panic(err)
	}
	redisConn.Do("INCR", redisKey)
}

// 更新失败次数
func UpdateFailureCount(context yawf.Context,gatewayHashKey string) {
	now := time.Now()
	year, month, day := now.Date()
	dateStr := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	var redisKey string = "gatewayFailureCount:" + gatewayHashKey + ":" + dateStr
	redisConn,err := utils.GetRedisConnection()
	defer redisConn.Close()
	if err != nil{
		panic(err)
	}
	redisConn.Do("INCR", redisKey)
}
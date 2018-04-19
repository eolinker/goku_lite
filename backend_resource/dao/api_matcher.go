package dao

import (
	"goku-ce-1.0/dao/cache"
	"github.com/farseer810/yawf"
	"github.com/garyburd/redigo/redis"
	"goku-ce-1.0/dao/database"
	"strconv"
	"goku-ce-1.0/utils"
)

func GetAllAPIPaths(context yawf.Context, gatewayHashKey string) []string {
	result := getAllAPIPathsFromCache(context, gatewayHashKey)
	return result
}

func getAllAPIPathsFromCache(context yawf.Context, gatewayHashKey string) []string {
	conn := cache.GetConnection(context)

	var redisKey string = "apiList:" + gatewayHashKey
	result, err := redis.Strings(conn.Do("LRANGE", redisKey, 0, -1))
	if err != nil {
		panic(err)
	}else if len(result) == 0{
		result = loadAllAPIPathFromDB(gatewayHashKey)
	}
	return result
}

func loadAllAPIPathFromDB(gatewayHashkey string) []string {
	db := database.GetConnection()
	sql := "SELECT gatewayProtocol, gatewayRequestType, gatewayRequestURI FROM eo_gateway_api INNER JOIN eo_gateway ON eo_gateway_api.gatewayID = eo_gateway.gatewayID WHERE eo_gateway.hashkey = ?;"
	apiList := make([]string,0)
	rows,err := db.Query(sql,gatewayHashkey)
	if err != nil {
		return apiList
	}
	defer rows.Close()
	for rows.Next(){
		var gatewayProtocol,gatewayRequestType int
		var gatewayRequestURI string 
		err = rows.Scan(&gatewayProtocol,&gatewayRequestType,&gatewayRequestURI)
		if err != nil {
			break
		}
		api := strconv.Itoa(gatewayProtocol) + ":" + strconv.Itoa(gatewayRequestType) + ":" + gatewayRequestURI
		apiList = append(apiList,api)
	}
	redisConn,err := utils.GetRedisConnection()
	listName := "apiList:" + gatewayHashkey
	for _,i := range apiList {
		_	,err := redisConn.Do("rpush",listName,i)
		if err != nil{
			panic(err)
		}
	}
	return apiList
}

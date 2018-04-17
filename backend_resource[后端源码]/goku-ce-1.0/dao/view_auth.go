package dao

import (
	"goku-ce-1.0/dao/cache"
	"goku-ce-1.0/utils"
	"goku-ce-1.0/dao/database"
	"github.com/farseer810/yawf"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"encoding/base64"
)

func loadAuthInfo(strategyKey string) utils.AuthInfo {
	var authInfo utils.AuthInfo
	db := database.GetConnection()
	sql := "SELECT eo_gateway_strategy_group.strategyID,authType,apiKey,userName,userPassword FROM eo_gateway_auth INNER JOIN eo_gateway_strategy_group ON eo_gateway_strategy_group.strategyID = eo_gateway_auth.strategyID WHERE strategyKey = ?;"
	err := db.QueryRow(sql,strategyKey).Scan(&authInfo.StrategyID,&authInfo.AuthType,&authInfo.ApiKey,&authInfo.UserName,&authInfo.UserPassword)
	if err != nil{
		authInfo.AuthType = 2
	}
	if authInfo.AuthType == 0{
		authStr := []byte(authInfo.UserName + ":" + authInfo.UserPassword)
		authInfo.Authorization = base64.StdEncoding.EncodeToString(authStr)
	}
	return authInfo
}

func getAuthInfo(context yawf.Context, hashKey string, strategyKey string) utils.AuthInfo {
	var redisKey string = "authInfo:" + hashKey + ":" + strategyKey
	var info utils.AuthInfo
	conn := cache.GetConnection(context)
	infoStr, err := redis.String(conn.Do("GET", redisKey))
	if err == redis.ErrNil {
		info = loadAuthInfo(strategyKey)
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

func GetAuthInfo(context yawf.Context, hashKey,strategyKey string) utils.AuthInfo {
	return getAuthInfo(context, hashKey,strategyKey)
}

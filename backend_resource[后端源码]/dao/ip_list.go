package dao

import (
	"encoding/json"
	"goku-ce-1.0/dao/cache"
	"goku-ce-1.0/dao/database"
	"goku-ce-1.0/utils"
	"fmt"
	"strings"
	"github.com/farseer810/yawf"
	"github.com/garyburd/redigo/redis"
)

func loadGatewayIPList(context yawf.Context, hashKey,strategyKey string) (*utils.IPListInfo,bool) {
	var ipList []string
	var blackList string
	var whiteList string
	var chooseType int
	var info *utils.IPListInfo = &utils.IPListInfo{}
	flag := true
	db := database.GetConnection()
	sql := `SELECT chooseType,IFNULL(blackList,''),IFNULL(whiteList,'') FROM eo_gateway_strategy_group WHERE ` +
		`gatewayID=(SELECT gatewayID FROM eo_gateway WHERE hashKey=? AND strategyKey = ?) `
	err := db.QueryRow(sql, hashKey,strategyKey).Scan(&chooseType, &blackList, &whiteList)
	if err != nil {
		panic(err)
		flag =false
	}else{
		if chooseType == 1 {
			blackList = strings.Replace(blackList,"；",";",-1)
			ipList = strings.Split(blackList,";")
		} else if chooseType == 2 {
			whiteList = strings.Replace(whiteList,"；",";",-1)
			ipList = strings.Split(whiteList,";")
		} else {
			ipList = []string{}
		}
		info.IPList = ipList
		info.ChooseType = chooseType
	}
	
	return info,flag
}

func getGatewayIPList(context yawf.Context, hashKey,strategyKey string) *utils.IPListInfo {
	var redisKey string = "IPList:" + hashKey + ":" + strategyKey
	var info *utils.IPListInfo
	fmt.Println(hashKey)
	conn := cache.GetConnection(context)
	infoStr, err := redis.String(conn.Do("GET", redisKey))
	if err == redis.ErrNil {
		ipInfo,flag := loadGatewayIPList(context, hashKey,strategyKey)
		if flag{
			infoStr = ipInfo.String()
			// 缓存时间为1 hour
			conn.Do("SET", redisKey, infoStr, "EX", 3600)
		}
		info = ipInfo
	} else if err != nil {
		panic(err)
	} else {
		info = &utils.IPListInfo{}
		err = json.Unmarshal([]byte(infoStr), &info)
		if err != nil {
			panic(err)
		}
	}
	return info
}

func GetIPList(context yawf.Context, hashKey,strategyKey string) *utils.IPListInfo {
	return getGatewayIPList(context, hashKey,strategyKey)
}

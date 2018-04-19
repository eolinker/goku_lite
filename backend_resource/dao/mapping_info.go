package dao

import (
	"encoding/json"
	"goku-ce-1.0/dao/cache"
	"goku-ce-1.0/dao/database"
	"goku-ce-1.0/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/farseer810/yawf"
	"github.com/garyburd/redigo/redis"
)

func loadMappingInfo(hashKey string, matchedURI string) *utils.MappingInfo {
	var apiID int
	var jsonStr,path string
	parsedArray := strings.Split(matchedURI, ":")
	protocol, method, uri := parsedArray[0], parsedArray[1], parsedArray[2]
	fmt.Println(protocol, method, uri, hashKey)
	db := database.GetConnection()
	sql := `SELECT apiID FROM eo_gateway_api WHERE ` +
		`gatewayID=(SELECT gatewayID FROM eo_gateway WHERE hashKey=?) ` +
		`and gatewayProtocol=? and gatewayRequestType=? and gatewayRequestURI=?`
	iProtocol, _ := strconv.Atoi(protocol)
	iMethod, _ := strconv.Atoi(method)
	err := db.QueryRow(sql, hashKey, iProtocol, iMethod, uri).Scan(&apiID)
	if err != nil {
		panic(err)
	}
	fmt.Println(apiID)
	sql = `SELECT apiJson,path FROM eo_gateway_api_cache WHERE apiID=?`
	
	err = db.QueryRow(sql, apiID).Scan(&jsonStr,&path)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonStr)
	info := utils.ParseDBJson(jsonStr,path)
	info.ApiID = apiID
	
	return info
}

func getMapping(context yawf.Context, hashKey string, matchedURI string) *utils.MappingInfo {
	var redisKey string = "apiInfo:" + hashKey + ":" + matchedURI
	var info *utils.MappingInfo
	conn := cache.GetConnection(context)
	infoStr, err := redis.String(conn.Do("GET", redisKey))
	if err == redis.ErrNil {
		info = loadMappingInfo(hashKey, matchedURI)
		infoStr = info.String()

		// 缓存时间为1 hour
		conn.Do("SET", redisKey, infoStr, "EX", 3600)
	} else if err != nil {
		panic(err)
	} else {
		info = &utils.MappingInfo{}
		err = json.Unmarshal([]byte(infoStr), &info)
		if err != nil {
			panic(err)
		}
	}
	return info
}

func GetMapping(context yawf.Context, hashKey string, matchedURI string) *utils.MappingInfo {
	return getMapping(context, hashKey, matchedURI)
}

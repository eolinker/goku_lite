package dao

import (
	"goku-ce-1.0/dao/cache"
	"github.com/farseer810/yawf"
	"github.com/garyburd/redigo/redis"
	"goku-ce-1.0/dao/database"
)

func loadGatewayHashKey(gatewayAlias string) string{
	db := database.GetConnection()
	var gatewayHashKey string
	sql := `SELECT hashKey FROM eo_gateway WHERE gatewayAlias = ?; `
	err := db.QueryRow(sql,gatewayAlias).Scan(&gatewayHashKey)
	if err != nil {
		panic(err)
	}
	return gatewayHashKey
}

func GetGatewayHashKey(context yawf.Context,gatewayAlias string) string {
	var redisKey string = "gatewayHashKey:" + gatewayAlias
	conn := cache.GetConnection(context)
	gatewayHashKey, err := redis.String(conn.Do("GET", redisKey))
	if err == redis.ErrNil {
		gatewayHashKey = loadGatewayHashKey(gatewayAlias)
		conn.Do("SET", redisKey, gatewayHashKey)
	} else if err != nil {
		panic(err)
	}
	return gatewayHashKey
}

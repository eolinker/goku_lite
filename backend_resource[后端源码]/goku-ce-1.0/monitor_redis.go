package main

import (
	"goku-ce-1.0/utils"
	"goku-ce-1.0/conf"
	"time"
	"strconv"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	redis "github.com/garyburd/redigo/redis"
	"encoding/json"
	
)
var db *sql.DB
var redisConn redis.Conn
func init(){
	var err error
	db = getConnection()
	redisConn,err = getRedisConnection()
	if err != nil{
		panic(err)
	}
}

func main() {
	dealTask()
}

func getRedisConnection() (redis.Conn, error) {
	redisDB, err := strconv.Atoi(conf.Configure["redis_db"])
	if err != nil {
		redisDB = 0
	}
	if _, hasPassword := conf.Configure["redis_password"]; hasPassword {
			return redis.Dial("tcp",
				conf.Configure["redis_host"]+":"+conf.Configure["redis_port"],
				redis.DialPassword(conf.Configure["redis_password"]),
				redis.DialDatabase(redisDB))
	} else {
			return redis.Dial("tcp",
				conf.Configure["redis_host"]+":"+conf.Configure["redis_port"],
				redis.DialDatabase(redisDB))
	}
}

func getConnection() *sql.DB {
	var err error
	var dsn string = conf.Configure["mysql_username"] + ":" + conf.Configure["mysql_password"]
	dsn = dsn + "@tcp(" + conf.Configure["mysql_host"] + ":" + conf.Configure["mysql_port"] + ")/" + conf.Configure["mysql_dbname"]
	dsn = dsn + "?charset=utf8"
	db, err = sql.Open("mysql", dsn)
	fmt.Println(dsn)
	if err == nil {
		if err = db.Ping();err !=nil {
			panic(err)
		}
		db.SetMaxOpenConns(2000)
		db.SetMaxIdleConns(1000)
	} else {
		panic(err)
	}
	return db
}

func dealTask(){
	last := time.Now()
	for {
		task,err := redis.StringMap(redisConn.Do("blpop","gatewayQueue", 1))
		if err != nil{
			_,err = redisConn.Do("get","gatewayQueueCloseSignal")
			if err !=nil { 
				break
			}
			now := time.Now()
			if now.Sub(last).Minutes() >= 1{
				last = now
				fmt.Println(last.Format("2006-01-02 15:04:05"))
			}
			continue
		}
		time.Sleep(1000)

		var taskQueue utils.QueryJson
		json.Unmarshal([]byte(task["gatewayQueue"]),&taskQueue)
		dataStr,err := json.Marshal(taskQueue.Data)
		if err != nil{
			panic(err)
		}

		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + ":" + task["gatewayQueue"])
		
		var data utils.OperationData
		json.Unmarshal([]byte(dataStr),&data)
		if taskQueue.OperationType == "gateway" {
			
			if taskQueue.Operation == "add" {
				hash_key, gateway_alias := data.GatewayHashKey, data.GatewayAlias
				passAddGateway(hash_key,gateway_alias)

			}else if taskQueue.Operation == "delete" {
				hash_key, gateway_alias := data.GatewayHashKey, data.GatewayAlias
				passDeleteGateway(hash_key,gateway_alias)
			}

		}else if taskQueue.OperationType == "backend" {
			hash_key, gateway_id := data.GatewayHashKey, data.GatewayID
			time.Sleep(1 * time.Second)
			deleteApiInfo(gateway_id,hash_key)
			
		}else if taskQueue.OperationType == "api" {
			hash_key, gateway_id := data.GatewayHashKey, data.GatewayID
			time.Sleep(1 * time.Second)
			loadAPIList(gateway_id,hash_key)
		}
    }
	
}

func passAddGateway(hash_key,gateway_alias string) {
	_,err := redisConn.Do("SET","gatewayHashKey:" + gateway_alias,hash_key)
	if err != nil{
		panic(err)
	}
}

func passDeleteGateway(hash_key,gateway_alias string) {
	redisConn.Do("apiList","apiList:" + hash_key)
	redisConn.Do("del","gatewayHashKey:" + gateway_alias)
}

func deleteApiInfo(gateway_id int,hash_key string) {
	keys,err := redis.Strings(redisConn.Do("keys","apiInfo:" + hash_key + "*"))
	if err != nil{
		panic(err)
	}
	if len(keys) > 0 {
		fmt.Println(keys)
		for _,key := range keys{
			_,err = redisConn.Do("del",key)
			if err != nil{
				panic(err)
			}
		}
	}
}

func loadAPIList(gateway_id int,hash_key string) {
	sql := "SELECT gatewayProtocol, gatewayRequestType, gatewayRequestURI FROM eo_gateway_api WHERE gatewayID = ?"
	rows,err := db.Query(sql,gateway_id)
	if err != nil {
		panic(err)
	}
	apis := make([]string,0)
	defer rows.Close()
	//获取记录列
	
		for rows.Next(){
			var gatewayProtocol,gatewayRequestType int
			var gatewayRequestURI string 
			err = rows.Scan(&gatewayProtocol,&gatewayRequestType,&gatewayRequestURI)
			if err != nil {
				break
			}
			api := strconv.Itoa(gatewayProtocol) + ":" + strconv.Itoa(gatewayRequestType) + ":" + gatewayRequestURI
			apis = append(apis,api)
		}
	listName := "apiList:" + hash_key
	for _,i := range apis {
		_,err := redisConn.Do("rpush",listName,i)
		if err != nil{
			panic(err)
		}
	}
	fmt.Print("apis:")
	fmt.Println(apis)
}

	

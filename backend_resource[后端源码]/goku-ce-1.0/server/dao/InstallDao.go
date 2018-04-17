package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"goku-ce-1.0/dao/database"
)

// 检查数据库是否可以连接
func CheckDBConnect(mysql_username,mysql_password,mysql_host,mysql_port,mysql_dbname string) bool{
	var dsn string = mysql_username + ":" + mysql_password
	dsn = dsn + "@tcp(" + mysql_host + ":" + mysql_port + ")/" + mysql_dbname + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err == nil {
		err = db.Ping()  
		if err != nil {  
			return false
		}else{
			return true
		}
	} else {
		return false
	}
}

// 检查Redis是否可以连接
func CheckRedisConnect(redis_db,redis_password,redis_host,redis_port string) (bool){
	db, err := strconv.Atoi(redis_db)
	if err != nil {
		db = 0
	}
	if redis_password != ""{
			_,err = redis.Dial("tcp",redis_host+":"+redis_port,redis.DialPassword(redis_password),redis.DialDatabase(db))
			if err == nil{
				return true
			}else{
				return false
			}
	} else {
			_,err =  redis.Dial("tcp",redis_host+":"+redis_port,redis.DialDatabase(db))
			if err == nil{
				return true
			}else{
				return false
			}
	}
}

// 检查是否安装
func CheckIsInstall() bool{
	db := database.GetConnection()
	err := db.Ping()
	if err != nil{
		return false
	}
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM eo_admin").Scan(&count)
	if err != nil{
		return false
	}
	if count == 0{
		return false
	}else{
		return true
	}
}


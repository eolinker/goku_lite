package module

import (
	"goku-ce-1.0/server/dao"
)

// 检查数据库是否可以连接
func CheckDBConnect(mysql_username,mysql_password,mysql_host,mysql_port,mysql_dbname string) bool{
	return dao.CheckDBConnect(mysql_username,mysql_password,mysql_host,mysql_port,mysql_dbname)
}

// 检查Redis是否可以连接
func CheckRedisConnect(redis_db,redis_password,redis_host,redis_port string) (bool){
	return dao.CheckRedisConnect(redis_db,redis_password,redis_host,redis_port)
}



func CheckIsInstall() bool{
	return dao.CheckIsInstall()
}

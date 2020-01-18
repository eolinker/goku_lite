package main

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/common/database"
	"github.com/eolinker/goku-api-gateway/common/pdao"
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
	"github.com/eolinker/goku-api-gateway/server/entity"
)



//ClusterDatabaseConfig 集群数据库配置
type ClusterDatabaseConfig entity.ClusterDB

//GetDriver 获取驱动类型
func (c *ClusterDatabaseConfig) GetDriver() string {
	return c.Driver
}

//GetSource 获取连接字符串
func (c *ClusterDatabaseConfig) GetSource() string {

	switch c.Driver {
	case database.MysqlDriver:
		{
			return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.UserName, c.Password, c.Host, c.Port, c.Database)
		}
	case database.Sqlite3Driver:
		{
			return c.Path
		}
	default:
		{
			return ""
		}
	}

}

//InitDatabase 初始化数据库
func InitDatabase() {
	console_sqlite3.DoRegister()

	def, err := getDefaultDatabase()
	if err != nil {
		panic(err)
	}
	c := ClusterDatabaseConfig(*def)
	db,e := database.InitConnection(&c)
	if e != nil {
		panic(e)
	}
	err =pdao.Build(c.GetDriver(),db)
	if err!=nil{
		panic(err)
	}

	err =pdao.Check()
	if err!=nil{
		panic(err)
	}
}


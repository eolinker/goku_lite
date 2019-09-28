package console

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/common/database"
	"github.com/eolinker/goku-api-gateway/server/entity"
)

//ClusterDatabaseConfig 集群数据库对象
type ClusterDatabaseConfig entity.ClusterDB

//GetDriver 获取驱动
func (c *ClusterDatabaseConfig) GetDriver() string {
	return c.Driver
}

//GetSource 获取链接字符串
func (c *ClusterDatabaseConfig) GetSource() string {

	//dsn = conf.Value("db_user") + ":" + conf.Value("db_password")
	//dsn = dsn + "@tcp(" + conf.Value("db_host") + ":" + conf.Value("db_port") + ")/" + conf.Value("db_name")
	//dsn = dsn + "?charset=utf8"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.UserName, c.Password, c.Host, c.Port, c.Database)
}

//InitDatabase 初始化数据库
func InitDatabase() {
	def, err := getDefaultDatabase()
	if err != nil {
		panic(err)
	}
	c := ClusterDatabaseConfig(*def)
	e := database.InitConnection(&c)
	if e != nil {
		panic(e)
	}
}

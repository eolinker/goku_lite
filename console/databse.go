package console

import (
	"fmt"
	"github.com/eolinker/goku/common/database"
	"github.com/eolinker/goku/server/entity"
)

type ClusterDatabaseConfig entity.ClusterDB

func (c *ClusterDatabaseConfig) GetDriver() string {
	return c.Driver
}

func (c *ClusterDatabaseConfig) GetSource() string {

	//dsn = conf.Value("db_user") + ":" + conf.Value("db_password")
	//dsn = dsn + "@tcp(" + conf.Value("db_host") + ":" + conf.Value("db_port") + ")/" + conf.Value("db_name")
	//dsn = dsn + "?charset=utf8"
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.UserName, c.Password, c.Host, c.Port, c.Database)
}

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

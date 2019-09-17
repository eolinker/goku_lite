package dao_service

import (
	"github.com/eolinker/goku/common/database"
	"time"
)

const sqlAdd="INSERT INTO `goku_service_config`(`name`,`driver`,`default`,`desc`,`config`,`clusterConfig`,`healthCheck`,`healthCheckPath`,`healthCheckPeriod`,`healthCheckCode`,`healthCheckTimeOut`,`createTime`,`updateTime`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);"
func Add(name, driver, desc, config, clusterConfig string,isDefault,healthCheck bool,healthCheckPath string,healthCheckCode string,healthCheckPeriod ,healthCheckTimeOut int) error {

	now:=time.Now()

	stmt, e := database.GetConnection().Prepare(sqlAdd)
	if e!=nil{
		return e
	}

	_,err:=stmt.Exec(name,driver,isDefault,desc,config,clusterConfig,healthCheck,healthCheckPath,healthCheckPeriod,healthCheckCode,healthCheckTimeOut,now,now)
	return err
}


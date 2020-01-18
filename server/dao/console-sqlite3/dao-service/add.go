package dao_service

import (
	"time"
)

const sqlAdd = "INSERT INTO `goku_service_config`(`name`,`driver`,`default`,`desc`,`config`,`clusterConfig`,`healthCheck`,`healthCheckPath`,`healthCheckPeriod`,`healthCheckCode`,`healthCheckTimeOut`,`createTime`,`updateTime`)VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);"

//Add 新增服务
func (d *ServiceDao) Add(name, driver, desc, config, clusterConfig string, isDefault, healthCheck bool, healthCheckPath string, healthCheckCode string, healthCheckPeriod, healthCheckTimeOut int) error {

	now := time.Now().Format("2006-01-02 15:04:05")

	stmt, e := d.db.Prepare(sqlAdd)
	if e != nil {
		return e
	}
	defer stmt.Close()

	_, err := stmt.Exec(name, driver, isDefault, desc, config, clusterConfig, healthCheck, healthCheckPath, healthCheckPeriod, healthCheckCode, healthCheckTimeOut, now, now)
	return err
}

package dao_service

import (
	"time"
)

const sqlSave = "UPDATE `goku_service_config` SET `desc`=?,`config`=?,`clusterConfig`=?,`healthCheck`=?,`healthCheckPath`=?,`healthCheckPeriod`=?,`healthCheckCode`=?,`healthCheckTimeOut`=?,`updateTime`=? WHERE `name`=?;"

//Save 存储服务发现信息
func (d *ServiceDao) Save(name, desc, config, clusterConfig string, healthCheck bool, healthCheckPath string, healthCheckCode string, healthCheckPeriod, healthCheckTimeOut int) error {
	now := time.Now().Format("2006-01-02 15:04:05")

	stmt, e := d.db.Prepare(sqlSave)
	if e != nil {
		return e
	}
	defer stmt.Close()
	_, err := stmt.Exec(desc, config, clusterConfig, healthCheck, healthCheckPath, healthCheckPeriod, healthCheckCode, healthCheckTimeOut, now, name)
	return err
}

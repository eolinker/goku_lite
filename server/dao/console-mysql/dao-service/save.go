package dao_service

import (
	"github.com/eolinker/goku/common/database"
	"time"
)

const sqlSave = "UPDATE `goku_service_config` SET `desc`=?,`config`=?,`clusterConfig`=?,`healthCheck`=?,`healthCheckPath`=?,`healthCheckPeriod`=?,`healthCheckCode`=?,`healthCheckTimeOut`=?,`updateTime`=? WHERE `name`=?;"

func Save(name, desc, config, clusterConfig string, healthCheck bool, healthCheckPath string, healthCheckCode string, healthCheckPeriod, healthCheckTimeOut int) error {
	now := time.Now()

	stmt, e := database.GetConnection().Prepare(sqlSave)
	if e != nil {
		return e
	}

	_, err := stmt.Exec(desc, config, clusterConfig, healthCheck, healthCheckPath, healthCheckPeriod, healthCheckCode, healthCheckTimeOut, now, name)
	return err
}

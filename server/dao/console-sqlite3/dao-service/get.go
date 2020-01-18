package dao_service

import (
	"fmt"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

const sqlGet = "SELECT `name`,`driver`,`default`,`desc`,`config`,`clusterConfig`,`healthCheck`,`healthCheckPath`,`healthCheckPeriod`,`healthCheckCode`,`healthCheckTimeOut`,`createTime`,`updateTime` FROM `goku_service_config` WHERE `name`=?; "

//Get 获取服务发现信息
func (d *ServiceDao) Get(name string) (*entity.Service, error) {

	stmt, e := d.db.Prepare(sqlGet)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, err := stmt.Query(name)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {

		v := new(entity.Service)
		er := rows.Scan(&v.Name,
			&v.Driver,
			&v.IsDefault,
			&v.Desc,
			&v.Config,
			&v.ClusterConfig,
			&v.HealthCheck,
			&v.HealthCheckPath,
			&v.HealthCheckPeriod,
			&v.HealthCheckCode,
			&v.HealthCheckTimeOut,
			&v.CreateTime,
			&v.UpdateTime,
		)
		if er != nil {
			return nil, er
		}

		return v, nil
	}

	return nil, fmt.Errorf("no that service:%s", name)

}

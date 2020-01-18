package dao_service

import (
	"fmt"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

const sqlList = "SELECT `name`,`driver`,`default`,`desc`,`config`,`clusterConfig`,`healthCheck`,`healthCheckPath`,`healthCheckPeriod`,`healthCheckCode`,`healthCheckTimeOut`,`createTime`,`updateTime` FROM `goku_service_config` %s ORDER BY `updateTime` DESC;"

//List 获取服务发现列表
func (d *ServiceDao) List(keyword string) ([]*entity.Service, error) {
	where := ""
	if keyword != "" {
		where = fmt.Sprint("where `name` like '%", keyword, "%' OR `driver` like '%", keyword, "%'")
	}

	sql := fmt.Sprintf(sqlList, where)
	stmt, e := d.db.Prepare(sql)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vs := make([]*entity.Service, 0, 10)

	for rows.Next() {

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

		vs = append(vs, v)

	}
	return vs, nil

}

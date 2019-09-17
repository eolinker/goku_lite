package dao_service

import (
	"github.com/eolinker/goku/common/database"
	entity "github.com/eolinker/goku/server/entity/node-entity"
)

const sqlList="SELECT `name`,`driver`,`default`,`desc`,`config`,`clusterConfig`,`healthCheck`,`healthCheckPath`,`healthCheckPeriod`,`healthCheckCode`,`healthCheckTimeOut` FROM `goku_service_config`"
func GetAll() ([]*entity.Service,error) {

	stmt, e := database.GetConnection().Prepare(sqlList)
	if e!=nil{
		return nil,e
	}
	defer stmt.Close()
	rows,err:=stmt.Query()

	if err!=nil{
		return nil,err
	}
	defer rows.Close()

	vs:=make([]*entity.Service,0,10)

	for rows.Next(){

		v:=new(entity.Service)
		er:=rows.Scan(&v.Name,
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

		)
		if er!=nil{
			return nil,er
		}

		vs = append(vs,v)
	}
	return vs,nil

}
package dao_service

import (
	"fmt"
	"github.com/eolinker/goku/common/database"
	entity "github.com/eolinker/goku/server/entity/console-entity"
)

const sqlGet  = "SELECT `name`,`driver`,`default`,`desc`,`config`,`clusterConfig`,`healthCheck`,`healthCheckPath`,`healthCheckPeriod`,`healthCheckCode`,`healthCheckTimeOut`,`createTime`,`updateTime` FROM `goku_service_config` WHERE `name`=?; "
func Get(name string)(*entity.Service, error)  {

	stmt, e := database.GetConnection().Prepare(sqlGet)
	if e!=nil{
		return nil,e
	}
	defer stmt.Close()
	rows,err:=stmt.Query(name)

	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	if rows.Next(){

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
			&v.CreateTime,
			&v.UpdateTime,
			)
		if er!=nil{
			return nil,er
		}


		return v,nil
	}

	return nil,fmt.Errorf("no that service:%s",name)

}
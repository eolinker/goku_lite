package dao_balance

import (
	"github.com/eolinker/goku/common/database"
	entity "github.com/eolinker/goku/server/entity/node-entity"
)

func GetAllBalance() ([]*entity.Balance ,error){
	const sql = "SELECT A.`balanceName`,A.`serviceName`,IFNULL(B.`driver`,''),A.`appName`,IFNULL(A.`static`,''),IFNULL(A.`staticCluster`,'') FROM `goku_balance` A LEFT JOIN `goku_service_config` B ON A.`serviceName` = B.`name`;"
	db:=database.GetConnection()
	rows,err:= db.Query(sql)
	if err!= nil{
		return nil,err
	}
	defer rows.Close()

	r:=make([]*entity.Balance,0,20)
	for rows.Next(){
		v:=new(entity.Balance)
		err:=rows.Scan( &v.Name,&v.ServiceName,&v.ServiceDriver,&v.AppName,&v.Static,&v.StaticCluster)
		if err!= nil{
			return nil,err
		}
		r= append(r,v.Type())
	}
	return r,nil
}
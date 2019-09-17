package dao_balance_update

import (
	"github.com/eolinker/goku/common/database"
	dao_service "github.com/eolinker/goku/server/dao/console-mysql/dao-service"
	entity "github.com/eolinker/goku/server/entity/balance-entity"
)

func GetAllOldVerSion() ([]*entity.BalanceInfoEntity ,error){
	const sql = "SELECT `balanceName`,IFNULL(`balanceDesc`, ''),IFNULL(`balanceConfig`, ''),IFNULL(`defaultConfig`, ''),IFNULL(`clusterConfig`, ''),`updateTime`,`createTime` FROM `goku_balance` WHERE `serviceName` = '';"
	db:=database.GetConnection()
	rows,err:= db.Query(sql)
	if err!= nil{
		return nil,err
	}
	defer rows.Close()

	r:=make([]*entity.BalanceInfoEntity,0,20)
	for rows.Next(){
		v:=new(entity.BalanceInfoEntity)
		err:=rows.Scan( &v.Name,&v.Desc,&v.OldVersionConfig, &v.DefaultConfig, &v.ClusterConfig, &v.UpdateTime, &v.CreateTime)
		if err!= nil{
			return nil,err
		}
		r= append(r,v)
	}

	return r,nil
}

func GetDefaultServiceStatic()string  {

	tx := database.GetConnection()
	name:=""
	err:=tx.QueryRow("SELECT `name` FROM `goku_service_config` WHERE `driver`='static' ORDER BY  `default` DESC LIMIT 1; ").Scan(&name)
	if err!=nil{
		name ="static"
	    dao_service.Add(name,"static","默认静态服务","","",false,false,"","",5,300)
	}

	return name
}
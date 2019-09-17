package balance

import (
	"encoding/json"
	"github.com/eolinker/goku/common/general"
	dao_balance_update "github.com/eolinker/goku/server/dao/console-mysql/dao-balance-update"
	entity "github.com/eolinker/goku/server/entity/balance-entity"
)

func init() {
	general.RegeditLater(Update)
}
func Update()error  {

	l,e:= dao_balance_update.GetAllOldVerSion()
	if e!=nil{
		return e
	}

	defStaticServiceName :=dao_balance_update.GetDefaultServiceStatic()
	for _,e:=range l{
		update(e,defStaticServiceName)
	}



	return nil

}

func update(e *entity.BalanceInfoEntity,serviceName string)  {

	if e==nil{
		return
	}

	param:=&Param{
		Name:          e.Name,
		ServiceName:   serviceName,
		AppName:       "",
		Static:        "",
		StaticCluster: "",
		Desc:          e.Desc,
	}

	info,err:=e.Decode()

	if err!=nil{
		return
	}


	if info.Default!= nil{

		param.Static = info.Default.ServersConfigOrg
	}
	if info.Cluster !=nil{
		cluster:=make(map[string]string)
		for clusterName,server:=range info.Cluster{
			cluster[clusterName] = server.ServersConfigOrg

		}

		data ,err:= json.Marshal(cluster)

		if err==nil{

			param.StaticCluster = string(data)
		}
	}

	Save(param)


}
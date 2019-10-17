package balance

import (
	"encoding/json"

	"github.com/eolinker/goku-api-gateway/common/general"
	dao_balance_update2 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/dao-balance-update"
	entity "github.com/eolinker/goku-api-gateway/server/entity/balance-entity"
)

func init() {
	general.RegeditLater(Update)
}

//Update 将旧负载配置更新为新负载配置
func Update() error {

	l, e := dao_balance_update2.GetAllOldVerSion()
	if e != nil {
		return e
	}

	defStaticServiceName := dao_balance_update2.GetDefaultServiceStatic()
	for _, e := range l {
		update(e, defStaticServiceName)
	}

	return nil

}

func update(e *entity.BalanceInfoEntity, serviceName string) {

	if e == nil {
		return
	}

	param := &Param{
		Name:          e.Name,
		ServiceName:   serviceName,
		AppName:       "",
		Static:        "",
		StaticCluster: "",
		Desc:          e.Desc,
	}

	info, err := e.Decode()

	if err != nil {
		return
	}

	if info.Default != nil {

		param.Static = info.Default.ServersConfigOrg
	}
	if info.Cluster != nil {
		cluster := make(map[string]string)
		for clusterName, server := range info.Cluster {
			cluster[clusterName] = server.ServersConfigOrg

		}

		data, err := json.Marshal(cluster)

		if err == nil {

			param.StaticCluster = string(data)
		}
	}

	Save(param)

}

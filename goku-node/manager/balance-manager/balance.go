package balance_manager

import (
	"encoding/json"
	log "github.com/eolinker/goku/goku-log"
	"github.com/eolinker/goku/goku-node/manager/updater"
	node_common "github.com/eolinker/goku/goku-node/node-common"
	"github.com/eolinker/goku/goku-service/application"
	"github.com/eolinker/goku/goku-service/balance"
	dao_balance "github.com/eolinker/goku/server/dao/node-mysql/dao-balance"
	"github.com/eolinker/goku/server/driver"
	entity "github.com/eolinker/goku/server/entity/node-entity"
	"strings"
)

func init() {
	updater.Add(loadBalanceInfo, 3, "goku_balance")
}

func Get(name string) (application.IHttpApplication, bool) {

	return balance.GetByName(name)
}

// 加载负载均衡信息
func loadBalanceInfo() {

	balanceList, err := getAllBalance()
	if err != nil {
		log.Info(err)
		return
	}

	balance.ResetBalances(balanceList)
}
func genBalance(e *entity.Balance) *balance.Balance {
	b := &balance.Balance{
		Name:      e.Name,
		Discovery: e.ServiceName,
		AppConfig: "",
	}

	switch e.ServiceType {
	case driver.Static:
		{
			var m map[string]string
			b.AppConfig = e.Static
			if err := json.Unmarshal([]byte(e.StaticCluster), &m); err == nil {
				if v, has := m[node_common.ClusterName()]; has {
					if len(strings.Trim(v, " ")) > 0 {
						b.AppConfig = v
					}
				}
			}
			return b
		}
	case driver.Discovery:
		{
			b.AppConfig = e.AppName
			return b
		}

	}
	return nil
}
func getAllBalance() ([]*balance.Balance, error) {

	entities, e := dao_balance.GetAllBalance()
	if e != nil {
		return nil, e
	}
	infos := make([]*balance.Balance, 0, len(entities))

	for _, ent := range entities {
		b := genBalance(ent)
		if b != nil {
			infos = append(infos, b)
		}
	}
	return infos, nil
}

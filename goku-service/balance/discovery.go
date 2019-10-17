package balance

import (
	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/goku-service/application"
	"github.com/eolinker/goku-api-gateway/goku-service/discovery"
)

//ResetBalances 重置负载列表
func ResetBalances(balances map[string]*config.BalanceConfig) {
	manager.set(balances)
}

//GetByName 通过名称获取负载
func GetByName(name string) (application.IHttpApplication, bool) {
	b, has := manager.get(name)
	if !has {
		return application.NewOrg(name), true
	}

	sources, has := discovery.GetDiscoverer(b.DiscoverName)
	if has {

		service, handler, yes := sources.GetApp(b.Config)
		if yes {
			return application.NewApplication(service, handler), true
		}
	}

	return nil, false
}

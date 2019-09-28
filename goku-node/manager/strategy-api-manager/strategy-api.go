package strategyapimanager

import (
	"sync"

	"github.com/eolinker/goku-api-gateway/goku-node/manager/updater"
	dao_strategy "github.com/eolinker/goku-api-gateway/server/dao/node-mysql/dao-strategy"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

func init() {
	updater.Add(loadStategyAPI, 6, "goku_conn_strategy_api", "goku_gateway_strategy", "goku_gateway_api")
}

var (
	apiOfStrategy = make(map[string]*_APIMap)
	apiLocker     sync.RWMutex
)

//GetAPIForStategy 获取指定策略、API 对应的 配置
func GetAPIForStategy(stategyID string, apiID int) (*entity.StrategyAPI, bool) {
	apis, has := getAPis(stategyID)
	if !has {
		return nil, false
	}
	return apis.Get(apiID)
}

func getAPis(strategyID string) (*_APIMap, bool) {
	apis, has := apiOfStrategy[strategyID]
	return apis, has
}
func resetAPIs(apis map[string]*_APIMap) {
	apiLocker.Lock()
	defer apiLocker.Unlock()

	apiOfStrategy = apis
}
func loadStategyAPI() {
	apis, e := dao_strategy.GetAllStrategyAPI()
	if e != nil {
		return
	}

	tem := make(map[string]*_APIMap)
	for _, api := range apis {

		as, has := tem[api.StrategyID]
		if !has {
			as = &_APIMap{
				apis: make(map[int]*entity.StrategyAPI),
			}
		}
		as.apis[api.APIID] = api
		tem[api.StrategyID] = as
	}
	resetAPIs(tem)
}

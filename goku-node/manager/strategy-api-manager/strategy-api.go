package strategy_api_manager

import (
	"sync"

	"github.com/eolinker/goku/goku-node/manager/updater"
	dao_strategy "github.com/eolinker/goku/server/dao/node-mysql/dao-strategy"
	entity "github.com/eolinker/goku/server/entity/node-entity"
)

func init() {
	updater.Add(loadStategyApi, 6, "goku_conn_strategy_api", "goku_gateway_strategy", "goku_gateway_api")
}

var (
	apiOfStrategy = make(map[string]*_ApiMap)
	apiLocker     sync.RWMutex
)

//GetAPIForStategy 获取指定策略、API 对应的 配置
func GetAPIForStategy(stategyId string, apiId int) (*entity.StrategyApi, bool) {
	apis, has := getAPis(stategyId)
	if !has {
		return nil, false
	}
	return apis.Get(apiId)
}

func getAPis(strategyId string) (*_ApiMap, bool) {
	apis, has := apiOfStrategy[strategyId]
	return apis, has
}
func resetApis(apis map[string]*_ApiMap) {
	apiLocker.Lock()
	defer apiLocker.Unlock()

	apiOfStrategy = apis
}
func loadStategyApi() {
	apis, e := dao_strategy.GetAllStrategyApi()
	if e != nil {
		return
	}

	tem := make(map[string]*_ApiMap)
	for _, api := range apis {

		as, has := tem[api.StrategyID]
		if !has {
			as = &_ApiMap{
				apis: make(map[int]*entity.StrategyApi),
			}
		}
		as.apis[api.ApiId] = api
		tem[api.StrategyID] = as
	}
	resetApis(tem)
}

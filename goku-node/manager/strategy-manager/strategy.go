package strategymanager

import (
	"sync"

	"github.com/eolinker/goku-api-gateway/goku-node/manager/updater"
	dao_strategy "github.com/eolinker/goku-api-gateway/server/dao/node-mysql/dao-strategy"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

func init() {
	updater.Add(loadStategy, 4, "goku_gateway_strategy")
	// loadStategy()
}

//GetAnonymous 获取匿名策略
//return: 开放策略ID，是否有效
func GetAnonymous() (*entity.Strategy, bool) {
	if defStrategy == nil || defStrategy.EnableStatus != 1 {
		return nil, false
	}
	return defStrategy, true
}

//CheckStategy 测试策略ID是否生效
func CheckStategy(stategyID string) bool {
	lockerStategy.RLock()
	defer lockerStategy.RUnlock()
	_, has := strategys[stategyID]
	return has
}

var (
	strategys     = make(map[string]*entity.Strategy)
	defStrategy   *entity.Strategy
	lockerStategy sync.RWMutex
)

func reset(s map[string]*entity.Strategy, def *entity.Strategy) {
	lockerStategy.Lock()
	defer lockerStategy.Unlock()
	strategys = s
	defStrategy = def
}

//Get get
func Get(id string) (*entity.Strategy, bool) {
	lockerStategy.RLock()
	defer lockerStategy.RUnlock()
	s, has := strategys[id]
	return s, has
}

//Has has
func Has(id string) bool {
	lockerStategy.RLock()
	defer lockerStategy.RUnlock()

	_, has := strategys[id]
	return has
}
func loadStategy() {
	strategies, def, e := dao_strategy.GetAllStrategy()
	if e != nil {
		return
	}
	reset(strategies, def)
}

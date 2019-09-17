package api_manager

import (
	"sync"

	"github.com/eolinker/goku/goku-node/manager/updater"
	dao_api "github.com/eolinker/goku/server/dao/node-mysql/dao-api"
	entity "github.com/eolinker/goku/server/entity/node-entity"
)

func init() {
	updater.Add(load, 5, "goku_gateway_api")
}

// apiList: api列表全局变量
var (
	apiList map[int]*entity.Api
	locker  sync.RWMutex
)

//GetAPI
func GetAPI(id int) (*entity.Api, bool) {
	locker.RLock()
	defer locker.RUnlock()
	a, has := apiList[id]
	return a, has
}

//GetAPI
func GetAPIs(ids []int) []*entity.Api {

	apis := make([]*entity.Api, 0, len(ids))

	locker.RLock()
	defer locker.RUnlock()
	for _, id := range ids {
		if api, has := apiList[id]; has {
			apis = append(apis, api)
		}
	}

	return apis
}

func reset(list map[int]*entity.Api) {
	locker.Lock()
	defer locker.Unlock()
	apiList = list
}

func load() {
	apis, e := dao_api.GetAllApi()
	if e != nil {
		return
	}
	reset(apis)
}

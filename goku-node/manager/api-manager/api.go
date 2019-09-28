package apimanager

import (
	"sync"

	"github.com/eolinker/goku-api-gateway/goku-node/manager/updater"
	dao_api "github.com/eolinker/goku-api-gateway/server/dao/node-mysql/dao-api"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

func init() {
	updater.Add(load, 5, "goku_gateway_api")
}

// apiList: api列表全局变量
var (
	apiList map[int]*entity.API
	locker  sync.RWMutex
)

//GetAPI 通过id获取API信息
func GetAPI(id int) (*entity.API, bool) {
	locker.RLock()
	defer locker.RUnlock()
	a, has := apiList[id]
	return a, has
}

//GetAPIs 获取接口信息列表
func GetAPIs(ids []int) []*entity.API {

	apis := make([]*entity.API, 0, len(ids))

	locker.RLock()
	defer locker.RUnlock()
	for _, id := range ids {
		if api, has := apiList[id]; has {
			apis = append(apis, api)
		}
	}

	return apis
}

func reset(list map[int]*entity.API) {
	locker.Lock()
	defer locker.Unlock()
	apiList = list
}

func load() {
	apis, e := dao_api.GetAllAPI()
	if e != nil {
		return
	}
	reset(apis)
}

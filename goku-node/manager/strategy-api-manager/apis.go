package strategyapimanager

import entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"

type _APIMap struct {
	apis map[int]*entity.StrategyAPI
}

func (m *_APIMap) Get(id int) (*entity.StrategyAPI, bool) {
	api, has := m.apis[id]
	return api, has
}

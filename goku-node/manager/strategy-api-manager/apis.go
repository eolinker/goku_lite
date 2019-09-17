package strategy_api_manager

import entity "github.com/eolinker/goku/server/entity/node-entity"

type _ApiMap struct {
	apis map[int]*entity.StrategyApi
}

func (m *_ApiMap) Get(id int) (*entity.StrategyApi, bool) {
	api, has := m.apis[id]
	return api, has
}

package strategy

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
)

var(
	strategyDao dao.StrategyDao
	strategyGroupDao dao.StrategyGroupDao
	strategyPluginDao dao.StrategyPluginDao

)

func init() {
	pdao.Need(&strategyDao,&strategyGroupDao,&strategyPluginDao)
}
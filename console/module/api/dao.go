package api

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
)

var(
	apiDao dao.APIDao
	apiGroupDao dao.APIGroupDao
	apiPluginDao dao.APIPluginDao
	apiStrategyDao dao.APIStrategyDao
	importDao dao.ImportDao
)

func init() {
	pdao.Need(&apiDao,&apiGroupDao,&apiPluginDao,&apiStrategyDao,&importDao)
}
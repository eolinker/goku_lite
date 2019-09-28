package pluginmanager

import (
	"github.com/eolinker/goku-api-gateway/goku-node/manager/updater"
	dao_plugin "github.com/eolinker/goku-api-gateway/server/dao/node-mysql/dao-plugin"
)

func init() {
	updater.Add(load, 1, "goku_plugin")
}

func load() {
	pis, e := dao_plugin.GetAll()
	if e != nil {
		return
	}
	reset(pis)
}

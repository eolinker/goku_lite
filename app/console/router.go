package main

import (
	"net/http"

	account_default "github.com/eolinker/goku-api-gateway/console/account"
	"github.com/eolinker/goku-api-gateway/console/controller/account"
	"github.com/eolinker/goku-api-gateway/console/controller/api"
	"github.com/eolinker/goku-api-gateway/console/controller/auth"
	"github.com/eolinker/goku-api-gateway/console/controller/balance"
	"github.com/eolinker/goku-api-gateway/console/controller/cluster"
	config_log "github.com/eolinker/goku-api-gateway/console/controller/config-log"
	"github.com/eolinker/goku-api-gateway/console/controller/discovery"
	"github.com/eolinker/goku-api-gateway/console/controller/gateway"
	"github.com/eolinker/goku-api-gateway/console/controller/monitor"
	"github.com/eolinker/goku-api-gateway/console/controller/node"
	"github.com/eolinker/goku-api-gateway/console/controller/plugin"
	"github.com/eolinker/goku-api-gateway/console/controller/project"
	"github.com/eolinker/goku-api-gateway/console/controller/strategy"
	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"
)

var (
	accountFactory = goku_handler.NewAccountHandlerFactory(account_default.NewDefaultAccount())
)

func router() http.Handler {
	s := goku_handler.NewGokuServer(accountFactory)

	// 账号管理模块
	s.Add("/guest", account.NewAccountController())
	s.Add("/user", account.NewUserController())

	// 接口管理模块
	s.Add("/apis", api.NewAPIHandlers())
	s.Add("/apis/group", api.NewGroupHandlers())
	s.Add("/import/ams", api.NewImportHandlers())
	s.Add("/plugin/api", api.NewPluginHandlers())

	// 鉴权模块
	s.Add("/auth", auth.NewHandlers())

	// 负载模块
	s.Add("/balance", balance.NewHandlers())

	// 集群模块
	s.Add("/cluster", cluster.NewHandlers())
	s.Add("/version/config", cluster.NewVersionHandlers())

	// 日志配置模块
	s.Add("/config/log", config_log.NewHandlers())

	// 服务发现模块
	s.Add("/balance/service", discovery.NewHandlers())

	// 网关模块
	s.Add("/monitor/gateway", gateway.NewHandlers())

	// 监控模块
	s.Add("/monitor/module/config", monitor.NewHandlers())

	// 节点模块
	s.Add("/node", node.NewNodeHandlers())
	s.Add("/node/group", node.NewGroupHandlers())

	// 插件模块
	s.Add("/plugin", plugin.NewHandlers())

	// 项目模块
	s.Add("/project", project.NewHandlers())

	// 策略模块
	s.Add("/strategy", strategy.NewStrategyHandlers())
	s.Add("/strategy/group", strategy.NewGroupHandlers())
	s.Add("/strategy/api", strategy.NewAPIStrategyHandlers())
	s.Add("/plugin/strategy", strategy.NewPluginHandlers())

	// 前端接入
	s.Add("/", new(staticHandlers))
	return s
}

type staticHandlers struct {
}

func (s *staticHandlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/": http.StripPrefix("/", http.FileServer(http.Dir("./static"))),
	}
}

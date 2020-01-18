package plugin_executor

import (
	"time"

	"github.com/eolinker/goku-api-gateway/config"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
	goku_plugin "github.com/eolinker/goku-plugin"
)

//Executor executor
type Executor interface {
	Execute(ctx *common.Context) (isContinue bool, e error)
	IsStop() bool
	IsAuth() bool
}
type executorInfo struct {
	Name   string
	isStop bool
	isAuth bool
}

func (ex *executorInfo) IsStop() bool {
	return ex.isStop
}

func (ex *executorInfo) IsAuth() bool {
	return ex.isAuth
}

func genExecutor(cfg *config.PluginConfig) executorInfo {
	return executorInfo{
		Name:   cfg.Name,
		isStop: cfg.IsStop,
		isAuth: cfg.IsAuth,
	}
}

type beforeExecutor struct {
	executorInfo
	plugin goku_plugin.PluginBeforeMatch
}

//Execute execute
func (ex *beforeExecutor) Execute(ctx *common.Context) (isContinue bool, e error) {
	requestID := ctx.RequestId()
	ctx.SetPlugin(ex.Name)
	log.Debug(requestID, " before plugin :", ex.Name, " start")
	now := time.Now()
	isContinue, err := ex.plugin.BeforeMatch(ctx)
	log.Debug(requestID, " before plugin :", ex.Name, " Duration:", time.Since(now))
	log.Debug(requestID, " before plugin :", ex.Name, " end")
	if err != nil {
		log.Warn(requestID, " before plugin:", ex.Name, " error:", err)
	}
	return isContinue, err
}

//NewBeforeExecutor 创建before阶段执行器
func NewBeforeExecutor(cfg *config.PluginConfig, p goku_plugin.PluginBeforeMatch) *beforeExecutor {
	return &beforeExecutor{
		executorInfo: genExecutor(cfg),
		plugin:       p,
	}

}

type accessExecutor struct {
	executorInfo
	plugin goku_plugin.PluginAccess
}

func (ex *accessExecutor) Execute(ctx *common.Context) (isContinue bool, e error) {
	requestID := ctx.RequestId()
	ctx.SetPlugin(ex.Name)

	log.Debug(requestID, " access plugin :", ex.Name, " start")
	now := time.Now()
	isContinue, err := ex.plugin.Access(ctx)
	log.Debug(requestID, " access plugin :", ex.Name, " Duration:", time.Since(now))
	log.Debug(requestID, " access plugin :", ex.Name, " end")
	if err != nil {
		log.Warn(requestID, " access plugin:", ex.Name, " error:", err)
	}
	return isContinue, err
}

//NewAccessExecutor 创建access阶段执行器
func NewAccessExecutor(cfg *config.PluginConfig, p goku_plugin.PluginAccess) *accessExecutor {
	return &accessExecutor{
		executorInfo: genExecutor(cfg),
		plugin:       p,
	}
}

type proxyExecutor struct {
	executorInfo
	plugin goku_plugin.PluginProxy
}

//Execute execute
func (ex *proxyExecutor) Execute(ctx *common.Context) (isContinue bool, e error) {
	requestID := ctx.RequestId()
	ctx.SetPlugin(ex.Name)

	log.Debug(requestID, " proxy plugin :", ex.Name, " start")
	now := time.Now()
	isContinue, err := ex.plugin.Proxy(ctx)
	log.Debug(requestID, " proxy plugin :", ex.Name, " Duration:", time.Since(now))
	log.Debug(requestID, " proxy plugin :", ex.Name, " end")
	if err != nil {
		log.Warn(requestID, " proxy plugin:", ex.Name, " error:", err)
	}
	return isContinue, err
}

//NewProxyExecutor 创建proxy阶段执行器
func NewProxyExecutor(cfg *config.PluginConfig, p goku_plugin.PluginProxy) *proxyExecutor {
	return &proxyExecutor{
		executorInfo: genExecutor(cfg),
		plugin:       p,
	}
}

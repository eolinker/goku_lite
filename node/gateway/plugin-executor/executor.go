package plugin_executor

import (
	goku_plugin "github.com/eolinker/goku-plugin"
	"github.com/eolinker/goku-api-gateway/config"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
	"time"
)

type Executor interface {
	Execute(ctx *common.Context)(isContinue bool, e error)
	IsStop()bool
}
type executorInfo struct {
	Name string
	isStop bool
}



func (ex*executorInfo) IsStop() bool {
	return  ex.isStop
}

func genExecutor(cfg *config.PluginConfig) executorInfo {
	return executorInfo{
		Name:cfg.Name,
		isStop:cfg.IsStop,
	}
}
type beforeExecutor struct {
	executorInfo
	plugin goku_plugin.PluginBeforeMatch

}

func (ex*beforeExecutor) Execute(ctx *common.Context) (isContinue bool, e error) {
	requestId:=ctx.RequestId()
	ctx.SetPlugin(ex.Name)
	log.Debug(requestId, " before plugin :", ex.Name, " start")
	now := time.Now()
	isContinue, err := ex.plugin.BeforeMatch(ctx)
	log.Debug(requestId, " before plugin :", ex.Name, " Duration:", time.Since(now))
	log.Debug(requestId, " before plugin :", ex.Name, " end")
	if err != nil {
		log.Warn(requestId, " before plugin:", ex.Name, " error:", err)
	}
	return isContinue,err
}

func NewBeforeExecutor(cfg  *config.PluginConfig,p goku_plugin.PluginBeforeMatch ) *beforeExecutor {
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
	requestId:=ctx.RequestId()
	ctx.SetPlugin(ex.Name)

	log.Debug(requestId, " access plugin :", ex.Name, " start")
	now := time.Now()
	isContinue, err := ex.plugin.Access(ctx)
	log.Debug(requestId, " access plugin :", ex.Name, " Duration:", time.Since(now))
	log.Debug(requestId, " access plugin :", ex.Name, " end")
	if err != nil {
		log.Warn(requestId, " access plugin:", ex.Name, " error:", err)
	}
	return isContinue,err
}

func NewAccessExecutor(cfg  *config.PluginConfig,p goku_plugin.PluginAccess ) *accessExecutor {
	return &accessExecutor{
		executorInfo: genExecutor(cfg),
		plugin:       p,
	}
}

type proxyExecutor struct {
	executorInfo
	plugin goku_plugin.PluginProxy
}

func (ex*proxyExecutor) Execute(ctx *common.Context) (isContinue bool, e error) {
	requestId:=ctx.RequestId()
	ctx.SetPlugin(ex.Name)

	log.Debug(requestId, " proxy plugin :", ex.Name, " start")
	now := time.Now()
	isContinue, err := ex.plugin.Proxy(ctx)
	log.Debug(requestId, " proxy plugin :", ex.Name, " Duration:", time.Since(now))
	log.Debug(requestId, " proxy plugin :", ex.Name, " end")
	if err != nil {
		log.Warn(requestId, " proxy plugin:", ex.Name, " error:", err)
	}
	return isContinue,err
}

func NewProxyExecutor(cfg  *config.PluginConfig,p goku_plugin.PluginProxy ) *proxyExecutor {
	return &proxyExecutor{
		executorInfo: genExecutor(cfg),
		plugin:       p,
	}
}
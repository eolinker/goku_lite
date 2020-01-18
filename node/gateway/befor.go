package gateway

import (
	"net/http"
	"time"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/goku-node/common"
	plugin_executor "github.com/eolinker/goku-api-gateway/node/gateway/plugin-executor"
	"github.com/eolinker/goku-api-gateway/node/utils"
)

//Before before
type Before struct {
	pluginBefor       []plugin_executor.Executor
	pluginGlobalBefor []plugin_executor.Executor

	strategies        map[string]*Strategy
	anonymousStrategy string
}

//Router 路由
func (r *Before) Router(w http.ResponseWriter, req *http.Request, ctx *common.Context) {
	start := time.Now()
	isBefore := r.BeforeMatch(ctx)
	log.Info(ctx.RequestId(), " BeforeMatch plugin duration:", time.Since(start))
	if !isBefore {
		log.Info(ctx.RequestId(), " stop by BeforeMatch plugin")
		return
	}
	r.rout(w, req, ctx)
}

//BeforeMatch 插件流程，匹配URI及策略前执行
func (r *Before) BeforeMatch(ctx *common.Context) bool {
	requestID := ctx.RequestId()
	defer func(ctx *common.Context) {
		log.Debug(requestID, " before plugin default: begin")
		for _, handler := range r.pluginGlobalBefor {

			_, _ = handler.Execute(ctx)

		}
		log.Debug(requestID, " before plugin default: end")
	}(ctx)
	log.Debug(requestID, " before plugin : begin")
	for _, handler := range r.pluginBefor {

		flag, _ := handler.Execute(ctx)

		if flag == false {
			if handler.IsStop() == true {
				return false
			}
		}
	}
	log.Debug(requestID, " before plugin : end")
	return true

}

func (r *Before) rout(w http.ResponseWriter, req *http.Request, ctx *common.Context) {
	strategyID := utils.GetStrateyID(ctx)

	if strategyID == "" {
		// 没有策略id
		if r.anonymousStrategy == "" {
			// 没有开放策略
			go log.Info("[ERROR]Missing Strategy ID!")
			ctx.SetStatus(500, "500")

			ctx.SetBody([]byte("[ERROR]Missing Strategy ID!"))
			return
		}

		strategyID = r.anonymousStrategy
	}

	v, has := r.strategies[strategyID]

	if !has {
		go log.Info("[ERROR]StrategyID dose not exist!")

		ctx.SetStatus(500, "500")

		ctx.SetBody([]byte("[ERROR]StrategyID dose not exist!"))
		return
	}
	v.Router(w, req, ctx)

}

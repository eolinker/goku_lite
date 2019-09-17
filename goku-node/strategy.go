package goku_node

import (
	log "github.com/eolinker/goku/goku-log"
	"github.com/eolinker/goku/goku-node/common"
	strategy_manager "github.com/eolinker/goku/goku-node/manager/strategy-manager"
)

func getStrateyID(ctx *common.Context) string {
	if value := ctx.Request().GetHeader("Strategy-Id"); value != "" {
		return value
	}
	if value := ctx.Request().URL().Query().Get("Strategy-Id"); value != "" {
		return value
	}
	if value := ctx.Request().GetForm("Strategy-Id"); value != "" {
		return value
	}

	return ""
}

func retrieveStrategyID(ctx *common.Context) (string, bool) {
	strategyID := getStrateyID(ctx)

	if strategyID == "" {
		if v, ok := strategy_manager.GetAnonymous(); ok {
			strategyID = v.StrategyID
			ctx.SetStrategyName(v.StrategyName)
			ctx.SetStrategyId(strategyID)

		} else {
			go log.Info("[ERROR]Missing Strategy ID!")
			ctx.SetStatus(500, "500")

			ctx.SetBody([]byte("[ERROR]Missing Strategy ID!"))
			return strategyID, false
		}
	} else {

		// 快速查找策略ID
		if v, ok := strategy_manager.Get(strategyID); ok {

			if v.EnableStatus != 1 {

				go log.Info("[ERROR]StrategyID is out of service!")

				ctx.SetStatus(500, "500")

				ctx.SetBody([]byte("[ERROR]StrategyID is out of service!"))
				return strategyID, false
			}

			ctx.SetStrategyName(v.StrategyName)
			ctx.SetStrategyId(strategyID)
		} else {
			go log.Info("[ERROR]StrategyID dose not exist!")

			ctx.SetStatus(500, "500")

			ctx.SetBody([]byte("[ERROR]StrategyID dose not exist!"))
			return strategyID, false
		}
	}
	return strategyID, true
}

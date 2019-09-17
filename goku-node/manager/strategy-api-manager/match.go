package strategy_api_manager

import (
	"github.com/eolinker/goku/utils"
	"strings"

	api_manager "github.com/eolinker/goku/goku-node/manager/api-manager"
	node_common "github.com/eolinker/goku/goku-node/node-common"
	entity "github.com/eolinker/goku/server/entity/node-entity"
)

func CheckApiFromStrategy(strategyId, requestPath string, requestMethod string) (*entity.ApiExtend, string, []string, bool) {
	apiMap, has := getAPis(strategyId)
	if !has {
		return nil, "", nil, false
	}
	for _, strategyApi := range apiMap.apis {
		apiInfo, has := api_manager.GetAPI(strategyApi.ApiId)
		if !has {
			continue
		}
		if apiInfo.StripSlash {
			requestPath = node_common.FilterSlash(requestPath)
			apiInfo.RequestURL = node_common.FilterSlash(apiInfo.RequestURL)
		}

		isMatch, splitURL, param := node_common.MatchURI(requestPath, apiInfo.RequestURL)
		if isMatch {
			method := strings.ToUpper(apiInfo.RequestMethod)
			if strings.Contains(method, requestMethod) {
				apiextend := &entity.ApiExtend{Api: apiInfo}

				if strategyApi.Target != "" {
					apiextend.Target = strategyApi.Target
				} else {
					apiextend.Target = apiInfo.BalanceName
				}

				apiextend.Target = 	utils.TrimSuffixAll(apiextend.Target,"/")

				//apiextend.TargetServer = balance_manager.ParseTargetServer(apiextend.Target)
				return apiextend, splitURL, param, true
			}
		}
	}
	return nil, "", nil, false
}

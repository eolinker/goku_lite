package strategyapimanager

import (
	"github.com/eolinker/goku-api-gateway/utils"
	"strings"

	api_manager "github.com/eolinker/goku-api-gateway/goku-node/manager/api-manager"
	node_common "github.com/eolinker/goku-api-gateway/goku-node/node-common"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

//CheckAPIFromStrategy 判断接口是否在策略中
func CheckAPIFromStrategy(strategyID, requestPath string, requestMethod string) (*entity.APIExtend, string, []string, bool) {
	apiMap, has := getAPis(strategyID)
	if !has {
		return nil, "", nil, false
	}
	for _, strategyAPI := range apiMap.apis {
		apiInfo, has := api_manager.GetAPI(strategyAPI.APIID)
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
				apiextend := &entity.APIExtend{API: apiInfo}

				if strategyAPI.Target != "" {
					apiextend.Target = strategyAPI.Target
				} else {
					apiextend.Target = apiInfo.BalanceName
				}

				apiextend.Target = utils.TrimSuffixAll(apiextend.Target, "/")

				return apiextend, splitURL, param, true
			}
		}
	}
	return nil, "", nil, false
}

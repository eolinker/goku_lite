package dao

import (
	"goku-ce/server/conf"
	"gopkg.in/yaml.v2"
)

// 编辑鉴权信息
func EditAuth(strategyConfPath,strategyID,auth,basicUserName,basicUserPassword,apiKey string) bool {
	_,strategy := conf.ParseStrategyInfo(strategyConfPath)
	strategy[strategyID].Auth = auth
	strategy[strategyID].BasicUserName = basicUserName
	strategy[strategyID].BasicUserPassword = basicUserPassword
	strategy[strategyID].ApiKey = apiKey
	
	strategyList := make([]*conf.StrategyInfo,0)
	for _,value := range strategy {
		strategyList = append(strategyList,value)
	}

	content, err :=  yaml.Marshal(conf.Strategy{
		StrategyList: strategyList,
	})
	if err != nil {
		return false
	}
	conf.WriteConfigToFile(strategyConfPath,content)
	return true
}

// 获取鉴权信息
func GetAuthInfo(strategyConfPath,strategyID string) map[string]string {
	_,strategy := conf.ParseStrategyInfo(strategyConfPath)
	_,ok := strategy[strategyID]
	if !ok {
		return map[string]string{}
	}
	auth := "none"
	if strategy[strategyID].Auth != "" {
		auth =  strategy[strategyID].Auth
	}
	return map[string]string{
		"auth" : auth,
		"basicUserName" : strategy[strategyID].BasicUserName,
		"basicUserPassword" : strategy[strategyID].BasicUserPassword,
		"apiKey" : strategy[strategyID].ApiKey,
	}
}


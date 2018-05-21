package dao

import (
	"strings"
	"goku-ce/server/conf"
	"gopkg.in/yaml.v2"
	"os"
)

// 修改网关黑白名单
func EditGatewayIPList(gatewayAlias,ipLimitType,ipWhiteList,ipBlackList string) bool {
	gateway := conf.ParseGatewayInfo(conf.GlobalConf.GatewayConfPath)
	ipWhiteList = strings.Replace(ipWhiteList,"；",";",-1)
	ipBlackList = strings.Replace(ipBlackList,"；",";",-1)
	whiteList := strings.Split(ipWhiteList,";")
	blackList := strings.Split(ipBlackList,";")
	gateway[gatewayAlias].IPLimitType = ipLimitType
	gateway[gatewayAlias].IPBlackList = blackList
	gateway[gatewayAlias].IPWhiteList = whiteList
	gatewayConf,err := yaml.Marshal(gateway[gatewayAlias])
	if err != nil {
		panic(err)
	}
	pthSep := string(os.PathSeparator)
	gatewayDir := conf.GlobalConf.GatewayConfPath + pthSep + gatewayAlias
	conf.WriteConfigToFile(gatewayDir + pthSep + "gateway.conf",gatewayConf)
	return true
}

// 修改策略组黑白名单
func EditStrategyIPList(strategyConfPath,strategyID,ipLimitType,ipWhiteList,ipBlackList string) bool {
	_,strategy := conf.ParseStrategyInfo(strategyConfPath)
	ipWhiteList = strings.Replace(ipWhiteList,"；",";",-1)
	ipBlackList = strings.Replace(ipBlackList,"；",";",-1)
	whiteList := strings.Split(ipWhiteList,";")
	blackList := strings.Split(ipBlackList,";")
	strategy[strategyID].IPLimitType = ipLimitType
	strategy[strategyID].IPBlackList = blackList
	strategy[strategyID].IPWhiteList = whiteList

	strategyList := make([]*conf.StrategyInfo,0)
	for _,value := range strategy {
		strategyList = append(strategyList,value)
	}
	strategyConf := conf.Strategy{}
	strategyConf.StrategyList = strategyList
	content, err :=  yaml.Marshal(strategyConf)
	if err != nil {
		return false
	}
	conf.WriteConfigToFile(strategyConfPath,content)
	return true
}

// 获取网关黑白名单
func GetGatewayIPList(gatewayAlias string) map[string]string {
	gateway := conf.ParseGatewayInfo(conf.GlobalConf.GatewayConfPath)
	value,ok := gateway[gatewayAlias]
	if !ok {
		return make(map[string]string)
	} else {
		ipLimitType := "none"
		if value.IPLimitType != "" {
			ipLimitType = value.IPLimitType
		}
		return map[string]string{
			"ipLimitType" : ipLimitType,
			"ipWhiteList" : strings.Join(value.IPWhiteList,";"),
			"ipBlackList" : strings.Join(value.IPBlackList,";"),
		}
	}
}

// 获取策略组黑白名单
func GetStrategyIPList(strategyConfPath,strategyID string) map[string]string {
	_,strategy := conf.ParseStrategyInfo(strategyConfPath)
	value,ok := strategy[strategyID]
	if !ok {
		return make(map[string]string)
	} else {
		ipLimitType := "none"
		if value.IPLimitType != "" {
			ipLimitType = value.IPLimitType
		}
		return map[string]string{
			"ipLimitType" : ipLimitType,
			"ipWhiteList" : strings.Join(value.IPWhiteList,";"),
			"ipBlackList" : strings.Join(value.IPBlackList,";"),
		}
	}
}
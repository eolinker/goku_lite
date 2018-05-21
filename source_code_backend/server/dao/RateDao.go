package dao

import (
	"goku-ce/server/conf"
	"gopkg.in/yaml.v2"
)

// 新增流量限制
func AddRateLimit(strategyConfPath,strategyID,period string,startTime,endTime,priority,limitCount int,allow bool) bool {
	rateLimitList,_,_ := conf.ParseRateLimitInfo(strategyConfPath,strategyID)
	maxID := 1 
	for _,r := range rateLimitList {
		if r.LimitID > maxID {
			maxID = r.LimitID
		}
	}
	rateLimitInfo := &conf.RateLimitInfo{
		LimitID: maxID + 1,
 		Allow : allow,
		Period : period,
		Limit : limitCount,
		Priority: priority,
		StartTime : startTime,
		EndTime : endTime,
	}
	rateLimitList = append(rateLimitList,rateLimitInfo)

	_,strategy := conf.ParseStrategyInfo(strategyConfPath)
	strategy[strategyID].RateLimitList = rateLimitList
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

// 修改流量限制
func EditRateLimit(strategyConfPath,strategyID,period string,rateLimitID,startTime,endTime,priority,limitCount int,allow bool) bool {
	rates,rateLimit,_ := conf.ParseRateLimitInfo(strategyConfPath,strategyID)
	_,ok := rateLimit[rateLimitID]
	if !ok {
		return false
	}
	rateLimit[rateLimitID] = &conf.RateLimitInfo{
		Allow : allow,
		Period : period,
		Limit : limitCount,
		Priority: priority,
		StartTime : startTime,
		EndTime : endTime,
	}

	rateLimitList := make([]*conf.RateLimitInfo,0)
	for i:= 0; i< len(rates);i++ {
		rateLimitList = append(rateLimitList,rateLimit[i+1])
	}

	_,strategy := conf.ParseStrategyInfo(strategyConfPath)
	strategy[strategyID].RateLimitList = rateLimitList
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

// 删除流量限制
func DeleteRateLimit(strategyConfPath,strategyID string,rateLimitID int) bool {
	_,rateLimit,_ := conf.ParseRateLimitInfo(strategyConfPath,strategyID)
	_,ok := rateLimit[rateLimitID]
	if !ok {
		return false
	}
	delete(rateLimit,rateLimitID)
	rateLimitList := make([]*conf.RateLimitInfo,0)
	for _,r := range rateLimit {
		rateLimitList = append(rateLimitList,r)
	}
	_,strategy := conf.ParseStrategyInfo(strategyConfPath)
	strategy[strategyID].RateLimitList = rateLimitList
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

// 获取流量限制列表
func GetRateLimitInfo(strategyConfPath,strategyID string,limitID int) (bool,*conf.RateLimitInfo) {
	_,rateLimit,_ := conf.ParseRateLimitInfo(strategyConfPath,strategyID)
	value,ok := rateLimit[limitID]
	if !ok {
		return false,&conf.RateLimitInfo{}
	}
	return true,value
}

// 获取流量限制列表
func GetRateLimitList(strategyConfPath,strategyID string) []map[string]interface{}{
	_,_,rateLimitList := conf.ParseRateLimitInfo(strategyConfPath,strategyID)
	return rateLimitList
}

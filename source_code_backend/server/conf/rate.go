package conf

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type RateLimitInfo struct {
	LimitID					int						`json:"limitID" yaml:"limitI_id"`
	Allow					bool					`json:"allow" yaml:"allow"`
	Period					string					`json:"period" yaml:"period"`
	Limit					int						`json:"limit" yaml:"limit"`
	Priority				int						`json:"priority" yaml:"priority"`
	StartTime				int						`json:"startTime" yaml:"start_time"`
	EndTime					int						`json:"endTime" yaml:"end_time"`
}

func ParseRateLimitInfo(path,strategyID string) ([]*RateLimitInfo,map[int]*RateLimitInfo,[]map[string]interface{}) {
	rateInfo := make(map[int]*RateLimitInfo)
	mapRateList := make([]map[string]interface{},0)
	rateList :=  make([]*RateLimitInfo,0)
	var strategy Strategy
	content,err := ioutil.ReadFile(path)
	if err != nil {
		return rateList,rateInfo,mapRateList
	}

	err = yaml.Unmarshal(content,&strategy)
	if err != nil {
		panic(err)
	}
	
	for _,s := range strategy.StrategyList {
		if s.StrategyID == strategyID {
			if len(s.RateLimitList) != 0 {
				rateList = s.RateLimitList
				maxID := 0
				for _,r := range s.RateLimitList {
					if r.LimitID > maxID {
						maxID = r.LimitID
					}
				}
				for _,r := range s.RateLimitList{
					limitID := r.LimitID
					if r.LimitID == 0 {
						limitID = maxID + 1 
					}
					rate := map[string]interface{}{
						"limitID": limitID,
						"allow": r.Allow,
						"period": r.Period,
						"limit": r.Limit,
						"priority": r.Priority,
						"startTime": r.StartTime,
						"endTime": r.EndTime,
					}
					rateInfo[limitID] = r
					maxID += 1
					mapRateList = append(mapRateList,rate)
				}
			}
			break
		}
	}
	return rateList,rateInfo,mapRateList
}
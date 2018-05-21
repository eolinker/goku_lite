package conf

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"sort"
	"time"
)

type Strategy struct {
	StrategyList				[]*StrategyInfo			`json:"strategy" yaml:"strategy"`
}

type StrategyInfo struct {
	StrategyName			string					`json:"strategyName" yaml:"strategy_name"`
	StrategyID				string					`json:"strategyID" yaml:"strategy_id"`
	Auth					string					`json:"auth,omitempty" yaml:"auth,omitempty"`
	BasicUserName			string					`json:"basicUserName,omitempty" yaml:"basic_user_name,omitempty"`
	BasicUserPassword		string					`json:"basicUserPassword,omitempty" yaml:"basic_user_password,omitempty"`
	ApiKey					string					`json:"apiKey,omitempty" yaml:"api_key"`
	IPLimitType				string					`json:"ipLimitType,omitempty" yaml:"ip_limit_type,omitempty"`
	IPWhiteList				[]string				`json:"ipWhiteList,omitempty" yaml:"ip_white_list,omitempty"`
	IPBlackList				[]string				`json:"ipBlackList,omitempty" yaml:"ip_black_list,omitempty"`
	RateLimitList			[]*RateLimitInfo		`json:"rateLimitList,omitempty" yaml:"rate_limit_list,omitempty"`
	UpdateTime				string					`json:"updateTime,omitempty" yaml:"update_time,omitempty"`	
	CreateTime				string					`json:"createTime,omitempty" yaml:"create_time,omitempty"`
}

type StrategySlice []map[string]interface{}

func (s StrategySlice) Len() int {    // 重写 Len() 方法
    return len(s)
}
func (s StrategySlice) Swap(i, j int){     // 重写 Swap() 方法
    s[i], s[j] = s[j], s[i]
}
func (s StrategySlice) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
	t1, t1Err := time.Parse("2006-01-02 15:04:05", s[j]["updateTime"].(string))
	t2, t2Err := time.Parse("2006-01-02 15:04:05", s[i]["updateTime"].(string))

	if t1Err == nil && t2Err != nil {
		return true
	} else if t1Err == nil && t2Err == nil{
		if t1.Before(t2) {
			return false
		} else {
			return true
		}
	} else if t1Err != nil && t2Err == nil {
		return false
	} else {
		str := []string{s[i]["strategyName"].(string),s[j]["strategyName"].(string)}
		sort.Strings(str)
		if str[0] == s[i]["strategyName"].(string) {
			return false
		}else {
			return true
		}
	}
}

// 读入策略组信息
func ParseStrategyInfo(path string) ([]*StrategyInfo,map[string]*StrategyInfo) {
	strategyInfo := make(map[string]*StrategyInfo)
	strategyList := make([]*StrategyInfo,0)
	var strategy Strategy
	content,err := ioutil.ReadFile(path)
	if err != nil {
		return strategyList,strategyInfo
	}

	err = yaml.Unmarshal(content,&strategy)
	if err != nil {
		panic(err)
	}
	
	if len(strategy.StrategyList) != 0 {
		strategyList = strategy.StrategyList
	}

	for _,s := range strategy.StrategyList {
		strategyInfo[s.StrategyID] = s
	}
	return strategyList,strategyInfo
}

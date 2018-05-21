package dao

import (
	"goku-ce/server/conf"
	"gopkg.in/yaml.v2"
	"goku-ce/utils"
	"time"
	"sort"
)

// 新增策略组
func AddStrategy(strategyConfPath,strategyName string) (bool,string) {
	_,strategy := conf.ParseStrategyInfo(strategyConfPath)
	strategyID := ""
	for i := 0;i<6;i++ {
		randomID := utils.GetRandomString(6)
		_,ok := strategy[randomID]
		if !ok {
			strategyID = randomID
			break
		}
	}

	if strategyID == "" {
		return false,""
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	strategy[strategyID] = &conf.StrategyInfo{
		StrategyID : strategyID,
		StrategyName : strategyName,
		UpdateTime : now,
		CreateTime : now,
	}

	strategyList := make([]*conf.StrategyInfo,0)
	for _,value := range strategy {
		strategyList = append(strategyList,value)
	}

	content, err :=  yaml.Marshal(conf.Strategy{
		StrategyList: strategyList,
	})
	if err != nil {
		return false,""
	}
	conf.WriteConfigToFile(strategyConfPath,content)
	return true,strategyID
}

// 修改策略组
func EditStrategy(strategyConfPath,strategyName,strategyID string) (bool) {
	_,strategy := conf.ParseStrategyInfo(strategyConfPath)
	value,ok := strategy[strategyID]
	if !ok {
		return false
	} else {
		now := time.Now().Format("2006-01-02 15:04:05")
		value.StrategyName = strategyName
		value.UpdateTime = now
	}

	strategyList := make([]*conf.StrategyInfo,0)
	for _,value := range strategy {
		strategyList = append(strategyList,value)
	}

	content, err :=  yaml.Marshal(conf.Strategy{
		StrategyList: strategyList,
	})
	if err != nil {
		panic(err);
	}
	
	conf.WriteConfigToFile(strategyConfPath,content)
	return true
}


// 删除策略组
func DeleteStrategy(strategyConfPath,strategyID string) (bool) {
	_,strategy := conf.ParseStrategyInfo(strategyConfPath)

	_,ok := strategy[strategyID]
	if !ok {
		return false
	} else {
		delete(strategy,strategyID)
	}
	
	strategyList := make([]*conf.StrategyInfo,0)
	for _,value := range strategy {
		strategyList = append(strategyList,value)
	}

	content, err :=  yaml.Marshal(conf.Strategy{
		StrategyList: strategyList,
	})
	if err != nil {
		panic(err);
	}
	conf.WriteConfigToFile(strategyConfPath,content)
	return true
}


// 获取策略组列表
func GetStrategyList(strategyConfPath string) []map[string]interface{}{
	strategy,_ := conf.ParseStrategyInfo(strategyConfPath)
	strategyList := make([]map[string]interface{},0)
	for _,s := range strategy {
		strategyInfo := map[string]interface{}{
			"strategyID":s.StrategyID,
			"strategyName":s.StrategyName,
			"updateTime":s.UpdateTime,
		}
		strategyList = append(strategyList,strategyInfo)
	}
	sort.Sort(sort.Reverse(conf.StrategySlice(strategyList)))
	return strategyList
}

// 获取简易策略组列表
func GetSimpleStrategyList(strategyConfPath string) []map[string]interface{}{
	strategy,_ := conf.ParseStrategyInfo(strategyConfPath)
	strategyList := make([]map[string]interface{},0)
	for _,s := range strategy {
		strategyInfo := map[string]interface{}{
			"strategyID":s.StrategyID,
			"strategyName":s.StrategyName,
		}
		strategyList = append(strategyList,strategyInfo)
	}
	sort.Sort(sort.Reverse(conf.StrategySlice(strategyList)))
	return strategyList
}

// 获取策略组数量
func GetStrategyCount(strategyConfPath string) int {
	startegyList,_ := conf.ParseStrategyInfo(strategyConfPath)
	return len(startegyList)
}
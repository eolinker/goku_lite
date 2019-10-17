package strategy

import (
	"errors"

	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddStrategy 新增策略组
func AddStrategy(strategyName string, groupID int) (bool, string, error) {
	flag := console_sqlite3.CheckIsOpenGroup(groupID)
	if flag {
		return false, "[ERROR]The group is an open group", errors.New("[ERROR]The group is an open group")
	}
	flag, result, err := console_sqlite3.AddStrategy(strategyName, groupID)

	return flag, result, err
}

//EditStrategy 修改策略组信息
func EditStrategy(strategyID, strategyName string, groupID int) (bool, string, error) {
	return console_sqlite3.EditStrategy(strategyID, strategyName, groupID)
}

// DeleteStrategy 删除策略组
func DeleteStrategy(strategyID string) (bool, string, error) {
	flag := console_sqlite3.CheckIsOpenStrategy(strategyID)
	if flag {
		return false, "[ERROR]The strategy is an open strategy", errors.New("[ERROR]The strategy is an open strategy")
	}
	flag, result, err := console_sqlite3.DeleteStrategy(strategyID)

	return flag, result, err
}

// GetOpenStrategy 获取策略组列表
func GetOpenStrategy() (bool, *entity.Strategy, error) {
	return console_sqlite3.GetOpenStrategy()
}

// GetStrategyList 获取策略组列表
func GetStrategyList(groupID int, keyword string, condition int) (bool, []*entity.Strategy, error) {
	return console_sqlite3.GetStrategyList(groupID, keyword, condition)
}

// GetStrategyInfo 获取策略组信息
func GetStrategyInfo(strategyID string) (bool, *entity.Strategy, error) {
	return console_sqlite3.GetStrategyInfo(strategyID)
}

// CheckStrategyIsExist 检查策略组ID是否存在
func CheckStrategyIsExist(strategyID string) (bool, error) {
	return console_sqlite3.CheckStrategyIsExist(strategyID)
}

// BatchEditStrategyGroup 批量修改策略组分组
func BatchEditStrategyGroup(strategyIDList string, groupID int) (bool, string, error) {
	return console_sqlite3.BatchEditStrategyGroup(strategyIDList, groupID)
}

// BatchDeleteStrategy 批量修改策略组分组
func BatchDeleteStrategy(strategyIDList string) (bool, string, error) {
	flag, result, err := console_sqlite3.BatchDeleteStrategy(strategyIDList)

	return flag, result, err
}

//CheckIsOpenStrategy 检查是否是开放策略
func CheckIsOpenStrategy(strategyID string) bool {
	return console_sqlite3.CheckIsOpenStrategy(strategyID)
}

//BatchUpdateStrategyEnableStatus 更新策略启用状态
func BatchUpdateStrategyEnableStatus(strategyIDList string, enableStatus int) (bool, string, error) {
	flag, result, err := console_sqlite3.BatchUpdateStrategyEnableStatus(strategyIDList, enableStatus)

	return flag, result, err
}

// GetBalanceListInStrategy 获取在策略中的负载列表
func GetBalanceListInStrategy(strategyID string, balanceType int) (bool, []string, error) {
	return console_sqlite3.GetBalanceListInStrategy(strategyID, balanceType)
}

// CopyStrategy 复制策略
func CopyStrategy(strategyID string, newStrategyID string, userID int) (string, error) {
	result, err := console_sqlite3.CopyStrategy(strategyID, newStrategyID, userID)

	return result, err
}

//GetStrategyIDList 获取策略ID列表
func GetStrategyIDList(groupID int, keyword string, condition int) (bool, []string, error) {
	return console_sqlite3.GetStrategyIDList(groupID, keyword, condition)
}

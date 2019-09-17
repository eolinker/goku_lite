package strategy

import (
	"errors"

	"github.com/eolinker/goku/server/dao"
	console_mysql "github.com/eolinker/goku/server/dao/console-mysql"
	entity "github.com/eolinker/goku/server/entity/console-entity"
)

// 新增策略组
func AddStrategy(strategyName string, groupID int) (bool, string, error) {
	flag := console_mysql.CheckIsOpenGroup(groupID)
	if flag {
		return false, "[ERROR]The group is an open group", errors.New("[ERROR]The group is an open group")
	} else {
		flag, result, err := console_mysql.AddStrategy(strategyName, groupID)
		if flag {
			tableName := "goku_gateway_strategy"
			dao.UpdateTable(tableName)
		}
		return flag, result, err
	}
}

// 修改策略组信息
func EditStrategy(strategyID, strategyName string, groupID int) (bool, string, error) {
	return console_mysql.EditStrategy(strategyID, strategyName, groupID)
}

// DeleteStrategy 删除策略组
func DeleteStrategy(strategyID string) (bool, string, error) {
	flag := console_mysql.CheckIsOpenStrategy(strategyID)
	if flag {
		return false, "[ERROR]The strategy is an open strategy", errors.New("[ERROR]The strategy is an open strategy")
	} else {
		tableName := "goku_gateway_strategy"
		flag, result, err := console_mysql.DeleteStrategy(strategyID)
		if flag {
			dao.UpdateTable(tableName)
		}
		return flag, result, err
	}
}

// GetOpenStrategy 获取策略组列表
func GetOpenStrategy() (bool, *entity.Strategy, error) {
	return console_mysql.GetOpenStrategy()
}

// GetStrategyList 获取策略组列表
func GetStrategyList(groupID int, keyword string, condition int) (bool, []*entity.Strategy, error) {
	return console_mysql.GetStrategyList(groupID, keyword, condition)
}

// GetStrategyInfo 获取策略组信息
func GetStrategyInfo(strategyID string) (bool, *entity.Strategy, error) {
	return console_mysql.GetStrategyInfo(strategyID)
}

// CheckStrategyIsExist 检查策略组ID是否存在
func CheckStrategyIsExist(strategyID string) (bool, error) {
	return console_mysql.CheckStrategyIsExist(strategyID)
}

// BatchEditStrategyGroup 批量修改策略组分组
func BatchEditStrategyGroup(strategyIDList string, groupID int) (bool, string, error) {
	return console_mysql.BatchEditStrategyGroup(strategyIDList, groupID)
}

// BatchDeleteStrategy 批量修改策略组分组
func BatchDeleteStrategy(strategyIDList string) (bool, string, error) {
	flag, result, err := console_mysql.BatchDeleteStrategy(strategyIDList)
	if flag {
		tableName := "goku_gateway_strategy"
		dao.UpdateTable(tableName)
	}
	return flag, result, err
}

func CheckIsOpenStrategy(strategyID string) bool {
	return console_mysql.CheckIsOpenStrategy(strategyID)
}

// 更新策略启用状态
func BatchUpdateStrategyEnableStatus(strategyIDList string, enableStatus int) (bool, string, error) {
	tableName := "goku_gateway_strategy"
	flag, result, err := console_mysql.BatchUpdateStrategyEnableStatus(strategyIDList, enableStatus)
	if flag {
		dao.UpdateTable(tableName)
		console_mysql.BatchUpdateStrategyPluginUpdateTag(strategyIDList)
		console_mysql.BatchUpdateApiPluginUpdateTag(strategyIDList)
	}
	return flag, result, err
}

// UpdateAllStrategyPluginUpdateTag 更新所有策略插件更新标识
func UpdateAllStrategyPluginUpdateTag() error {
	return console_mysql.UpdateAllStrategyPluginUpdateTag()
}

// GetBalanceListInStrategy 获取在策略中的负载列表
func GetBalanceListInStrategy(strategyID string, balanceType int) (bool, []string, error) {
	return console_mysql.GetBalanceListInStrategy(strategyID, balanceType)
}

// CopyStrategy 复制策略
func CopyStrategy(strategyID string, newStrategyID string, userID int) (string, error) {
	result, err := console_mysql.CopyStrategy(strategyID, newStrategyID, userID)
	if err == nil {
		dao.UpdateTable("goku_gateway_strategy")
		dao.UpdateTable("goku_conn_strategy_api")
		dao.UpdateTable("goku_conn_plugin_strategy")
		dao.UpdateTable("goku_conn_plugin_api")
	}
	return result, err
}

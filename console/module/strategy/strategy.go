package strategy

import (
	"errors"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//AddStrategy 新增策略组
func AddStrategy(strategyName string, groupID, userID int) (bool, string, error) {
	flag := strategyGroupDao.CheckIsOpenGroup(groupID)
	if flag {
		return false, "[ERROR]The group is an open group", errors.New("[ERROR]The group is an open group")
	}
	flag, result, err := strategyDao.AddStrategy(strategyName, groupID, userID)

	return flag, result, err
}

//EditStrategy 修改策略组信息
func EditStrategy(strategyID, strategyName string, groupID, userID int) (bool, string, error) {
	return strategyDao.EditStrategy(strategyID, strategyName, groupID, userID)
}

// DeleteStrategy 删除策略组
func DeleteStrategy(strategyID string) (bool, string, error) {
	flag := strategyDao.CheckIsOpenStrategy(strategyID)
	if flag {
		return false, "[ERROR]The strategy is an open strategy", errors.New("[ERROR]The strategy is an open strategy")
	}
	flag, result, err := strategyDao.DeleteStrategy(strategyID)

	return flag, result, err
}

// GetOpenStrategy 获取策略组列表
func GetOpenStrategy() (bool, *entity.Strategy, error) {
	return strategyDao.GetOpenStrategy()
}

// GetStrategyList 获取策略组列表
func GetStrategyList(groupID int, keyword string, condition, page, pageSize int) (bool, []*entity.Strategy, int, error) {
	return strategyDao.GetStrategyList(groupID, keyword, condition, page, pageSize)
}

// GetStrategyInfo 获取策略组信息
func GetStrategyInfo(strategyID string) (bool, *entity.Strategy, error) {
	return strategyDao.GetStrategyInfo(strategyID)
}

// CheckStrategyIsExist 检查策略组ID是否存在
func CheckStrategyIsExist(strategyID string) (bool, error) {
	return strategyDao.CheckStrategyIsExist(strategyID)
}

// BatchEditStrategyGroup 批量修改策略组分组
func BatchEditStrategyGroup(strategyIDList string, groupID int) (bool, string, error) {
	return strategyDao.BatchEditStrategyGroup(strategyIDList, groupID)
}

// BatchDeleteStrategy 批量修改策略组分组
func BatchDeleteStrategy(strategyIDList string) (bool, string, error) {
	flag, result, err := strategyDao.BatchDeleteStrategy(strategyIDList)

	return flag, result, err
}

//BatchUpdateStrategyEnableStatus 更新策略启用状态
func BatchUpdateStrategyEnableStatus(strategyIDList string, enableStatus int) (bool, string, error) {
	flag, result, err := strategyDao.BatchUpdateStrategyEnableStatus(strategyIDList, enableStatus)

	return flag, result, err
}

// GetBalanceListInStrategy 获取在策略中的负载列表
func GetBalanceListInStrategy(strategyID string, balanceType int) (bool, []string, error) {
	return strategyDao.GetBalanceListInStrategy(strategyID, balanceType)
}

// CopyStrategy 复制策略
func CopyStrategy(strategyID string, newStrategyID string, userID int) (string, error) {
	result, err := strategyDao.CopyStrategy(strategyID, newStrategyID, userID)

	return result, err
}

//GetStrategyIDList 获取策略ID列表
func GetStrategyIDList(groupID int, keyword string, condition int) (bool, []string, error) {
	return strategyDao.GetStrategyIDList(groupID, keyword, condition)
}

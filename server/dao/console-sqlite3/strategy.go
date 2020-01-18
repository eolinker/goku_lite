package console_sqlite3

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/eolinker/goku-api-gateway/server/dao"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
	"github.com/eolinker/goku-api-gateway/utils"
)

//StrategyDao StrategyDao
type StrategyDao struct {
	db *sql.DB
}

//NewStrategyDao new StrategyDao
func NewStrategyDao() *StrategyDao {
	return &StrategyDao{}
}

//Create create
func (d *StrategyDao) Create(db *sql.DB) (interface{}, error) {
	d.db = db
	var i dao.StrategyDao = d

	return &i, nil
}

//AddStrategy  新增策略组
func (d *StrategyDao) AddStrategy(strategyName string, groupID, userID int) (bool, string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	// 随机生成字符串
	sqlCode := "SELECT strategyID FROM goku_gateway_strategy WHERE strategyID = ?"
	var strategyID string
	for i := 0; i < 5; i++ {
		randomStr := utils.GetRandomString(6)
		err := db.QueryRow(sqlCode, randomStr).Scan(&strategyID)
		if err == nil {
			continue
		}
		strategyID = randomStr
		break
	}
	if strategyID == "" {
		return false, "[ERROR]Empty strategy id !", nil
	}
	stmt, err := db.Prepare(`INSERT INTO goku_gateway_strategy (strategyID,strategyName,updateTime,createTime,groupID) VALUES (?,?,?,?,?)`)

	if err != nil {
		return false, "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(strategyID, strategyName, now, now, groupID)
	if err != nil {
		return false, "[ERROR]Failed to insert data!", err
	}
	return true, strategyID, nil
}

//EditStrategy 修改策略组信息
func (d *StrategyDao) EditStrategy(strategyID, strategyName string, groupID, userID int) (bool, string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.Prepare(`UPDATE goku_gateway_strategy SET strategyName = ?,groupID = ?,updateTime = ? WHERE strategyID = ?`)
	if err != nil {
		return false, "[ERROR]Illegal SQL statement!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(strategyName, groupID, now, strategyID)
	if err != nil {
		return false, "[ERROR]Failed to update data!", err
	}
	return true, strategyID, nil
}

//DeleteStrategy 删除策略组
func (d *StrategyDao) DeleteStrategy(strategyID string) (bool, string, error) {
	db := d.db
	Tx, _ := db.Begin()
	_, err := Tx.Exec(`DELETE FROM goku_gateway_strategy WHERE strategyID = ?;`, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Failed to delete data!", err
	}
	// 删除绑定的接口
	sql := "DELETE FROM goku_conn_strategy_api WHERE strategyID = ?;"
	_, err = Tx.Exec(sql, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	// 删除策略组插件
	sql = "DELETE FROM goku_conn_plugin_strategy WHERE strategyID = ?;"
	_, err = Tx.Exec(sql, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	// 删除接口插件
	sql = "DELETE FROM goku_conn_plugin_api WHERE strategyID = ?;"
	_, err = Tx.Exec(sql, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	Tx.Commit()
	return true, "", nil
}

// GetStrategyList 获取策略组列表
func (d *StrategyDao) GetStrategyList(groupID int, keyword string, condition, page, pageSize int) (bool, []*entity.Strategy, int, error) {
	rule := make([]string, 0, 2)

	rule = append(rule, "A.strategyType != 1")
	if groupID > -1 {
		groupRule := fmt.Sprintf("A.groupID = %d", groupID)
		rule = append(rule, groupRule)
	}
	if keyword != "" {
		searchRule := "(A.strategyID LIKE '%" + keyword + "%' OR A.strategyName LIKE '%" + keyword + "%')"
		rule = append(rule, searchRule)
	}
	if condition > 0 {
		filterRule := fmt.Sprintf("A.enableStatus = %d", condition-1)
		rule = append(rule, filterRule)
	}
	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}
	sql := fmt.Sprintf("SELECT A.strategyID,A.strategyName,IFNULL(A.updateTime,''),IFNULL(A.createTime,''),A.enableStatus,A.groupID,IFNULL(B.groupName,'未分组') FROM goku_gateway_strategy A LEFT JOIN goku_gateway_strategy_group B ON A.groupID = B.groupID %s", ruleStr)
	count := getCountSQL(d.db, sql)
	rows, err := getPageSQL(d.db, sql, "A.updateTime", "DESC", page, pageSize)
	if err != nil {
		return false, make([]*entity.Strategy, 0), 0, err
	}
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	strategyList := make([]*entity.Strategy, 0)
	for rows.Next() {
		var strategy entity.Strategy
		err = rows.Scan(&strategy.StrategyID, &strategy.StrategyName, &strategy.UpdateTime, &strategy.CreateTime, &strategy.EnableStatus, &strategy.GroupID, &strategy.GroupName)
		if err != nil {
			return false, make([]*entity.Strategy, 0), 0, err
		}
		strategyList = append(strategyList, &strategy)
	}
	return true, strategyList, count, nil
}

// GetOpenStrategy 获取策略组列表
func (d *StrategyDao) GetOpenStrategy() (bool, *entity.Strategy, error) {
	var openStrategy entity.Strategy
	db := d.db
	sql := `SELECT strategyID,strategyName,IFNULL(updateTime,""),IFNULL(createTime,""),enableStatus,strategyType FROM goku_gateway_strategy WHERE strategyType = 1 ORDER BY updateTime DESC;`
	err := db.QueryRow(sql).Scan(&openStrategy.StrategyID, &openStrategy.StrategyName, &openStrategy.UpdateTime, &openStrategy.CreateTime, &openStrategy.EnableStatus, &openStrategy.StrategyType)
	if err != nil {
		return false, nil, err
	}
	openStrategy.GroupName = "开放分组"
	return true, &openStrategy, nil
}

//GetStrategyInfo 获取策略组信息
func (d *StrategyDao) GetStrategyInfo(strategyID string) (bool, *entity.Strategy, error) {
	db := d.db
	sql := `SELECT strategyID,strategyName,IFNULL(updateTime,''),strategyType,enableStatus FROM goku_gateway_strategy WHERE strategyID = ?;`
	strategy := new(entity.Strategy)
	err := db.QueryRow(sql, strategyID).Scan(&strategy.StrategyID, &strategy.StrategyName, &strategy.UpdateTime, &strategy.StrategyType, &strategy.EnableStatus)
	if err != nil {
		return false, nil, err
	}
	return true, strategy, err
}

//CheckStrategyIsExist 检查策略组ID是否存在
func (d *StrategyDao) CheckStrategyIsExist(strategyID string) (bool, error) {
	db := d.db
	sql := "SELECT strategyID FROM goku_gateway_strategy WHERE strategyID = ?;"
	var id string
	err := db.QueryRow(sql, strategyID).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, err
}

//BatchEditStrategyGroup 批量修改策略组分组
func (d *StrategyDao) BatchEditStrategyGroup(strategyIDList string, groupID int) (bool, string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	strategy := strings.Split(strategyIDList, ",")
	code := make([]string, 0, len(strategy))
	s := make([]interface{}, 0, len(strategy)+2)
	s = append(s, groupID, now)
	for i := 0; i < len(strategy); i++ {
		code = append(code, "?")
		s = append(s, strategy[i])
	}
	sql := "UPDATE goku_gateway_strategy SET groupID = ?,updateTime = ? WHERE strategyID IN (" + strings.Join(code, ",") + ");"
	_, err := Tx.Exec(sql, s...)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	Tx.Commit()
	return true, "", nil
}

//BatchDeleteStrategy 批量修改策略组
func (d *StrategyDao) BatchDeleteStrategy(strategyIDList string) (bool, string, error) {
	db := d.db
	Tx, _ := db.Begin()
	strategy := strings.Split(strategyIDList, ",")
	code := ""
	s := make([]interface{}, 0)
	for i := 0; i < len(strategy); i++ {
		code += "?"
		if i < len(strategy)-1 {
			code += ","
		}
		s = append(s, strategy[i])
	}
	// 删除绑定的接口
	sql := "DELETE FROM goku_conn_strategy_api WHERE strategyID IN (" + code + ");"
	_, err := Tx.Exec(sql, s...)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	sql = "DELETE FROM goku_gateway_strategy WHERE strategyID IN (" + code + ");"
	_, err = Tx.Exec(sql, s...)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	// 删除策略组插件
	sql = "DELETE FROM goku_conn_plugin_strategy WHERE strategyID IN (" + code + ");"
	_, err = Tx.Exec(sql, s...)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	// 删除接口插件
	sql = "DELETE FROM goku_conn_plugin_api WHERE strategyID IN (" + code + ");"
	_, err = Tx.Exec(sql, s...)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	Tx.Commit()
	return true, "", nil
}

//CheckIsOpenStrategy 判断是否是开放策略
func (d *StrategyDao) CheckIsOpenStrategy(strategyID string) bool {
	db := d.db
	var strategyType int
	sql := "SELECT strategyType FROM goku_gateway_strategy WHERE strategyID = ?;"
	err := db.QueryRow(sql, strategyID).Scan(&strategyType)
	if err != nil {
		return false
	}
	if strategyType == 1 {
		return true
	}
	return false
}

//BatchUpdateStrategyEnableStatus 更新策略启动状态
func (d *StrategyDao) BatchUpdateStrategyEnableStatus(strategyIDList string, enableStatus int) (bool, string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	strategy := strings.Split(strategyIDList, ",")
	code := ""
	s := make([]interface{}, 0)
	s = append(s, enableStatus)
	s = append(s, now)
	for i := 0; i < len(strategy); i++ {
		code += "?"
		if i < len(strategy)-1 {
			code += ","
		}
		s = append(s, strategy[i])
	}
	sql := "UPDATE goku_gateway_strategy SET enableStatus = ?,updateTime = ? WHERE strategyID IN (" + code + ");"
	result, err := db.Exec(sql, s...)
	if err != nil {
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	affectRow, err := result.RowsAffected()
	if err != nil || affectRow == 0 {
		return false, "[ERROR]Fail to update data!", err
	}
	return true, "", nil
}

// GetBalanceListInStrategy 获取在策略中的负载列表
func (d *StrategyDao) GetBalanceListInStrategy(strategyID string, balanceType int) (bool, []string, error) {
	db := d.db

	sql := "SELECT DISTINCT(IFNULL(A.balanceName,'')) FROM goku_gateway_api A INNER JOIN goku_conn_strategy_api B ON A.apiID = B.apiID WHERE B.strategyID = ?;"
	if balanceType == 1 {
		sql = "SELECT DISTINCT(IFNULL(target,'')) FROM goku_conn_strategy_api WHERE strategyID = ?;"
	}
	rows, err := db.Query(sql, strategyID)
	if err != nil {
		return false, nil, err
	}
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	balanceNameList := make([]string, 0, 10)
	for rows.Next() {
		var balanceName string
		err = rows.Scan(&balanceName)
		if err != nil {
			return false, nil, err
		}
		if balanceName == "" {
			continue
		}
		balanceNameList = append(balanceNameList, balanceName)
	}
	return true, balanceNameList, nil
}

// CopyStrategy 复制策略
func (d *StrategyDao) CopyStrategy(strategyID string, newStrategyID string, userID int) (string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "INSERT INTO goku_conn_strategy_api (strategyID,apiID,apiMonitorStatus,strategyMonitorStatus,target,updateTime) SELECT ?,apiID,apiMonitorStatus,strategyMonitorStatus,target,? FROM goku_conn_strategy_api WHERE strategyID = ?"
	_, err := db.Exec(sql, newStrategyID, now, strategyID)
	if err != nil {
		return "", err
	}
	sql = "INSERT INTO goku_conn_plugin_strategy (strategyID,pluginName,pluginConfig,pluginInfo,createTime,updateTime,pluginStatus,updateTag,updaterID) SELECT ?,pluginName,pluginConfig,pluginInfo,?,?,pluginStatus,updateTag,? FROM goku_conn_plugin_strategy WHERE strategyID = ?"
	_, err = db.Exec(sql, newStrategyID, now, now, userID, strategyID)
	if err != nil {
		return "", err
	}

	sql = "INSERT INTO goku_conn_plugin_api (strategyID,apiID,pluginName,pluginConfig,pluginInfo,createTime,updateTime,pluginStatus,updateTag,updaterID) SELECT ?,apiID,pluginName,pluginConfig,pluginInfo,?,?,pluginStatus,updateTag,? FROM goku_conn_plugin_api WHERE strategyID = ?"
	_, err = db.Exec(sql, newStrategyID, now, now, userID, strategyID)
	if err != nil {
		return "", err
	}
	return newStrategyID, nil
}

//GetStrategyIDList 获取策略ID列表
func (d *StrategyDao) GetStrategyIDList(groupID int, keyword string, condition int) (bool, []string, error) {
	db := d.db
	rule := make([]string, 0, 2)

	rule = append(rule, fmt.Sprintf("A.strategyType != 1"))
	if groupID > -1 {
		groupRule := fmt.Sprintf("A.groupID = %d", groupID)
		rule = append(rule, groupRule)
	}
	if keyword != "" {
		searchRule := "(A.strategyID LIKE '%" + keyword + "%' OR A.strategyName LIKE '%" + keyword + "%')"
		rule = append(rule, searchRule)
	}
	if condition > 0 {
		filterRule := fmt.Sprintf("A.enableStatus = %d", condition-1)
		rule = append(rule, filterRule)
	}
	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}
	sql := fmt.Sprintf("SELECT A.strategyID FROM goku_gateway_strategy A %s ORDER BY A.updateTime DESC;", ruleStr)
	rows, err := db.Query(sql)
	if err != nil {
		return false, nil, err
	}
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	strategyList := make([]string, 0)
	for rows.Next() {
		var strategyID string
		err = rows.Scan(&strategyID)
		if err != nil {
			return false, nil, err
		}
		strategyList = append(strategyList, strategyID)
	}
	return true, strategyList, nil
}

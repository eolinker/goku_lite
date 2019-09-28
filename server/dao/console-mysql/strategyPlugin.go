package consolemysql

import (
	SQL "database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	database2 "github.com/eolinker/goku-api-gateway/common/database"
)

var strategyPlugins = []string{"goku-oauth2_auth", "goku-rate_limiting", "goku-replay_attack_defender"}

//AddPluginToStrategy 新增策略组插件
func AddPluginToStrategy(pluginName, config, strategyID string) (bool, interface{}, error) {
	db := database2.GetConnection()
	// 查询接口是否添加该插件
	sql := "SELECT strategyID FROM goku_conn_plugin_strategy WHERE strategyID = ? AND pluginName = ?;"
	var id string
	err := db.QueryRow(sql, strategyID, pluginName).Scan(&id)
	if err == nil {
		return false, "[ERROR]The strategy plugin is already exist", errors.New("[ERROR]The strategy plugin is already exist")
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	result, err := Tx.Exec("INSERT INTO goku_conn_plugin_strategy (pluginName,pluginConfig,strategyID,createTime,updateTime,pluginStatus) VALUES (?,?,?,?,?,?);", pluginName, config, strategyID, now, now, 1)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to insert data", errors.New("[ERROR]Fail to insert data")
	}
	connID, err := result.LastInsertId()
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to insert data", errors.New("[ERROR]Fail to insert data")
	}

	sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?;"
	_, err = Tx.Exec(sql, now, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data!", err
	}
	Tx.Commit()
	return true, connID, nil
}

//EditStrategyPluginConfig 新增策略组插件配置
func EditStrategyPluginConfig(pluginName, config, strategyID string) (bool, string, error) {
	db := database2.GetConnection()
	// 查询策略组是否添加该插件
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	updateTag := t.Format("20060102150405")
	sql := "SELECT strategyID FROM goku_conn_plugin_strategy WHERE strategyID = ? AND pluginName = ?;"
	var id string
	err := db.QueryRow(sql, strategyID, pluginName).Scan(&id)
	if err != nil {
		return false, "[ERROR]The strategy plugin is not exist", errors.New("[ERROR]The strategy plugin is not exist")
	}
	Tx, _ := db.Begin()
	_, err = Tx.Exec("UPDATE goku_conn_plugin_strategy SET updateTag = ?,pluginConfig = ?,updateTime = ? WHERE pluginName = ? AND strategyID = ?;", updateTag, config, now, pluginName, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data", errors.New("[ERROR]Fail to update data")
	}

	sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?;"
	_, err = Tx.Exec(sql, now, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data!", err
	}
	Tx.Commit()
	return true, "", nil
}

// GetStrategyPluginList 获取策略插件列表
func GetStrategyPluginList(strategyID, keyword string, condition int) (bool, []map[string]interface{}, error) {
	db := database2.GetConnection()

	rule := make([]string, 0, 3)

	rule = append(rule, fmt.Sprintf("A.strategyID = '%s'", strategyID))
	if keyword != "" {
		searchRule := "(A.pluginName LIKE '%" + keyword + "%' OR B.pluginDesc LIKE '%" + keyword + "%')"
		rule = append(rule, searchRule)
	}
	if condition > 0 {
		rule = append(rule, fmt.Sprintf("IF(B.pluginStatus=0,-1,A.pluginStatus) = %d", condition-1))
	}

	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}
	sql := fmt.Sprintf(`SELECT A.connID,A.pluginName,A.pluginConfig,IFNULL(A.createTime,""),IFNULL(A.updateTime,""),B.pluginPriority,IF(B.pluginStatus=0,-1,A.pluginStatus) as pluginStatus,IFNULL(B.pluginDesc,"") FROM goku_conn_plugin_strategy A INNER JOIN goku_plugin B ON B.pluginName = A.pluginName %s ORDER BY pluginStatus DESC,A.updateTime DESC;`, ruleStr)
	rows, err := db.Query(sql)
	if err != nil {
		return false, make([]map[string]interface{}, 0), err
	}
	defer rows.Close()
	pluginList := make([]map[string]interface{}, 0)
	//获取记录列
	for rows.Next() {
		var pluginPriority, pluginStatus, connID int
		var pluginName, pluginDesc, pluginConfig, createTime, updateTime string
		err = rows.Scan(&connID, &pluginName, &pluginConfig, &createTime, &updateTime, &pluginPriority, &pluginStatus, &pluginDesc)
		if err != nil {
		}
		pluginInfo := map[string]interface{}{
			"connID":         connID,
			"pluginName":     pluginName,
			"pluginConfig":   pluginConfig,
			"pluginPriority": pluginPriority,
			"pluginStatus":   pluginStatus,
			"createTime":     createTime,
			"updateTime":     updateTime,
			"pluginDesc":     pluginDesc,
		}
		pluginList = append(pluginList, pluginInfo)
	}
	return true, pluginList, nil
}

//GetStrategyPluginConfig 通过策略组ID获取配置信息
func GetStrategyPluginConfig(strategyID, pluginName string) (bool, string, error) {
	db := database2.GetConnection()
	sql := "SELECT pluginConfig FROM goku_conn_plugin_strategy WHERE strategyID = ? AND pluginName = ?;"
	var p string
	err := db.QueryRow(sql, strategyID, pluginName).Scan(&p)
	if err != nil {
		if err == SQL.ErrNoRows {
			return false, "", errors.New("[ERROR]Can not find the plugin")
		}
		return false, "", err
	}
	return true, p, nil
}

//CheckPluginIsExistInStrategy 检查策略组是否绑定插件
func CheckPluginIsExistInStrategy(strategyID, pluginName string) (bool, error) {
	db := database2.GetConnection()
	sql := "SELECT strategyID FROM goku_conn_plugin_strategy WHERE strategyID = ? AND pluginName = ?;"
	var id string
	err := db.QueryRow(sql, strategyID, pluginName).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, err
}

//GetStrategyPluginStatus 检查策略组插件是否开启
func GetStrategyPluginStatus(strategyID, pluginName string) (bool, error) {
	db := database2.GetConnection()
	sql := "SELECT pluginStatus FROM goku_conn_plugin_strategy WHERE strategyID = ? AND pluginName = ?;"
	var pluginStatus int
	err := db.QueryRow(sql, strategyID, pluginName).Scan(&pluginStatus)
	if err != nil {
		return false, err
	}
	if pluginStatus != 1 {
		return false, nil
	}
	return true, nil
}

// GetConnIDFromStrategyPlugin 获取ConnID
func GetConnIDFromStrategyPlugin(pluginName, strategyID string) (bool, int, error) {
	db := database2.GetConnection()
	sql := "SELECT connID FROM goku_conn_plugin_strategy WHERE strategyID = ? AND pluginName = ?;"
	var connID int
	err := db.QueryRow(sql, strategyID, pluginName).Scan(&connID)
	if err != nil {
		return false, 0, err
	}
	return true, connID, nil
}

//BatchEditStrategyPluginStatus 批量修改策略组插件状态
func BatchEditStrategyPluginStatus(connIDList, strategyID string, pluginStatus int) (bool, string, error) {
	db := database2.GetConnection()
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	updateTag := t.Format("20060102150405")
	Tx, _ := db.Begin()
	sql := "UPDATE goku_conn_plugin_strategy SET updateTag = ?,pluginStatus = ?, updateTime = ? WHERE connID IN (" + connIDList + ");"
	_, err := Tx.Exec(sql, updateTag, pluginStatus, now)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?;"
	_, err = Tx.Exec(sql, now, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data!", err
	}
	Tx.Commit()
	return true, "", nil
}

//BatchDeleteStrategyPlugin 批量删除策略组插件
func BatchDeleteStrategyPlugin(connIDList, strategyID string) (bool, string, error) {
	db := database2.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	sql := "DELETE FROM goku_conn_plugin_strategy WHERE connID IN (" + connIDList + ");"
	_, err := Tx.Exec(sql)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?;"
	_, err = Tx.Exec(sql, now, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data!", err
	}
	Tx.Commit()
	return true, "", nil
}

//CheckStrategyPluginIsExistByConnIDList 通过connIDList判断插件是否存在
func CheckStrategyPluginIsExistByConnIDList(connIDList, pluginName string) (bool, error) {
	db := database2.GetConnection()
	sql := "SELECT pluginStatus FROM goku_conn_plugin_strategy WHERE connID IN (" + connIDList + ") AND pluginName = ?;"
	var pluginStatus int
	err := db.QueryRow(sql, pluginName).Scan(&pluginStatus)
	if err != nil {
		return false, err
	}
	return true, nil
}

//UpdateStrategyTagByPluginName 更新策略插件标志位
func UpdateStrategyTagByPluginName(strategyID string, pluginNameList string) error {
	db := database2.GetConnection()
	plugins := strings.Split(pluginNameList, ",")
	code := make([]string, 0, len(plugins))
	updateTag := time.Now().Format("20060102150405")
	s := make([]interface{}, 0, len(plugins)+2)
	s = append(s, updateTag, strategyID)
	for i := 0; i < len(plugins); i++ {
		code = append(code, "?")
		s = append(s, plugins[i])
	}
	sql := "UPDATE goku_conn_plugin_strategy SET updateTag = ? WHERE strategyID = ? AND pluginName IN (" + strings.Join(code, ",") + ");"
	log.Debug("UpdateStrategyTagByPluginName:", strategyID, pluginNameList, sql, s)
	_, err := db.Exec(sql, s...)
	if err != nil {
		return err
	}
	return nil
}

//BatchUpdateStrategyPluginUpdateTag 批量更新策略插件更新标志位
func BatchUpdateStrategyPluginUpdateTag(strategyIDList string) error {
	db := database2.GetConnection()
	strategyIDs := strings.Split(strategyIDList, ",")
	strategyCode := make([]string, 0, len(strategyIDs))

	code := make([]string, 0, len(strategyPlugins))
	updateTag := time.Now().Format("20060102150405")
	s := make([]interface{}, 0, len(strategyPlugins)+1+len(strategyIDs))
	s = append(s, updateTag)
	for i := 0; i < len(strategyIDs); i++ {
		strategyCode = append(strategyCode, "?")
		s = append(s, strategyIDs[i])
	}
	for i := 0; i < len(strategyPlugins); i++ {
		code = append(code, "?")
		s = append(s, strategyPlugins[i])
	}
	sql := "UPDATE goku_conn_plugin_strategy SET updateTag = ? WHERE strategyID IN (" + strings.Join(strategyCode, ",") + ") AND pluginName IN (" + strings.Join(code, ",") + ");"
	log.Debug("BatchUpdateStrategyPluginUpdateTag:", sql, " ", s)
	_, err := db.Exec(sql, s...)
	if err != nil {
		return err
	}
	return nil
}

//UpdateAllStrategyPluginUpdateTag 更新策略插件更新标志位
func UpdateAllStrategyPluginUpdateTag() error {
	db := database2.GetConnection()
	updateTag := time.Now().Format("20060102150405")

	sql := "UPDATE goku_conn_plugin_strategy SET updateTag = ?;"
	_, err := db.Exec(sql, updateTag)
	if err != nil {
		return err
	}
	return nil
}

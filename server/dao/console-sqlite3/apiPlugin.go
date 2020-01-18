package console_sqlite3

import (
	SQL "database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/eolinker/goku-api-gateway/server/dao"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	"time"
)

//APIPluginDao APIPluginDao
type APIPluginDao struct {
	db *SQL.DB
}

//NewAPIPluginDao new APIPluginDao
func NewAPIPluginDao() *APIPluginDao {
	return &APIPluginDao{}
}

//Create create
func (d *APIPluginDao) Create(db *SQL.DB) (interface{}, error) {
	d.db = db
	var i dao.APIPluginDao = d
	return &i, nil
}

//AddPluginToAPI 新增接口插件
func (d *APIPluginDao) AddPluginToAPI(pluginName, config, strategyID string, apiID, userID int) (bool, interface{}, error) {
	db := d.db
	// 查询接口是否添加该插件
	sql := "SELECT apiID FROM goku_conn_plugin_api WHERE strategyID = ? AND pluginName = ? AND apiID = ?;"
	var id int
	err := db.QueryRow(sql, strategyID, pluginName, apiID).Scan(&id)
	if err == nil {
		return false, "[ERROR]The api plugin is already exist", errors.New("[ERROR]The api plugin is already exist")
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	result, err := Tx.Exec("INSERT INTO goku_conn_plugin_api (pluginName,pluginConfig,strategyID,apiID,updateTime,createTime,pluginStatus,updaterID) VALUES (?,?,?,?,?,?,?,?);", pluginName, config, strategyID, apiID, now, now, 1, userID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to insert data", errors.New("[ERROR]Fail to insert data")
	}
	connID, _ := result.LastInsertId()
	sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?;"
	_, err = Tx.Exec(sql, now, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Failed to update data!", err
	}
	Tx.Commit()
	return true, int(connID), nil
}

//EditAPIPluginConfig 修改接口插件配置
func (d *APIPluginDao) EditAPIPluginConfig(pluginName, config, strategyID string, apiID, userID int) (bool, interface{}, error) {
	db := d.db
	// 查询接口是否添加该插件
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")

	sql := "SELECT connID,apiID FROM goku_conn_plugin_api WHERE strategyID = ? AND pluginName = ? AND apiID = ?;"
	var id, aID int
	err := db.QueryRow(sql, strategyID, pluginName, apiID).Scan(&id, &aID)
	if err != nil {
		return false, "[ERROR]The api plugin is not exist", errors.New("[ERROR]The api plugin is not exist")
	}
	updateTag := t.Format("20060102150405")
	Tx, _ := db.Begin()
	_, err = Tx.Exec("UPDATE goku_conn_plugin_api SET updateTag = ?,pluginConfig = ?,updateTime = ?,updaterID = ? WHERE strategyID = ? AND apiID = ? AND pluginName = ?;", updateTag, config, now, userID, strategyID, apiID, pluginName)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data", errors.New("[ERROR]Fail to update data")
	}

	sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?;"
	_, err = Tx.Exec(sql, now, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Failed to update data!", err
	}
	Tx.Commit()
	return true, id, nil
}

//GetAPIPluginList 获取接口插件列表
func (d *APIPluginDao) GetAPIPluginList(apiID int, strategyID string) (bool, []map[string]interface{}, error) {
	db := d.db
	sql := `SELECT goku_conn_plugin_api.connID,goku_conn_plugin_api.pluginName,IFNULL(goku_conn_plugin_api.createTime,""),IFNULL(goku_conn_plugin_api.updateTime,""),goku_conn_plugin_api.pluginConfig,goku_plugin.pluginPriority, IF(goku_plugin.pluginStatus=0,-1,goku_conn_plugin_api.pluginStatus) as pluginStatus,goku_gateway_api.requestURL FROM goku_conn_plugin_api INNER JOIN goku_plugin ON goku_plugin.pluginName = goku_conn_plugin_api.pluginName INNER goku_gateway_api.apiID = goku_conn_plugin_api.apiID WHERE goku_conn_plugin_api.apiID = ? AND goku_conn_plugin_api.strategyID = ? ORDER BY pluginStatus DESC,goku_conn_plugin_api.updateTime DESC;`
	rows, err := db.Query(sql, apiID, strategyID)
	if err != nil {
		return false, nil, err
	}
	defer rows.Close()
	pluginList := make([]map[string]interface{}, 0)
	//获取记录列

	for rows.Next() {
		var pluginPriority, pluginStatus, connID int
		var pluginName, pluginConfig, createTime, updateTime, requestURL string
		err = rows.Scan(&connID, &pluginName, &pluginConfig, &createTime, &updateTime, &pluginPriority, &pluginStatus, &requestURL)
		if err != nil {
			info := err.Error()
			log.Info(info)
		}
		pluginInfo := map[string]interface{}{
			"connID":         connID,
			"pluginName":     pluginName,
			"pluginConfig":   pluginConfig,
			"pluginPriority": pluginPriority,
			"pluginStatus":   pluginStatus,
			"createTime":     createTime,
			"updateTime":     updateTime,
			"requestURL":     requestURL,
		}
		pluginList = append(pluginList, pluginInfo)
	}
	return true, pluginList, nil
}

//GetPluginIndex 获取插件优先级
func (d *APIPluginDao) GetPluginIndex(pluginName string) (bool, int, error) {
	db := d.db
	var pluginPriority int
	sql := "SELECT pluginPriority FROM goku_plugin WHERE pluginName = ?;"
	err := db.QueryRow(sql, pluginName).Scan(pluginPriority)
	if err != nil {
		return false, 0, err
	}
	return true, pluginPriority, nil
}

//GetAPIPluginConfig 通过APIID获取配置信息
func (d *APIPluginDao) GetAPIPluginConfig(apiID int, strategyID, pluginName string) (bool, map[string]string, error) {
	db := d.db
	sql := "SELECT goku_gateway_api.apiName,goku_gateway_api.requestURL,goku_conn_plugin_api.pluginConfig FROM goku_conn_plugin_api INNER JOIN goku_gateway_api ON goku_gateway_api.apiID = goku_conn_plugin_api.apiID WHERE goku_conn_plugin_api.apiID = ? AND goku_conn_plugin_api.strategyID = ? AND goku_conn_plugin_api.pluginName = ?;"
	var p, apiName, requestURL string
	err := db.QueryRow(sql, apiID, strategyID, pluginName).Scan(&apiName, &requestURL, &p)
	if err != nil {
		if err == SQL.ErrNoRows {
			return false, nil, errors.New("[ERROR]Can not find the plugin")
		}
		return false, nil, err
	}
	apiPluginInfo := map[string]string{
		"pluginConfig": p,
		"apiName":      apiName,
		"requestURL":   requestURL,
	}
	return true, apiPluginInfo, nil
}

//CheckPluginIsExistInAPI 检查策略组是否绑定插件
func (d *APIPluginDao) CheckPluginIsExistInAPI(strategyID, pluginName string, apiID int) (bool, error) {
	db := d.db
	sql := "SELECT apiID FROM goku_conn_plugin_api WHERE strategyID = ? AND pluginName = ? AND apiID = ?;"
	var id int
	err := db.QueryRow(sql, strategyID, pluginName, apiID).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, err
}

// GetAPIPluginInStrategyByAPIID 通过接口ID获取策略组中接口插件列表
func (d *APIPluginDao) GetAPIPluginInStrategyByAPIID(strategyID string, apiID int, keyword string, condition int) (bool, []map[string]interface{}, map[string]interface{}, error) {
	db := d.db
	var (
		apiName       string
		requestURL    string
		targetURL     string
		target        string
		rewriteTarget string
	)
	sql := "SELECT A.apiName,A.requestURL,IFNULL(A.targetURL,''),IFNULL(A.balanceName,''),IFNULL(B.target,'') FROM goku_gateway_api A INNER JOIN goku_conn_strategy_api B ON A.apiID = B.apiID WHERE B.apiID = ? AND B.strategyID = ?;"
	err := db.QueryRow(sql, apiID, strategyID).Scan(&apiName, &requestURL, &targetURL, &target, &rewriteTarget)
	if err != nil {
		return false, nil, nil, err
	}
	apiInfo := map[string]interface{}{
		"apiID":         apiID,
		"apiName":       apiName,
		"requestURL":    requestURL,
		"targetURL":     targetURL,
		"target":        target,
		"rewriteTarget": rewriteTarget,
	}

	rule := make([]string, 0, 3)

	rule = append(rule, fmt.Sprintf("A.strategyID = '%s'", strategyID))
	rule = append(rule, fmt.Sprintf("A.apiID = %d", apiID))
	if keyword != "" {
		searchRule := "(A.pluginName LIKE '%" + keyword + "%' OR C.pluginDesc LIKE '%" + keyword + "%')"
		rule = append(rule, searchRule)
	}
	if condition > 0 {
		rule = append(rule, fmt.Sprintf("CASE WHEN C.pluginStatus=0 THEN -1 ELSE A.pluginStatus END = %d", condition-1))
	}
	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}
	sql = fmt.Sprintf(`SELECT A.connID,A.pluginName,IFNULL(A.createTime,""),IFNULL(A.updateTime,""),CASE WHEN C.pluginStatus=0 THEN -1 ELSE A.pluginStatus END as pluginStatus,IFNULL(C.pluginDesc,""),CASE WHEN B.remark is null or B.remark = "" THEN B.loginCall ELSE B.remark END AS updaterName FROM goku_conn_plugin_api A LEFT JOIN goku_admin B ON A.updaterID=B.userID INNER JOIN goku_plugin C ON C.pluginName = A.pluginName %s ORDER BY pluginStatus DESC,A.updateTime DESC;`, ruleStr)
	rows, err := db.Query(sql)
	if err != nil {
		return false, nil, nil, err
	}
	defer rows.Close()
	pluginList := make([]map[string]interface{}, 0)
	//获取记录列
	for rows.Next() {
		var updaterName SQL.NullString
		var pluginStatus, connID int
		var pluginName, pluginDesc, createTime, updateTime string
		err = rows.Scan(&connID, &pluginName, &createTime, &updateTime, &pluginStatus, &pluginDesc, &updaterName)
		if err != nil {
			return false, nil, nil, err
		}

		pluginInfo := map[string]interface{}{
			"connID":       connID,
			"pluginName":   pluginName,
			"pluginStatus": pluginStatus,
			"createTime":   createTime,
			"updateTime":   updateTime,
			"updaterName":  updaterName.String,
			"pluginDesc":   pluginDesc,
		}
		pluginList = append(pluginList, pluginInfo)
	}
	return true, pluginList, apiInfo, nil
}

//GetAllAPIPluginInStrategy 获取策略组中所有接口插件列表
func (d *APIPluginDao) GetAllAPIPluginInStrategy(strategyID string) (bool, []map[string]interface{}, error) {
	db := d.db
	sql := `SELECT goku_conn_plugin_api.connID,goku_conn_plugin_api.apiID,goku_gateway_api.apiName,goku_gateway_api.requestURL,goku_conn_plugin_api.pluginName,IFNULL(goku_conn_plugin_api.createTime,""),IFNULL(goku_conn_plugin_api.updateTime,""),IF(goku_plugin.pluginStatus=0,-1,goku_conn_plugin_api.pluginStatus) as pluginStatus,IFNULL(goku_plugin.pluginDesc,"") FROM goku_conn_plugin_api INNER JOIN goku_gateway_api ON goku_gateway_api.apiID = goku_conn_plugin_api.apiID INNER JOIN goku_plugin ON goku_plugin.pluginName = goku_conn_plugin_api.pluginName WHERE goku_conn_plugin_api.strategyID = ? ORDER BY pluginStatus DESC,goku_conn_plugin_api.updateTime DESC;`
	rows, err := db.Query(sql, strategyID)
	if err != nil {
		return false, make([]map[string]interface{}, 0), err
	}
	defer rows.Close()
	pluginList := make([]map[string]interface{}, 0)
	//获取记录列

	for rows.Next() {
		var pluginStatus, apiID, connID int
		var apiName, pluginName, pluginDesc, createTime, updateTime, requestURL string
		err = rows.Scan(&connID, &apiID, &apiName, &requestURL, &pluginName, &createTime, &updateTime, &pluginStatus, &pluginDesc)
		if err != nil {
			return false, make([]map[string]interface{}, 0), err
		}
		pluginInfo := map[string]interface{}{
			"connID":       connID,
			"apiID":        apiID,
			"apiName":      apiName,
			"pluginName":   pluginName,
			"pluginStatus": pluginStatus,
			"createTime":   createTime,
			"updateTime":   updateTime,
			"requestURL":   requestURL,
			"pluginDesc":   pluginDesc,
		}
		pluginList = append(pluginList, pluginInfo)
	}
	return true, pluginList, nil
}

//BatchEditAPIPluginStatus 批量修改策略组插件状态
func (d *APIPluginDao) BatchEditAPIPluginStatus(connIDList, strategyID string, pluginStatus, userID int) (bool, string, error) {
	db := d.db
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	updateTag := t.Format("20060102150405")
	Tx, _ := db.Begin()
	sql := "UPDATE goku_conn_plugin_api SET updateTag = ?,pluginStatus = ?,updateTime = ?,updaterID = ? WHERE connID IN (" + connIDList + ");"
	_, err := Tx.Exec(sql, updateTag, pluginStatus, now, userID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	// 根据connID获取apiID
	sql = "SELECT apiID FROM goku_conn_plugin_api WHERE connID IN (" + connIDList + ");"
	rows, err := db.Query(sql)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Illegal SQL Statement!", err
	}
	defer rows.Close()
	//获取记录列

	for rows.Next() {
		var apiID int
		err = rows.Scan(&apiID)
		if err != nil {
			Tx.Rollback()
			return false, "[ERROR]Fail to get data!", err
		}
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

//BatchDeleteAPIPlugin 批量删除策略组插件
func (d *APIPluginDao) BatchDeleteAPIPlugin(connIDList, strategyID string) (bool, string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	apiIDList := make([]int, 0)
	// 根据connID获取apiID
	sql := "SELECT apiID FROM goku_conn_plugin_api WHERE connID IN (" + connIDList + ");"
	rows, err := Tx.Query(sql)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Illegal SQL Statement!", err
	}
	defer rows.Close()
	//获取记录列

	for rows.Next() {
		var apiID int
		err = rows.Scan(&apiID)
		if err != nil {
			Tx.Rollback()
			return false, "[ERROR]Fail to get data!", err
		}
		apiIDList = append(apiIDList, apiID)
	}
	sql = "DELETE FROM goku_conn_plugin_api WHERE connID IN (" + connIDList + ");"
	_, err = Tx.Exec(sql)
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

//GetAPIPluginName 通过connID获取插件名称
func (d *APIPluginDao) GetAPIPluginName(connID int) (bool, string, error) {
	db := d.db
	var pluginName string
	sql := "SELECT pluginName FROM goku_conn_plugin_api WHERE connID = ?"
	err := db.QueryRow(sql, connID).Scan(&pluginName)
	if err != nil {
		return false, "[ERROR]The plugin is not existing!", err
	}
	return true, "", nil
}

//CheckAPIPluginIsExistByConnIDList 通过connIDList判断插件是否存在
func (d *APIPluginDao) CheckAPIPluginIsExistByConnIDList(connIDList, pluginName string) (bool, []int, error) {
	db := d.db
	sql := "SELECT apiID FROM goku_conn_plugin_api WHERE connID IN (" + connIDList + ") AND pluginName = ?;"
	rows, err := db.Query(sql, pluginName)
	if err != nil {
		return false, make([]int, 0), err
	}
	defer rows.Close()
	apiIDList := make([]int, 0)

	for rows.Next() {
		var apiID int
		err = rows.Scan(&apiID)
		if err != nil {
			return false, make([]int, 0), err
		}
		apiIDList = append(apiIDList, apiID)
	}
	return true, apiIDList, nil
}

//GetAPIPluginListWithNotAssignAPIList 获取没有绑定嵌套插件列表
func (d *APIPluginDao) GetAPIPluginListWithNotAssignAPIList(strategyID string) (bool, []map[string]interface{}, error) {
	db := d.db
	sql := "SELECT pluginID,pluginDesc,pluginName FROM goku_plugin WHERE pluginType = 2 AND pluginStatus = 1;"
	rows, err := db.Query(sql)
	if err != nil {
		return false, make([]map[string]interface{}, 0), err
	}
	defer rows.Close()
	pluginList := make([]map[string]interface{}, 0)
	//获取记录列

	sql = "SELECT goku_gateway_api.apiID,goku_gateway_api.apiName,goku_gateway_api.requestURL FROM goku_gateway_api INNER JOIN goku_conn_strategy_api ON goku_gateway_api.apiID = goku_conn_strategy_api.apiID WHERE goku_conn_strategy_api.strategyID = ? AND goku_gateway_api.apiID NOT IN (SELECT goku_conn_plugin_api.apiID FROM goku_conn_plugin_api WHERE goku_conn_plugin_api.strategyID = ? AND goku_conn_plugin_api.pluginName = ?);"
	for rows.Next() {
		var pluginID int
		var pluginName, chineseName string
		err = rows.Scan(&pluginID, &chineseName, &pluginName)
		if err != nil {
			info := err.Error()
			log.Info(info)
			return false, make([]map[string]interface{}, 0), err
		}
		r, err := db.Query(sql, strategyID, strategyID, pluginName)
		if err != nil {
			return false, make([]map[string]interface{}, 0), err
		}
		defer r.Close()
		apiList := make([]map[string]interface{}, 0)
		for r.Next() {
			var (
				apiID      int
				apiName    string
				requestURL string
			)
			err = r.Scan(&apiID, &apiName, &requestURL)
			if err != nil {
				return false, make([]map[string]interface{}, 0), err
			}
			apiList = append(apiList, map[string]interface{}{
				"apiID":      apiID,
				"apiName":    apiName,
				"requestURL": requestURL,
			})

		}
		pluginInfo := map[string]interface{}{
			"chineseName": chineseName,
			"pluginName":  pluginName,
			"pluginID":    pluginID,
			"apiList":     apiList,
		}
		pluginList = append(pluginList, pluginInfo)
	}
	return true, pluginList, nil
}

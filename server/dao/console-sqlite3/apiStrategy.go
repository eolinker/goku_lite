package console_sqlite3

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/eolinker/goku-api-gateway/server/dao"
)

//APIStrategyDao APIStrategyDao
type APIStrategyDao struct {
	db *sql.DB
}

//NewAPIStrategyDao new APIStrategyDao
func NewAPIStrategyDao() *APIStrategyDao {
	return &APIStrategyDao{}
}

//Create create
func (d *APIStrategyDao) Create(db *sql.DB) (interface{}, error) {
	d.db = db
	var i dao.APIStrategyDao = d
	return &i, nil
}

//AddAPIToStrategy 将接口加入策略组
func (d *APIStrategyDao) AddAPIToStrategy(apiList []string, strategyID string) (bool, string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	sql2 := "SELECT apiID FROM goku_conn_strategy_api WHERE apiID = ? AND strategyID = ?"
	sql1 := "SELECT apiID FROM goku_gateway_api WHERE apiID = ?"
	sql3 := "INSERT INTO goku_conn_strategy_api (apiID,strategyID,updateTime) VALUES (?,?,?)"
	Tx, _ := db.Begin()
	stmt1, _ := Tx.Prepare(sql1)
	stmt2, _ := Tx.Prepare(sql2)
	stmt3, _ := Tx.Prepare(sql3)
	defer stmt1.Close()
	defer stmt2.Close()
	defer stmt3.Close()

	for _, apiID := range apiList {
		id, err := strconv.Atoi(apiID)
		if err != nil {
			continue
		}
		// 查询ID是否存在,若不存在，则跳过

		var aID int
		err = stmt1.QueryRow(apiID).Scan(&aID)
		if err != nil {
			continue
		}
		// 查询此接口是否被加入策略组

		err = stmt2.QueryRow(apiID, strategyID).Scan(&aID)
		if err == nil {
			continue
		}
		_, err = stmt3.Exec(id, strategyID, now)
		if err != nil {
			Tx.Rollback()
			return false, "[ERROR]Failed to insert data!", err
		}
	}
	// 更新策略修改时间
	sql4 := "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?"
	_, err := Tx.Exec(sql4, now, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Failed to update data!", err
	}
	Tx.Commit()
	return true, "", nil
}

// SetAPITargetOfStrategy 重定向接口负载
func (d *APIStrategyDao) SetAPITargetOfStrategy(apiID int, strategyID string, target string) (bool, string, error) {
	db := d.db
	sql := "UPDATE goku_conn_strategy_api SET `target` = ? where apiID = ? AND strategyID = ? "
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, err.Error(), err
	}
	defer stmt.Close()
	_, e := stmt.Exec(target, apiID, strategyID)

	if e != nil {
		return false, e.Error(), e
	}

	return true, "", nil
}

// BatchSetAPITargetOfStrategy 批量重定向接口负载
func (d *APIStrategyDao) BatchSetAPITargetOfStrategy(apiIds []int, strategyID string, target string) (bool, string, error) {
	idLen := len(apiIds)
	s := make([]interface{}, 0, idLen+2)
	c := ""
	s = append(s, target, strategyID)
	for i, id := range apiIds {
		c += "?"
		if i < idLen-1 {
			c += ","
		}
		s = append(s, id)
	}
	db := d.db
	sql := fmt.Sprintf("UPDATE goku_conn_strategy_api SET `target` = ? where strategyID = ? AND apiID IN (%s) ", c)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, err.Error(), err
	}
	defer stmt.Close()
	_, e := stmt.Exec(s...)

	if e != nil {
		return false, e.Error(), e
	}

	return true, "", nil
}

func (d *APIStrategyDao) getAPIOfStrategyRule(condition int, balanceNames []string, ids []int) []string {
	rule := make([]string, 0, 2)
	switch condition {
	case 1, 2:
		{
			balenceNameLen := len(balanceNames)
			nameType := "A.balanceName"
			if condition == 2 {
				nameType = "S.`target`"
			}
			nameStr := ""
			for i := 0; i < balenceNameLen; i++ {
				nameStr += fmt.Sprintf("'%s'", balanceNames[i])
				if i < balenceNameLen-1 {
					nameStr += ","
				}
			}
			rule = append(rule, fmt.Sprintf("%s IN (%s)", nameType, nameStr))
		}
	case 3, 4:
		{
			idsStr := ""
			idLen := len(ids)
			if len(ids) < 1 {
				break
			}
			for i, id := range ids {
				idsStr += strconv.Itoa(id)
				if i < idLen-1 {
					idsStr += ","
				}
			}
			if condition == 3 {
				rule = append(rule, fmt.Sprintf("A.managerID IN (%s)", idsStr))
			} else if condition == 4 {
				rule = append(rule, fmt.Sprintf("A.lastUpdateUserID IN (%s)", idsStr))
			}
		}
	}
	return rule
}

// GetAPIIDListFromStrategy 获取策略组接口列表
func (d *APIStrategyDao) GetAPIIDListFromStrategy(strategyID, keyword string, condition int, ids []int, balanceNames []string) (bool, []int, error) {
	rule := make([]string, 0, 10)

	rule = append(rule, fmt.Sprintf("S.strategyID = '%s'", strategyID))
	if keyword != "" {
		searchRule := "(A.apiName LIKE '%" + keyword + "%' OR A.requestURL LIKE '%" + keyword + "%' "
		searchRule += " OR IFNULL(A.balanceName,'') LIKE '%" + keyword + "%' OR A.targetURL LIKE '%" + keyword + "%' OR S.`target` LIKE '%" + keyword + "%')"
		rule = append(rule, searchRule)
	}
	if condition > 0 {
		rule = append(rule, d.getAPIOfStrategyRule(condition, balanceNames, ids)...)
	}
	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}
	sql := fmt.Sprintf("SELECT A.`apiID` FROM `goku_gateway_api` A INNER JOIN `goku_conn_strategy_api` S ON S.`apiID` = A.`apiID` %s", ruleStr)
	rows, err := d.db.Query(sql)
	if err != nil {
		return false, make([]int, 0), err
	}
	defer rows.Close()

	//获取记录列
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

// GetAPIListFromStrategy 获取策略组接口列表
func (d *APIStrategyDao) GetAPIListFromStrategy(strategyID, keyword string, condition, page, pageSize int, ids []int, balanceNames []string) (bool, []map[string]interface{}, int, error) {
	rule := make([]string, 0, 2)

	rule = append(rule, fmt.Sprintf("S.strategyID = '%s'", strategyID))
	if keyword != "" {
		searchRule := "(A.apiName LIKE '%" + keyword + "%' OR A.requestURL LIKE '%" + keyword + "%' "
		searchRule += " OR IFNULL(A.balanceName,'') LIKE '%" + keyword + "%' OR A.targetURL LIKE '%" + keyword + "%' OR S.`target` LIKE '%" + keyword + "%')"
		rule = append(rule, searchRule)
	}
	if condition > 0 {
		rule = append(rule, d.getAPIOfStrategyRule(condition, balanceNames, ids)...)
	}
	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}

	sql := fmt.Sprintf("SELECT A.`apiID`, A.`apiName`, A.`requestURL`,A.`requestMethod`,CASE WHEN A.`apiType`=0 THEN A.`targetURL` ELSE '' END,A.apiType,IFNULL(A.`targetMethod`,''), A.`isFollow`, IFNULL(A.`updateTime`,'') AS updateTime, A.`lastUpdateUserID`, A.`managerID`, IFNULL(A.`balanceName`,'') As `target`, IFNULL(S.`target`,'') as `rewriteTarget`,  CASE WHEN AD.`remark` is null or AD.`remark` = '' THEN AD.`loginCall` ELSE AD.`remark` END AS managerName, CASE WHEN AD2.`remark` is null or AD2.`remark` = '' THEN AD2.`loginCall` ELSE AD2.`remark` END AS updaterName  FROM `goku_gateway_api` A INNER JOIN `goku_conn_strategy_api` S ON S.`apiID` = A.`apiID` LEFT JOIN `goku_admin` AD ON A.`managerID` = AD.`userID` LEFT JOIN `goku_admin` AD2 ON A.`lastUpdateUserID` = AD2.`userID` %s", ruleStr)
	count := getCountSQL(d.db, sql)
	rows, err := getPageSQL(d.db, sql, "S.`connID`", "DESC", page, pageSize)
	if err != nil {
		return false, make([]map[string]interface{}, 0), 0, err
	}
	defer rows.Close()

	//获取记录列
	apiList := make([]map[string]interface{}, 0)
	for rows.Next() {
		var apiID, updaterID, managerID, apiType int
		var apiName, requestURL, updateTime, updaterName, managerName, target, targetURL, rewriteTarget, requestMethod, targetMethod string
		var isFollow bool
		err = rows.Scan(&apiID, &apiName, &requestURL, &requestMethod, &targetURL, &apiType, &targetMethod, &isFollow, &updateTime, &updaterID, &managerID, &target, &rewriteTarget, &managerName, &updaterName)
		if err != nil {
			return false, make([]map[string]interface{}, 0), 0, err
		}
		apiInfo := map[string]interface{}{
			"apiID":         apiID,
			"apiName":       apiName,
			"requestURL":    requestURL,
			"updateTime":    updateTime,
			"updaterName":   updaterName,
			"managerName":   managerName,
			"target":        target,
			"targetURL":     targetURL,
			"rewriteTarget": rewriteTarget,
			"requestMethod": strings.ToUpper(requestMethod),
			"targetMethod":  strings.ToUpper(targetMethod),
			"isFollow":      isFollow,
			"apiType":       apiType,
		}
		apiList = append(apiList, apiInfo)
	}
	return true, apiList, count, nil
}

// CheckIsExistAPIInStrategy 检查插件是否添加进策略组
func (d *APIStrategyDao) CheckIsExistAPIInStrategy(apiID int, strategyID string) (bool, string, error) {
	db := d.db
	var id int
	sql := "SELECT connID FROM goku_conn_strategy_api WHERE apiID = ? AND strategyID = ?"
	err := db.QueryRow(sql, apiID, strategyID).Scan(&id)
	if err != nil {
		return false, "", err
	}
	return true, "", nil
}

// 获取策略绑定的简易接口列表
func (d *APIStrategyDao) getSimpleAPIListInStrategy(strategyID string, projectID int) map[string]string {
	db := d.db
	sql := "SELECT goku_gateway_api.requestURL,GROUP_CONCAT(DISTINCT goku_gateway_api.requestMethod) AS requestMethod FROM goku_gateway_api INNER JOIN goku_conn_strategy_api ON goku_gateway_api.apiID = goku_conn_strategy_api.apiID where goku_conn_strategy_api.strategyID = ? AND goku_gateway_api.projectID = ? GROUP BY requestURL"
	rows, err := db.Query(sql, strategyID, projectID)
	if err != nil {
		return nil
	}
	defer rows.Close()
	simpleMap := make(map[string]string)
	for rows.Next() {
		var requestURL, requestMethod string
		err = rows.Scan(&requestURL, &requestMethod)
		if err != nil {
			return nil
		}
		simpleMap[requestURL] = requestMethod
	}
	return simpleMap
}

// GetAPIIDListNotInStrategy 获取未被该策略组绑定的接口ID列表(通过项目)
func (d *APIStrategyDao) GetAPIIDListNotInStrategy(strategyID string, projectID, groupID int, keyword string) (bool, []int, error) {
	requestMap := d.getSimpleAPIListInStrategy(strategyID, projectID)
	rule := make([]string, 0, 3)

	rule = append(rule, fmt.Sprintf("A.projectID = %d", projectID))
	rule = append(rule, fmt.Sprintf("A.apiID NOT IN (SELECT apiID FROM goku_conn_strategy_api WHERE strategyID = '%s')", strategyID))
	if keyword != "" {
		searchRule := "(A.apiName LIKE '%" + keyword + "%' OR A.requestURL LIKE '%" + keyword + "%'"
		searchRule += " OR IFNULL(A.balanceName,'') LIKE '%" + keyword + "%' OR A.targetURL LIKE '%" + keyword + "%')"
		rule = append(rule, searchRule)
	}
	groupRule, err := d.getAPIGroupRule(projectID, groupID)
	if err != nil {
		return false, make([]int, 0), err
	}
	if groupRule != "" {
		rule = append(rule, groupRule)
	}
	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}

	sql := fmt.Sprintf("SELECT A.apiID,A.requestURL,A.requestMethod FROM goku_gateway_api A %s", ruleStr)
	rows, err := d.db.Query(sql)
	if err != nil {
		return false, make([]int, 0), err
	}
	defer rows.Close()
	apiIDList := make([]int, 0)
	//获取记录列
	for rows.Next() {
		var apiID int
		var requestURL, requestMethod string
		err = rows.Scan(&apiID, &requestURL, &requestMethod)
		if err != nil {
			return false, make([]int, 0), err
		}
		if value, ok := requestMap[requestURL]; ok {
			if strings.Contains(strings.ToUpper(value), strings.ToUpper(requestMethod)) {
				continue
			}
		}
		apiIDList = append(apiIDList, apiID)
	}
	return true, apiIDList, nil
}

func (d *APIStrategyDao) getAPIGroupRule(projectID, groupID int) (string, error) {
	db := d.db
	if groupID < 1 {
		if groupID == 0 {
			groupRule := fmt.Sprintf("A.groupID = %d", groupID)
			return groupRule, nil
		}
		return "", nil
	}
	var groupPath string
	sql := "SELECT groupPath FROM goku_gateway_api_group WHERE groupID = ?;"
	err := db.QueryRow(sql, groupID).Scan(&groupPath)
	if err != nil {
		return "", err
	}
	// 获取分组ID列表
	sql = "SELECT GROUP_CONCAT(DISTINCT groupID) AS groupID FROM goku_gateway_api_group WHERE projectID = ? AND groupPath LIKE ?;"
	groupIDList := ""
	err = db.QueryRow(sql, projectID, groupPath+"%").Scan(&groupIDList)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("A.groupID IN (%s)", groupIDList), nil
}

// GetAPIListNotInStrategy 获取未被该策略组绑定的接口列表(通过项目)
func (d *APIStrategyDao) GetAPIListNotInStrategy(strategyID string, projectID, groupID, page, pageSize int, keyword string) (bool, []map[string]interface{}, int, error) {
	requestMap := d.getSimpleAPIListInStrategy(strategyID, projectID)
	rule := make([]string, 0, 3)

	rule = append(rule, fmt.Sprintf("A.projectID = %d", projectID))
	rule = append(rule, fmt.Sprintf("A.apiID NOT IN (SELECT apiID FROM goku_conn_strategy_api WHERE strategyID = '%s')", strategyID))
	if keyword != "" {
		searchRule := "(A.apiName LIKE '%" + keyword + "%' OR A.requestURL LIKE '%" + keyword + "%'"
		searchRule += " OR IFNULL(A.balanceName,'') LIKE '%" + keyword + "%' OR A.targetURL LIKE '%" + keyword + "%')"
		rule = append(rule, searchRule)
	}

	groupRule, err := d.getAPIGroupRule(projectID, groupID)
	if err != nil {
		return false, make([]map[string]interface{}, 0), 0, err
	}
	if groupRule != "" {
		rule = append(rule, groupRule)
	}

	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}
	sql := fmt.Sprintf("SELECT A.apiID,A.apiName,A.requestURL,A.requestMethod,IFNULL(A.balanceName,''),CASE WHEN A.apiType=0 THEN A.targetURL ELSE '' END,A.apiType,IFNULL(A.`targetMethod`,''), A.`isFollow`,A.groupID,IFNULL(G.groupPath,A.groupID) FROM goku_gateway_api A LEFT JOIN goku_gateway_api_group G ON G.groupID = A.groupID  %s", ruleStr)
	count := getCountSQL(d.db, sql)
	rows, err := getPageSQL(d.db, sql, "A.`updateTime`", "DESC", page, pageSize)
	if err != nil {
		return false, make([]map[string]interface{}, 0), 0, err
	}
	defer rows.Close()
	apiList := make([]map[string]interface{}, 0)
	//获取记录列
	for rows.Next() {
		var apiID, groupID, apiType int
		var apiName, requestURL, requestMethod, targetServer, groupPath, targetURL, targetMethod string
		var isFollow bool
		err = rows.Scan(&apiID, &apiName, &requestURL, &requestMethod, &targetServer, &targetURL, &apiType, &targetMethod, &isFollow, &groupID, &groupPath)
		if err != nil {
			return false, make([]map[string]interface{}, 0), 0, err
		}
		if value, ok := requestMap[requestURL]; ok {
			if strings.Contains(strings.ToUpper(value), strings.ToUpper(requestMethod)) {
				count = count - 1
				continue
			}
		}
		apiInfo := map[string]interface{}{
			"apiID":         apiID,
			"apiName":       apiName,
			"requestURL":    requestURL,
			"requestMethod": strings.ToUpper(requestMethod),
			"target":        targetServer,
			"targetURL":     targetURL,
			"groupID":       groupID,
			"groupPath":     groupPath,
			"targetMethod":  strings.ToUpper(targetMethod),
			"isFollow":      isFollow,
			"apiType":       apiType,
		}
		apiList = append(apiList, apiInfo)
	}
	return true, apiList, count, nil
}

//BatchDeleteAPIInStrategy 批量删除策略组接口
func (d *APIStrategyDao) BatchDeleteAPIInStrategy(apiIDList, strategyID string) (bool, string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	sql := "DELETE FROM goku_conn_strategy_api WHERE strategyID = ? AND apiID IN (" + apiIDList + ")"
	_, err := Tx.Exec(sql, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	sql = "DELETE FROM goku_conn_plugin_api WHERE strategyID = ? AND apiID IN (" + apiIDList + ")"
	_, err = Tx.Exec(sql, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	sql = "UPDATE goku_gateway_strategy SET updateTime = ? WHERE strategyID = ?"
	_, err = Tx.Exec(sql, now, strategyID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Failed to update data!", err
	}
	Tx.Commit()
	return true, "", nil
}

package console_sqlite3

import (
	SQL "database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/eolinker/goku-api-gateway/server/dao"

	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//APIDao APIDao
type APIDao struct {
	db *SQL.DB
}

//NewAPIDao new APIDao
func NewAPIDao() *APIDao {
	return &APIDao{}
}

//Create create
func (d *APIDao) Create(db *SQL.DB) (interface{}, error) {
	d.db = db
	var i dao.APIDao = d
	return &i, nil
}

// AddAPI 新增接口
func (d *APIDao) AddAPI(apiName, alias, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkAPIs, staticResponse, responseDataType, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, managerID, userID, apiType int) (bool, int, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	res, err := Tx.Exec("INSERT INTO goku_gateway_api (projectID,groupID,apiName,requestURL,targetURL,requestMethod,targetMethod,protocol,linkAPIs,staticResponse,responseDataType,balanceName,isFollow,timeout,retryCount,alertValve,createTime,updateTime,managerID,lastUpdateUserID,createUserID,apiType) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);", projectID, groupID, apiName, requestURL, targetURL, requestMethod, targetMethod, protocol, linkAPIs, staticResponse, responseDataType, balanceName, isFollow, timeout, retryCount, alertValve, now, now, managerID, userID, userID, apiType)

	if err != nil {
		Tx.Rollback()
		return false, 0, err
	}
	apiID, _ := res.LastInsertId()
	// 更新项目更新时间
	_, err = Tx.Exec("UPDATE goku_gateway_project SET updateTime = ? WHERE projectID = ?;", now, projectID)
	if err != nil {
		Tx.Rollback()
		return false, 0, err
	}
	Tx.Commit()
	return true, int(apiID), nil
}

// EditAPI 修改接口
func (d *APIDao) EditAPI(apiName, alias, requestURL, targetURL, requestMethod, targetMethod, isFollow, linkAPIs, staticResponse, responseDataType, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, apiID, managerID, userID int) (bool, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	_, err := Tx.Exec("UPDATE goku_gateway_api SET projectID = ?,groupID = ?,apiName = ?,requestURL = ?,targetURL = ?,requestMethod = ?,protocol = ?,balanceName = ?,targetMethod = ?,isFollow = ?,linkAPIs = ?,staticResponse = ?,responseDataType = ?,timeout = ?,retryCount = ?,alertValve = ?,updateTime = ?,managerID = ?,lastUpdateUserID = ? WHERE apiID = ?", projectID, groupID, apiName, requestURL, targetURL, requestMethod, protocol, balanceName, targetMethod, isFollow, linkAPIs, staticResponse, responseDataType, timeout, retryCount, alertValve, now, managerID, userID, apiID)

	if err != nil {
		Tx.Rollback()
		return false, err
	}
	// 更新项目更新时间
	_, err = Tx.Exec("UPDATE goku_gateway_project SET updateTime = ? WHERE projectID = ?;", now, projectID)
	if err != nil {
		Tx.Rollback()
		return false, err
	}
	Tx.Commit()
	return true, nil
}

// GetAPIInfo 获取接口信息
func (d *APIDao) GetAPIInfo(apiID int) (bool, *entity.API, error) {
	db := d.db
	sql := `SELECT A.apiID,A.groupID,A.apiName,A.requestURL,A.targetURL,A.requestMethod,A.targetMethod,IFNULL(A.protocol,"http"),IFNULL(A.balanceName,""),A.isFollow,A.timeout,A.retryCount,A.alertValve,A.createTime,A.updateTime,A.managerID,A.lastUpdateUserID,A.createUserID,IFNULL(goku_gateway_api_group.groupPath,"0"),A.apiType,IFNULL(A.linkAPIs,''),IFNULL(A.staticResponse,''),IFNULL(A.responseDataType,'origin') FROM goku_gateway_api A LEFT JOIN goku_gateway_api_group ON A.groupID = goku_gateway_api_group.groupID WHERE A.apiID = ?`
	api := &entity.API{}
	var managerInfo entity.ManagerInfo
	var linkAPIs string
	err := db.QueryRow(sql, apiID).Scan(&api.APIID, &api.GroupID, &api.APIName, &api.RequestURL, &api.ProxyURL, &api.RequestMethod, &api.TargetMethod, &api.Protocol, &api.BalanceName, &api.IsFollow, &api.Timeout, &api.RetryConut, &api.Valve, &api.CreateTime, &api.UpdateTime, &managerInfo.ManagerID, &managerInfo.UpdaterID, &managerInfo.CreateUserID, &api.GroupPath, &api.APIType, &linkAPIs, &api.StaticResponse, &api.ResponseDataType)
	if err != nil {
		return false, &entity.API{}, err
	}
	json.Unmarshal([]byte(linkAPIs), &api.LinkAPIs)
	api.RequestMethod = strings.ToUpper(api.RequestMethod)

	sql = `SELECT IFNULL(remark,loginCall) as userName FROM goku_admin WHERE userID = ?;`
	err = db.QueryRow(sql, managerInfo.ManagerID).Scan(&managerInfo.ManagerName)
	if err != nil {
		if err != SQL.ErrNoRows {
			return false, &entity.API{}, err
		}
	}
	err = db.QueryRow(sql, managerInfo.UpdaterID).Scan(&managerInfo.UpdaterName)
	if err != nil {
		if err != SQL.ErrNoRows {
			return false, &entity.API{}, err
		}
	}
	err = db.QueryRow(sql, managerInfo.CreateUserID).Scan(&managerInfo.CreateUserName)
	if err != nil {
		if err != SQL.ErrNoRows {
			return false, nil, err
		}
	}
	api.ManagerInfo = &managerInfo
	return true, api, nil
}

//GetAPIListByGroupList 通过分组列表获取接口列表
func (d *APIDao) GetAPIListByGroupList(projectID int, groupIDList string) (bool, []map[string]interface{}, error) {
	db := d.db
	// 获取分组ID列表
	sql := `SELECT A.apiID,A.apiName,A.requestURL,IFNULL(A.updateTime,""),A.lastUpdateUserID,A.managerID FROM goku_gateway_api A WHERE A.projectID = ? AND A.groupID IN (` + groupIDList + `) ORDER BY A.updateTime DESC;`

	rows, err := db.Query(sql, projectID)
	if err != nil {
		return false, make([]map[string]interface{}, 0), err
	}
	defer rows.Close()
	apiList := make([]map[string]interface{}, 0)
	//获取记录列

	for rows.Next() {
		var apiID, updaterID, managerID int
		var apiName, requestURL, updateTime, managerName, updaterName string
		err = rows.Scan(&apiID, &apiName, &requestURL, &updateTime, &updaterID, &managerID)
		if err != nil {
			return false, make([]map[string]interface{}, 0), err
		}
		sql = `SELECT IFNULL(remark,loginCall) as userName FROM goku_admin WHERE userID = ?;`
		err = db.QueryRow(sql, managerID).Scan(&managerName)
		if err != nil {
			if err != SQL.ErrNoRows {
				return false, make([]map[string]interface{}, 0), err
			}
		}
		err = db.QueryRow(sql, updaterID).Scan(&updaterName)
		if err != nil {
			if err != SQL.ErrNoRows {
				return false, make([]map[string]interface{}, 0), err
			}
		}
		apiInfo := map[string]interface{}{
			"apiID":       apiID,
			"apiName":     apiName,
			"requestURL":  requestURL,
			"updateTime":  updateTime,
			"updaterName": updaterName,
			"managerName": managerName,
		}
		apiList = append(apiList, apiInfo)
	}
	return true, apiList, nil
}

// getAPIRule
func getAPIRule(projectID int, keyword string, condition int, ids []int) []string {
	rule := make([]string, 0, 5)
	rule = append(rule, fmt.Sprintf("A.projectID = %d", projectID))
	if keyword != "" {
		searchRule := "(A.apiName LIKE '%" + keyword + "%' OR A.requestURL LIKE '%" + keyword + "%'"
		searchRule += " OR IFNULL(A.balanceName,'') LIKE '%" + keyword + "%' OR A.targetURL LIKE '%" + keyword + "%')"
		rule = append(rule, searchRule)
	}
	switch condition {
	case 0:
		{
			break
		}
	case 1, 2:
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
			if condition == 1 {
				rule = append(rule, fmt.Sprintf("A.managerID IN (%s)", idsStr))
			} else if condition == 2 {
				rule = append(rule, fmt.Sprintf("A.lastUpdateUserID IN (%s)", idsStr))
			}

		}
	default:
		{
			break
		}
	}
	return rule
}

// GetAPIIDList 获取接口ID列表
func (d *APIDao) GetAPIIDList(projectID int, groupID int, keyword string, condition int, ids []int) (bool, []int, error) {
	db := d.db
	rule := getAPIRule(projectID, keyword, condition, ids)

	if groupID < 1 {
		if groupID == 0 {
			groupRule := fmt.Sprintf("A.groupID = %d", groupID)
			rule = append(rule, groupRule)
		}
	} else {
		var groupPath string
		sql := "SELECT groupPath FROM goku_gateway_api_group WHERE groupID = ?;"
		err := db.QueryRow(sql, groupID).Scan(&groupPath)
		if err != nil {
			return false, make([]int, 0), err
		}
		// 获取分组ID列表
		sql = "SELECT GROUP_CONCAT(DISTINCT groupID) AS groupID FROM goku_gateway_api_group WHERE projectID = ? AND groupPath LIKE ?;"
		groupIDList := ""
		err = db.QueryRow(sql, projectID, groupPath+"%").Scan(&groupIDList)
		if err != nil {
			return false, make([]int, 0), err
		}
		rule = append(rule, fmt.Sprintf("A.groupID IN (%s)", groupIDList))
	}

	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}

	sql := fmt.Sprintf(`SELECT A.apiID FROM goku_gateway_api A  %s`, ruleStr)
	rows, err := db.Query(sql)
	if err != nil {
		return false, make([]int, 0), err
	}
	defer rows.Close()
	apiIDList := make([]int, 0)
	//获取记录列
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

// GetAPIList 获取所有接口列表
func (d *APIDao) GetAPIList(projectID int, groupID int, keyword string, condition, page, pageSize int, ids []int) (bool, []map[string]interface{}, int, error) {
	db := d.db
	rule := getAPIRule(projectID, keyword, condition, ids)

	if groupID < 1 {
		if groupID == 0 {
			groupRule := fmt.Sprintf("A.groupID = %d", groupID)
			rule = append(rule, groupRule)
		}
	} else {
		var groupPath string
		sql := "SELECT groupPath FROM goku_gateway_api_group WHERE groupID = ?;"
		err := db.QueryRow(sql, groupID).Scan(&groupPath)
		if err != nil {
			return false, make([]map[string]interface{}, 0), 0, err
		}
		// 获取分组ID列表
		sql = "SELECT GROUP_CONCAT(DISTINCT groupID) AS groupID FROM goku_gateway_api_group WHERE projectID = ? AND groupPath LIKE ?;"
		groupIDList := ""
		err = db.QueryRow(sql, projectID, groupPath+"%").Scan(&groupIDList)
		if err != nil {
			return false, make([]map[string]interface{}, 0), 0, err
		}
		rule = append(rule, fmt.Sprintf("A.groupID IN (%s)", groupIDList))
	}

	ruleStr := ""
	if len(rule) > 0 {
		ruleStr += "WHERE " + strings.Join(rule, " AND ")
	}

	sql := fmt.Sprintf(`SELECT A.apiID,A.apiName,A.requestURL,A.requestMethod,CASE WHEN A.apiType=0 THEN A.targetURL ELSE '' END,A.apiType,IFNULL(A.balanceName,''),IFNULL(A.updateTime,""),CASE WHEN B.remark is null or B.remark = "" THEN B.loginCall ELSE B.remark END AS updaterName,CASE WHEN C.remark is null or C.remark = "" THEN C.loginCall ELSE C.remark END AS managerName,A.lastUpdateUserID,A.managerID,A.isFollow,IFNULL(A.protocol,"http"),A.targetMethod,A.groupID,IFNULL(D.groupPath,"0"),IFNULL(D.groupName,"未分组") FROM goku_gateway_api A  INNER JOIN goku_admin B ON A.lastUpdateUserID = B.userID INNER JOIN goku_admin C ON A.managerID=C.userID LEFT JOIN goku_gateway_api_group D ON D.groupID = A.groupID %s`, ruleStr)
	count := getCountSQL(d.db, sql)
	rows, err := getPageSQL(d.db, sql, "A.updateTime", "DESC", page, pageSize)
	if err != nil {
		return false, make([]map[string]interface{}, 0), 0, err
	}
	defer rows.Close()
	apiList := make([]map[string]interface{}, 0)
	//获取记录列
	for rows.Next() {
		var apiID, updaterID, managerID, groupID, apiType int
		var apiName, requestURL, updateTime, managerName, updaterName, requestMethod, targetURL, balanceName, targetMethod, protocol, groupPath, groupName string
		var isFollow bool

		err = rows.Scan(&apiID, &apiName, &requestURL, &requestMethod, &targetURL, &apiType, &balanceName, &updateTime, &updaterName, &managerName, &updaterID, &managerID, &isFollow, &protocol, &targetMethod, &groupID, &groupPath, &groupName)
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
			"requestMethod": strings.ToUpper(requestMethod),
			"targetURL":     targetURL,
			"target":        balanceName,
			"protocol":      protocol,
			"targetMethod":  strings.ToUpper(targetMethod),
			"isFollow":      isFollow,
			"groupID":       groupID,
			"groupPath":     groupPath,
			"groupName":     groupName,
			"apiType":       apiType,
		}
		apiList = append(apiList, apiInfo)
	}
	return true, apiList, count, nil
}

//CheckURLIsExist 接口路径是否存在
func (d *APIDao) CheckURLIsExist(requestURL, requestMethod string, projectID, apiID int) bool {
	db := d.db
	var id int
	var m string
	sql := "SELECT apiID,requestMethod FROM goku_gateway_api A WHERE requestURL = ? AND projectID = ?;"
	err := db.QueryRow(sql, requestURL, projectID).Scan(&id, &m)
	if err != nil {
		return false
	}
	method := strings.Split(requestMethod, ",")
	mList := strings.Split(m, ",")
	if apiID == id {
		return false
	}
	for _, v := range mList {
		for _, n := range method {
			if strings.ToUpper(n) == strings.ToUpper(v) {
				return true
			}
		}
	}
	return false
}

//CheckAPIIsExist 检查接口是否存在
func (d *APIDao) CheckAPIIsExist(apiID int) (bool, error) {
	db := d.db
	sql := "SELECT apiID FROM goku_gateway_api A WHERE apiID = ?;"
	var id int
	err := db.QueryRow(sql, apiID).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, err
}

//CheckAliasIsExist 检查别名是否存在
func (d *APIDao) CheckAliasIsExist(apiID int, alias string) bool {
	if alias == "" {
		return false
	}

	db := d.db
	sql := "SELECT apiID FROM goku_gateway_api A WHERE alias = ?;"
	var id int
	err := db.QueryRow(sql, apiID).Scan(&id)
	if err != nil {
		return false
	}
	if id != 0 && apiID == id {
		return false
	}

	return true
}

//BatchEditAPIBalance 批量修改接口负载
func (d *APIDao) BatchEditAPIBalance(apiIDList []string, balance string) (string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")

	sqlTpl := "UPDATE `goku_gateway_api` A LEFT JOIN `goku_gateway_project` P ON A.`projectID` = P.`projectID` SET A.`updateTime` = ?,  P.`updateTime`=?, A.`balanceName`=? WHERE A.`apiID` IN (%s)"

	sql := fmt.Sprint(sqlTpl, strings.Join(apiIDList, ","))

	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Fail to Prepare SQL!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(now, now, balance)
	if err != nil {
		return "[ERROR]Fail to excute SQL statement!", err
	}
	return "", nil
}

//BatchEditAPIGroup 批量修改接口分组
func (d *APIDao) BatchEditAPIGroup(apiIDList []string, groupID int) (string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")

	sqlTpl := "UPDATE `goku_gateway_api` A LEFT JOIN `goku_gateway_project` P ON A.`projectID` = P.`projectID` SET A.`updateTime` = ?,  P.`updateTime`=?, A.`groupID`=? WHERE A.`apiID` IN (%s)"

	sql := fmt.Sprintf(sqlTpl, strings.Join(apiIDList, ","))
	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Fail to Prepare SQL!", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(now, now, groupID)
	if err != nil {
		return "[ERROR]Fail to excute SQL statement!", err
	}
	return "", nil

}

//BatchDeleteAPI 批量修改接口
func (d *APIDao) BatchDeleteAPI(apiIDList string) (bool, string, error) {
	db := d.db
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	sql := "DELETE FROM goku_gateway_api WHERE apiID IN (" + apiIDList + ");"
	_, err := Tx.Exec(sql)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	sql = "DELETE FROM goku_conn_strategy_api WHERE apiID IN (" + apiIDList + ");"
	_, err = Tx.Exec(sql)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	sql = "DELETE FROM goku_conn_plugin_api WHERE apiID IN (" + apiIDList + ");"
	_, err = Tx.Exec(sql)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to delete data!", err
	}

	// 查询接口的projectID
	rows, err := db.Query("SELECT projectID FROM goku_gateway_api A WHERE apiID IN (" + apiIDList + ");")
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	projectIDList := ""
	defer rows.Close()
	if _, err = rows.Columns(); err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}
	for rows.Next() {
		var projectID int
		err = rows.Scan(&projectID)
		if err != nil {
			Tx.Rollback()
			return false, "[ERROR]Fail to get data!", err
		}
		projectIDList += strconv.Itoa(projectID) + ","
	}
	if projectIDList != "" {
		projectIDList = projectIDList[:len(projectIDList)-1]
		// 更新项目更新时间
		_, err = Tx.Exec("UPDATE goku_gateway_project SET updateTime = ? WHERE projectID = ?;", now, projectIDList)
		if err != nil {
			Tx.Rollback()
			return false, "[ERROR]Fail to update data!", err
		}
	}
	Tx.Commit()
	return true, "", nil
}

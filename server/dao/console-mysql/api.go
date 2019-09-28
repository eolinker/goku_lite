package consolemysql

import (
	SQL "database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	database2 "github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

// AddAPI 新增接口
func AddAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, stripPrefix, stripSlash, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, managerID, userID int) (bool, int, error) {
	db := database2.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	res, err := Tx.Exec("INSERT INTO goku_gateway_api (projectID,groupID,apiName,requestURL,targetURL,requestMethod,targetMethod,protocol,stripSlash,balanceName,isFollow,stripPrefix,timeout,retryCount,alertValve,createTime,updateTime,managerID,lastUpdateUserID,createUserID) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);", projectID, groupID, apiName, requestURL, targetURL, requestMethod, targetMethod, protocol, stripSlash, balanceName, isFollow, stripPrefix, timeout, retryCount, alertValve, now, now, managerID, userID, userID)

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

// EditAPI 新增接口
func EditAPI(apiName, requestURL, targetURL, requestMethod, targetMethod, isFollow, stripPrefix, stripSlash, balanceName, protocol string, projectID, groupID, timeout, retryCount, alertValve, apiID, managerID, userID int) (bool, error) {
	db := database2.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	_, err := Tx.Exec("UPDATE goku_gateway_api SET projectID = ?,groupID = ?,apiName = ?,requestURL = ?,targetURL = ?,requestMethod = ?,protocol = ?,balanceName = ?,targetMethod = ?,isFollow = ?,stripPrefix = ?,stripSlash = ?,timeout = ?,retryCount = ?,alertValve = ?,updateTime = ?,managerID = ?,lastUpdateUserID = ? WHERE apiID = ?", projectID, groupID, apiName, requestURL, targetURL, requestMethod, protocol, balanceName, targetMethod, isFollow, stripPrefix, stripSlash, timeout, retryCount, alertValve, now, managerID, userID, apiID)

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
func GetAPIInfo(apiID int) (bool, *entity.API, error) {
	db := database2.GetConnection()
	sql := `SELECT goku_gateway_api.apiID,goku_gateway_api.groupID,goku_gateway_api.apiName,goku_gateway_api.requestURL,goku_gateway_api.targetURL,goku_gateway_api.requestMethod,goku_gateway_api.targetMethod,IFNULL(goku_gateway_api.protocol,"http"),goku_gateway_api.stripSlash,IFNULL(goku_gateway_api.balanceName,""),goku_gateway_api.isFollow,goku_gateway_api.stripPrefix,goku_gateway_api.timeout,goku_gateway_api.retryCount,goku_gateway_api.alertValve,goku_gateway_api.createTime,goku_gateway_api.updateTime,goku_gateway_api.managerID,goku_gateway_api.lastUpdateUserID,goku_gateway_api.createUserID,IFNULL(goku_gateway_api_group.groupPath,"0") FROM goku_gateway_api LEFT JOIN goku_gateway_api_group ON goku_gateway_api.groupID = goku_gateway_api_group.groupID WHERE goku_gateway_api.apiID = ?`
	api := &entity.API{}
	var managerInfo entity.ManagerInfo
	err := db.QueryRow(sql, apiID).Scan(&api.APIID, &api.GroupID, &api.APIName, &api.RequestURL, &api.ProxyURL, &api.RequestMethod, &api.TargetMethod, &api.Protocol, &api.StripSlash, &api.BalanceName, &api.IsFollow, &api.StripPrefix, &api.Timeout, &api.RetryConut, &api.Valve, &api.CreateTime, &api.UpdateTime, &managerInfo.ManagerID, &managerInfo.UpdaterID, &managerInfo.CreateUserID, &api.GroupPath)
	if err != nil {
		panic(err)
		return false, &entity.API{}, err
	}

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

// GetAPIListByGroupList 通过分组列表获取接口列表
func GetAPIListByGroupList(projectID int, groupIDList string) (bool, []map[string]interface{}, error) {
	db := database2.GetConnection()
	// 获取分组ID列表
	sql := `SELECT goku_gateway_api.apiID,goku_gateway_api.apiName,goku_gateway_api.requestURL,IFNULL(goku_gateway_api.updateTime,""),goku_gateway_api.lastUpdateUserID,goku_gateway_api.managerID FROM goku_gateway_api WHERE goku_gateway_api.projectID = ? AND goku_gateway_api.groupID IN (` + groupIDList + `) ORDER BY goku_gateway_api.updateTime DESC;`

	rows, err := db.Query(sql, projectID)
	if err != nil {
		return false, make([]map[string]interface{}, 0), err
	}
	defer rows.Close()
	apiList := make([]map[string]interface{}, 0)
	//获取记录列
	if _, err = rows.Columns(); err != nil {
		return false, make([]map[string]interface{}, 0), err
	}
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
func GetAPIIDList(projectID int, groupID int, keyword string, condition int, ids []int) (bool, []int, error) {
	db := database2.GetConnection()
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

	sql := fmt.Sprintf(`SELECT A.apiID FROM goku_gateway_api A %s`, ruleStr)
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
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
func GetAPIList(projectID int, groupID int, keyword string, condition, page, pageSize int, ids []int) (bool, []map[string]interface{}, int, error) {
	db := database2.GetConnection()
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

	sql := fmt.Sprintf(`SELECT A.apiID,A.apiName,A.requestURL,A.requestMethod,A.targetURL,IFNULL(A.balanceName,''),IFNULL(A.updateTime,""),IF(B.remark is null or B.remark = "",B.loginCall,B.remark) AS updaterName,IF(C.remark is null or C.remark = "",C.loginCall,C.remark) AS managerName,A.lastUpdateUserID,A.managerID,A.isFollow,IFNULL(A.protocol,"http"),A.targetMethod,A.groupID,IFNULL(D.groupPath,"0"),IFNULL(D.groupName,"未分组") FROM goku_gateway_api A INNER JOIN goku_admin B ON A.lastUpdateUserID = B.userID INNER JOIN goku_admin C ON A.managerID=C.userID LEFT JOIN goku_gateway_api_group D ON D.groupID = A.groupID %s`, ruleStr)
	count := getCountSQL(sql)
	rows, err := getPageSQL(sql, "A.updateTime", "DESC", page, pageSize)
	if err != nil {
		panic(err)
		return false, make([]map[string]interface{}, 0), 0, err
	}
	defer rows.Close()
	apiList := make([]map[string]interface{}, 0)
	//获取记录列
	for rows.Next() {
		var apiID, updaterID, managerID, groupID int
		var apiName, requestURL, updateTime, managerName, updaterName, requestMethod, targetURL, balanceName, targetMethod, protocol, groupPath, groupName string
		var isFollow bool

		err = rows.Scan(&apiID, &apiName, &requestURL, &requestMethod, &targetURL, &balanceName, &updateTime, &updaterName, &managerName, &updaterID, &managerID, &isFollow, &protocol, &targetMethod, &groupID, &groupPath, &groupName)
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
		}
		apiList = append(apiList, apiInfo)
	}
	return true, apiList, count, nil
}

// CheckURLIsExist 接口路径是否存在
func CheckURLIsExist(requestURL, requestMethod string, projectID, apiID int) bool {
	db := database2.GetConnection()
	var id int
	var m string
	sql := "SELECT apiID,requestMethod FROM goku_gateway_api WHERE requestURL = ? AND projectID = ?;"
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

// CheckAPIIsExist CheckAPIIsExist 检查接口是否存在
func CheckAPIIsExist(apiID int) (bool, error) {
	db := database2.GetConnection()
	sql := "SELECT apiID FROM goku_gateway_api WHERE apiID = ?;"
	var id int
	err := db.QueryRow(sql, apiID).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, err
}

// BatchEditAPIBalance 批量修改接口负载
func BatchEditAPIBalance(apiIDList []string, balance string) (string, error) {
	db := database2.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")

	sqlTpl := "UPDATE `goku_gateway_api` A LEFT JOIN `goku_gateway_project` P ON A.`projectID` = P.`projectID` SET A.`updateTime` = ?,  P.`updateTime`=?, A.`balanceName`=? WHERE A.`apiID` IN (%s)"

	sql := fmt.Sprint(sqlTpl, strings.Join(apiIDList, ","))

	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Fail to Prepare SQL!", err
	}
	_, err = stmt.Exec(now, now, balance)
	if err != nil {
		return "[ERROR]Fail to excute SQL statement!", err
	}
	return "", nil
}

// BatchEditAPIGroup 批量修改接口分组
func BatchEditAPIGroup(apiIDList []string, groupID int) (string, error) {
	db := database2.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")

	sqlTpl := "UPDATE `goku_gateway_api` A LEFT JOIN `goku_gateway_project` P ON A.`projectID` = P.`projectID` SET A.`updateTime` = ?,  P.`updateTime`=?, A.`groupID`=? WHERE A.`apiID` IN (%s)"

	sql := fmt.Sprintf(sqlTpl, strings.Join(apiIDList, ","))
	stmt, err := db.Prepare(sql)
	if err != nil {
		return "[ERROR]Fail to Prepare SQL!", err
	}
	_, err = stmt.Exec(now, now, groupID)
	if err != nil {
		return "[ERROR]Fail to excute SQL statement!", err
	}
	return "", nil

}

// BatchDeleteAPI 批量修改接口
func BatchDeleteAPI(apiIDList string) (bool, string, error) {
	db := database2.GetConnection()
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
	rows, err := db.Query("SELECT projectID FROM goku_gateway_api WHERE apiID IN (" + apiIDList + ");")
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

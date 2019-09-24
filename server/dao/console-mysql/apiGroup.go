package console_mysql

import (
	SQL "database/sql"
	"strconv"
	"time"

	database2 "github.com/eolinker/goku-api-gateway/common/database"
	log "github.com/eolinker/goku-api-gateway/goku-log"
)

// 新建接口分组
func AddApiGroup(groupName string, projectID, parentGroupID int) (bool, interface{}, error) {
	db := database2.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	groupPath := ""
	groupDepth := 1
	sql := ""

	// 查询父分组信息
	if parentGroupID > 0 {
		const sql = "SELECT groupDepth,groupPath FROM goku_gateway_api_group WHERE groupID = ?;"
		err := Tx.QueryRow(sql, parentGroupID).Scan(&groupDepth, &groupPath)
		if err != nil {
			if err != SQL.ErrNoRows {
				Tx.Rollback()
				return false, "[ERROR]Illegal SQL statement!", err
			}
		}
		if groupDepth > 4 {
			Tx.Rollback()
			return false, "[ERROR]Exceeding the grouping level!", err
		}
	}

	const sql2 = "INSERT INTO goku_gateway_api_group (projectID,groupName,parentGroupID) VALUES (?,?,?);"
	result, err := Tx.Exec(sql2, projectID, groupName, parentGroupID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Illegal SQL statement!", err
	}
	groupID, err := result.LastInsertId()
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to insert data!", err
	}
	if groupPath == "" {
		groupPath = strconv.Itoa(int(groupID))
	} else {
		groupDepth += 1
		groupPath += "," + strconv.Itoa(int(groupID))
	}

	// 更新groupDepth和groupPath
	sql = "UPDATE goku_gateway_api_group SET groupPath = ?,groupDepth = ? WHERE groupID = ?;"
	_, err = Tx.Exec(sql, groupPath, groupDepth, groupID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data!", err
	}
	// 更新项目更新时间
	_, err = Tx.Exec("UPDATE goku_gateway_project SET updateTime = ? WHERE projectID = ?;", now, projectID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data!", err
	}
	Tx.Commit()
	return true, groupID, nil
}

// 修改接口分组
func EditApiGroup(groupName string, groupID, projectID int) (bool, string, error) {
	db := database2.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	Tx, _ := db.Begin()
	sql := "UPDATE goku_gateway_api_group SET groupName = ? WHERE groupID = ? AND projectID = ?;"
	_, err := Tx.Exec(sql, groupName, groupID, projectID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data!", err
	}
	// 更新项目更新时间
	_, err = Tx.Exec("UPDATE goku_gateway_project SET updateTime = ? WHERE projectID = ?;", now, projectID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data!", err
	}
	Tx.Commit()
	return true, "", nil
}

// 删除接口分组
func DeleteApiGroup(projectID, groupID int) (bool, string, error) {
	db := database2.GetConnection()
	Tx, _ := db.Begin()
	var groupPath string
	// 获取分组信息
	sql := "SELECT groupPath FROM goku_gateway_api_group WHERE groupID = ?"
	err := Tx.QueryRow(sql, groupID).Scan(&groupPath)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to delete data!", err
	}
	var concatGroupID string
	sql = "SELECT GROUP_CONCAT(DISTINCT groupID) AS groupID FROM goku_gateway_api_group WHERE projectID = ? AND groupPath LIKE ?;"
	err = Tx.QueryRow(sql, projectID, groupPath+"%").Scan(&concatGroupID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to delete data!", err
	}
	sql = "DELETE FROM goku_gateway_api_group WHERE groupID IN (" + concatGroupID + ");"
	_, err = Tx.Exec(sql)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to delete data!", err
	}
	flag, apiList, _ := GetApiListByGroupList(projectID, concatGroupID)
	if flag {
		listLen := len(apiList)
		if listLen > 0 {
			apiIDList := ""
			for i, api := range apiList {
				apiIDList += strconv.Itoa(api["apiID"].(int))
				if i < listLen-1 {
					apiIDList += ","
				}
			}
			sql = "DELETE FROM goku_gateway_api WHERE apiID IN (" + apiIDList + ");"
			_, err = Tx.Exec(sql)
			if err != nil {
				Tx.Rollback()
				return false, "[ERROR]Fail to delete data!", err
			}
			//sql = "DELETE FROM goku_gateway_api_cache WHERE apiID IN (" + apiIDList + ");"
			//_, err = Tx.Exec(sql)
			//if err != nil {
			//	Tx.Rollback()
			//	return false, "[ERROR]Fail to delete data!", err
			//}

			sql = "DELETE FROM goku_conn_strategy_api WHERE apiID IN (" + apiIDList + ");"
			_, err = Tx.Exec(sql)
			if err != nil {
				Tx.Rollback()
				return false, "[ERROR]Fail to delete data!", err
			}

			sql = "DELETE FROM goku_conn_plugin_api WHERE apiID IN (" + apiIDList + ");"
			_, err = Tx.Exec(sql)
			if err != nil {
				Tx.Rollback()
				return false, "[ERROR]Fail to delete data!", err
			}
			//sql = "DELETE FROM goku_conn_plugin_api_cache WHERE apiID IN (" + apiIDList + ");"
			//_, err = Tx.Exec(sql)
			//if err != nil {
			//	Tx.Rollback()
			//	return false, "[ERROR]Fail to delete data!", err
			//}
		}
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	// 更新项目更新时间
	_, err = Tx.Exec("UPDATE goku_gateway_project SET updateTime = ? WHERE projectID = ?;", now, projectID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data!", err
	}
	// 获取接口分组下面的接口ID
	Tx.Commit()
	return true, "", nil
}

// 获取接口分组列表
func GetApiGroupList(projectID int) (bool, []map[string]interface{}, error) {
	db := database2.GetConnection()
	sql := "SELECT groupID,groupName,parentGroupID,groupDepth FROM goku_gateway_api_group WHERE projectID = ?;"
	rows, err := db.Query(sql, projectID)
	if err != nil {
		return false, make([]map[string]interface{}, 0), err
	}
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
		info := err.Error()
		log.Info(info)
		return false, make([]map[string]interface{}, 0), err
	} else {
		groupList := make([]map[string]interface{}, 0)
		for rows.Next() {
			var groupID, parentGroupID, groupDepth int
			var groupName string
			err = rows.Scan(&groupID, &groupName, &parentGroupID, &groupDepth)
			if err != nil {
				return false, make([]map[string]interface{}, 0), err
			}
			groupInfo := map[string]interface{}{
				"groupID":       groupID,
				"groupName":     groupName,
				"groupDepth":    groupDepth,
				"parentGroupID": parentGroupID,
			}
			groupList = append(groupList, groupInfo)
		}
		return true, groupList, nil
	}
}

// 更新接口分组脚本
func UpdateApiGroupScript() bool {
	db := database2.GetConnection()
	// 获取一级分组
	sql := "SELECT groupID,groupName,parentGroupID FROM goku_gateway_api_group WHERE isChild = ?;"
	rows, err := db.Query(sql, 0)
	if err != nil {
		return false
	}
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
		return false
	} else {
		groupList := make([]map[string]interface{}, 0)
		for rows.Next() {
			var groupID, parentGroupID int
			var groupName, groupPath string
			err = rows.Scan(&groupID, &groupName, &parentGroupID)
			if err != nil {
				return false
			}
			sql = "SELECT groupID,groupName,parentGroupID FROM goku_gateway_api_group WHERE isChild = ? AND parentGroupID = ?;"
			r, err := db.Query(sql, 1, groupID)
			if err != nil {
				return false
			}
			groupPath = strconv.Itoa(groupID)
			defer r.Close()
			if _, err = r.Columns(); err != nil {
				return false
			} else {
				for r.Next() {
					var childGroupID, childParentGroupID int
					var childGroupName, childGroupPath string
					err = r.Scan(&childGroupID, &childGroupName, &childParentGroupID)
					if err != nil {
						return false
					}
					childGroupPath = groupPath + "," + strconv.Itoa(childGroupID)
					sql = "SELECT groupID,groupName,parentGroupID FROM goku_gateway_api_group WHERE isChild = ? AND parentGroupID = ?;"
					rw, err := db.Query(sql, 2, childGroupID)
					if err != nil {
						return false
					}
					defer rw.Close()
					if _, err = rw.Columns(); err != nil {
						return false
					} else {
						for rw.Next() {
							var secChildGroupID, secChildParentGroupID int
							var secChildGroupName, secChildGroupPath string
							err = rw.Scan(&secChildGroupID, &secChildGroupName, &secChildParentGroupID)
							if err != nil {
								return false
							}
							secChildGroupPath = childGroupPath + "," + strconv.Itoa(secChildGroupID)
							groupList = append(groupList, map[string]interface{}{
								"groupID":       secChildGroupID,
								"groupName":     secChildGroupName,
								"groupDepth":    3,
								"parentGroupID": secChildParentGroupID,
								"groupPath":     secChildGroupPath,
							})
						}
					}
					groupList = append(groupList, map[string]interface{}{
						"groupID":       childGroupID,
						"groupName":     childGroupName,
						"groupDepth":    2,
						"parentGroupID": childParentGroupID,
						"groupPath":     childGroupPath,
					})
				}
			}
			groupInfo := map[string]interface{}{
				"groupID":       groupID,
				"groupName":     groupName,
				"groupDepth":    1,
				"parentGroupID": parentGroupID,
				"groupPath":     groupPath,
			}
			groupList = append(groupList, groupInfo)
		}
		Tx, _ := db.Begin()
		for _, groupInfo := range groupList {
			_, err = Tx.Exec("UPDATE goku_gateway_api_group SET groupPath = ?,groupDepth =? WHERE groupID = ?;", groupInfo["groupPath"].(string), groupInfo["groupDepth"].(int), groupInfo["groupID"].(int))
			if err != nil {
				Tx.Rollback()

			}
		}
		Tx.Commit()
		return true
	}
}

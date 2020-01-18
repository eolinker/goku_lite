package console_sqlite3

import (
	SQL "database/sql"
	"strconv"
	"time"

	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
)

var (
	apiDao dao.APIDao
)

func init() {
	pdao.Need(&apiDao)
}

//APIGroupDao APIGroupDao
type APIGroupDao struct {
	db *SQL.DB
}

//NewAPIGroupDao new APIGroupDao
func NewAPIGroupDao() *APIGroupDao {
	return &APIGroupDao{}
}

//Create create
func (d *APIGroupDao) Create(db *SQL.DB) (interface{}, error) {
	d.db = db
	var i dao.APIGroupDao = d
	return &i, nil
}

//AddAPIGroup 新建接口分组
func (d *APIGroupDao) AddAPIGroup(groupName string, projectID, parentGroupID int) (bool, interface{}, error) {
	db := d.db
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
		groupDepth = groupDepth + 1
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

//EditAPIGroup 修改接口分组
func (d *APIGroupDao) EditAPIGroup(groupName string, groupID, projectID int) (bool, string, error) {
	db := d.db
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

//DeleteAPIGroup 删除接口分组
func (d *APIGroupDao) DeleteAPIGroup(projectID, groupID int) (bool, string, error) {
	db := d.db
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
	flag, apiList, _ := apiDao.GetAPIListByGroupList(projectID, concatGroupID)
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

//GetAPIGroupList 获取接口分组列表
func (d *APIGroupDao) GetAPIGroupList(projectID int) (bool, []map[string]interface{}, error) {
	db := d.db
	sql := "SELECT groupID,groupName,parentGroupID,groupDepth FROM goku_gateway_api_group WHERE projectID = ?;"
	rows, err := db.Query(sql, projectID)
	if err != nil {
		return false, make([]map[string]interface{}, 0), err
	}
	defer rows.Close()
	//获取记录列

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

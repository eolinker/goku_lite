package console_sqlite3

import (
	SQL "database/sql"
	"fmt"

	"strconv"
	"strings"
	"time"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

//ProjectDao ProjectDao
type ProjectDao struct {
	db *SQL.DB
}

//NewProjectDao new ProjectDao
func NewProjectDao() *ProjectDao {
	return &ProjectDao{}
}

//Create create
func (d *ProjectDao) Create(db *SQL.DB) (interface{}, error) {
	d.db = db
	var i dao.ProjectDao = d
	return &i, nil
}

//AddProject 新建项目
func (d *ProjectDao) AddProject(projectName string) (bool, interface{}, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	db := d.db
	sql := "INSERT INTO goku_gateway_project (projectName,createTime,updateTime) VALUES (?,?,?);"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, err.Error(), err
	}
	defer stmt.Close()
	r, err := stmt.Exec(projectName, now, now)
	if err != nil {
		return false, "[ERROR]Fail to insert data!", err
	}
	projectID, _ := r.LastInsertId()
	return true, projectID, nil
}

//EditProject 修改项目信息
func (d *ProjectDao) EditProject(projectName string, projectID int) (bool, string, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	db := d.db
	sql := "UPDATE goku_gateway_project SET projectName = ?,updateTime = ? WHERE projectID = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, err.Error(), err
	}
	defer stmt.Close()
	_, err = stmt.Exec(projectName, now, projectID)
	if err != nil {
		return false, "[ERROR]Fail to update data!", err
	}
	return true, "", nil
}

//DeleteProject 修改项目信息
func (d *ProjectDao) DeleteProject(projectID int) (bool, string, error) {
	db := d.db
	Tx, _ := db.Begin()
	// 获取项目分组列表
	sql := "SELECT groupID FROM goku_gateway_api_group WHERE projectID = ?;"
	rows, err := Tx.Query(sql, projectID)
	if err != nil {
		Tx.Rollback()
		return false, "", err
	}
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	groupIDList := ""

	for rows.Next() {
		var groupID int
		err = rows.Scan(&groupID)
		if err != nil {
			Tx.Rollback()
			log.Info(err.Error())
			return false, "", err
		}
		groupIDList += strconv.Itoa(groupID) + ","
	}
	groupLen := len(groupIDList)
	if groupLen > 0 {
		if string(groupIDList[groupLen-1]) == "," {
			groupIDList = groupIDList[:groupLen-1]
		}
		sql = "DELETE FROM goku_gateway_api_group WHERE groupID IN (" + groupIDList + ");"
		_, err := Tx.Exec(sql)
		if err != nil {
			Tx.Rollback()
			return false, "[ERROR]Fail to excute SQL statement!", err
		}
		// 获取接口ID列表
		sql = "SELECT apiID FROM goku_gateway_api WHERE projectID = ?;"
		r, err := Tx.Query(sql, projectID)
		if err != nil {
			Tx.Rollback()
			return false, "", err
		}
		if _, err = r.Columns(); err != nil {
			Tx.Rollback()
			return false, "", err
		}
		apiIDList := ""
		for r.Next() {
			var apiID int
			err = r.Scan(&apiID)
			if err != nil {
				Tx.Rollback()
				log.Info(err.Error())
			}
			apiIDList += strconv.Itoa(apiID) + ","
		}
		apiLen := len(apiIDList)
		if apiLen > 0 {
			if string(apiIDList[apiLen-1]) == "," {
				apiIDList = apiIDList[:apiLen-1]
			}
			sql = "DELETE FROM goku_gateway_api WHERE apiID IN (" + apiIDList + ");"
			_, err = Tx.Exec(sql)
			if err != nil {
				Tx.Rollback()
				return false, "[ERROR]Fail to excute SQL statement!", err
			}

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

	sql = "DELETE FROM goku_gateway_project WHERE projectID = ?;"
	_, err = Tx.Exec(sql, projectID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	Tx.Commit()
	return true, "", nil
}

//BatchDeleteProject 批量删除项目
func (d *ProjectDao) BatchDeleteProject(projectIDList string) (bool, string, error) {
	db := d.db
	Tx, _ := db.Begin()
	// 获取项目分组列表
	sql := "SELECT groupID FROM goku_gateway_api_group WHERE projectID IN (" + projectIDList + ");"
	rows, err := Tx.Query(sql)
	if err != nil {
		Tx.Rollback()
		return false, "", err
	}
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	groupIDList := ""
	if _, err = rows.Columns(); err != nil {
		Tx.Rollback()
		return false, "", err
	}
	for rows.Next() {
		var groupID int
		err = rows.Scan(&groupID)
		if err != nil {
			Tx.Rollback()
			return false, "", err
		}
		groupIDList += strconv.Itoa(groupID) + ","
	}
	groupLen := len(groupIDList)
	if groupLen > 0 && string(groupIDList[groupLen-1]) == "," {
		groupIDList = groupIDList[:groupLen-1]
		sql = "DELETE FROM goku_gateway_api_group WHERE groupID IN (" + groupIDList + ");"
		_, err := Tx.Exec(sql)
		if err != nil {
			Tx.Rollback()
			return false, "[ERROR]Fail to excute SQL statement!", err
		}
	}
	// 获取接口ID列表
	sql = "SELECT apiID FROM goku_gateway_api WHERE projectID IN (" + projectIDList + ");"
	r, err := Tx.Query(sql)
	if err != nil {
		Tx.Rollback()
		return false, "", err
	}
	if _, err = r.Columns(); err != nil {
		Tx.Rollback()
		return false, "", err
	}
	apiIDList := ""
	for r.Next() {
		var apiID int
		err = r.Scan(&apiID)
		if err != nil {
			log.Info(err.Error())
		}
		apiIDList += strconv.Itoa(apiID) + ","
	}
	apiLen := len(apiIDList)
	if apiLen != 0 && string(apiIDList[apiLen-1]) == "," {
		apiIDList = apiIDList[:apiLen-1]
		sql = "DELETE FROM goku_gateway_api WHERE apiID IN (" + apiIDList + ");"
		_, err = Tx.Exec(sql)
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
	}

	sql = "DELETE FROM goku_gateway_project WHERE projectID IN (" + projectIDList + ");"
	_, err = Tx.Exec(sql)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to excute SQL statement!", err
	}

	Tx.Commit()
	return true, "", nil
}

//GetProjectInfo 获取项目信息
func (d *ProjectDao) GetProjectInfo(projectID int) (bool, entity.Project, error) {
	db := d.db
	var project entity.Project
	sql := "SELECT projectID,projectName,createTime,updateTime FROM goku_gateway_project WHERE projectID = ?;"
	err := db.QueryRow(sql, projectID).Scan(&project.ProjectID, &project.ProjectName, &project.CreateTime, &project.UpdateTime)
	if err != nil {
		return false, entity.Project{}, err
	}
	return true, project, nil
}

//GetProjectList 获取项目列表
func (d *ProjectDao) GetProjectList(keyword string) (bool, []*entity.Project, error) {

	sql := "SELECT `projectID`,`projectName`,`updateTime` FROM `goku_gateway_project` %s ORDER BY `updateTime` DESC;"
	keywordValue := strings.Trim(keyword, "%")
	arg := []interface{}{}
	where := ""
	if keywordValue != "" {

		kvp := fmt.Sprint("%", keywordValue, "%")
		where = fmt.Sprint("WHERE `projectName` LIKE ?")
		arg = []interface{}{
			kvp,
		}
	}
	sql = fmt.Sprintf(sql, where)
	db := d.db
	rows, err := db.Query(sql, arg...)
	if err != nil {
		return false, nil, err
	}
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	projectList := make([]*entity.Project, 0, 25)
	for rows.Next() {
		var project entity.Project
		err = rows.Scan(&project.ProjectID, &project.ProjectName, &project.UpdateTime)
		if err != nil {
			return false, nil, err
		}
		projectList = append(projectList, &project)
	}
	return true, projectList, nil

}

//CheckProjectIsExist 检查项目是否存在
func (d *ProjectDao) CheckProjectIsExist(projectID int) (bool, error) {
	db := d.db
	sql := "SELECT projectID FROM goku_gateway_project WHERE projectID = ?;"
	var id int
	err := db.QueryRow(sql, projectID).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, err
}

//GetAPIListFromProjectNotInStrategy 获取项目列表中没有被策略组绑定的接口
func (d *ProjectDao) GetAPIListFromProjectNotInStrategy() (bool, []map[string]interface{}, error) {
	db := d.db
	sql := "SELECT projectID,projectName FROM goku_gateway_project;"
	projectRows, err := db.Query(sql)
	if err != nil {
		return false, nil, err
	}
	//延时关闭Rows
	defer projectRows.Close()
	//获取记录列
	projectList := make([]map[string]interface{}, 0, 20)

	for projectRows.Next() {
		var projectID int
		var projectName string
		err = projectRows.Scan(&projectID, &projectName)
		if err != nil {
			return false, nil, err
		}
		sql = "SELECT groupID,groupName,parentGroupID,groupDepth FROM goku_gateway_api_group WHERE projectID = ?;"
		rows, err := db.Query(sql, projectID)
		if err != nil {
			return false, nil, err
		}
		defer rows.Close()
		//获取记录列
		groupList := make([]map[string]interface{}, 0, 20)
		for rows.Next() {
			var groupID, parentGroupID, groupDepth int
			var groupName string
			err = rows.Scan(&groupID, &groupName, &parentGroupID, &groupDepth)
			if err != nil {
				return false, nil, err
			}
			groupInfo := map[string]interface{}{
				"groupID":       groupID,
				"groupName":     groupName,
				"groupDepth":    groupDepth,
				"parentGroupID": parentGroupID,
			}
			groupList = append(groupList, groupInfo)
		}
		projectInfo := map[string]interface{}{
			"projectID":   projectID,
			"projectName": projectName,
			"groupList":   groupList,
		}
		projectList = append(projectList, projectInfo)
	}
	return true, projectList, nil
}

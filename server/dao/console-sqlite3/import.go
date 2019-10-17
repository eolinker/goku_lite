package console_sqlite3

import (
	SQL "database/sql"
	"net/url"
	"strconv"
	"strings"
	"time"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	database2 "github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

var method = []string{"POST", "GET", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH"}

// 导入接口信息
func importAPIInfo(Tx *SQL.Tx, api entity.AmsAPIInfo, projectID, groupID, userID int, now string) bool {
	// 新增API
	requestURL := ""
	host := ""
	protocol := "http"
	u, err := url.ParseRequestURI(api.BaseInfo.APIURI)
	if err == nil {
		requestURL = u.Path
		if u.Scheme != "" {
			protocol = strings.ToLower(u.Scheme)
			if u.Host != "" {
				host = strings.ToLower(u.Host)
			}
		}
	} else {
		requestURL = api.BaseInfo.APIURI
	}
	stripSlash := true
	log.Debug(protocol, host, stripSlash)
	requestMethod := method[api.BaseInfo.APIRequestType]
	_, err = Tx.Exec("INSERT INTO goku_gateway_api (projectID,groupID,apiName,requestURL,targetURL,requestMethod,targetMethod,isFollow,stripPrefix,timeout,retryCount,createTime,updateTime,protocol,balanceName,stripSlash,responseDataType,managerID,lastUpdateUserID,createUserID) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);", projectID, groupID, api.BaseInfo.APIName, requestURL, requestURL, requestMethod, requestMethod, "true", "true", 2000, 0, now, now, protocol, host, stripSlash, "origin", userID, userID, userID)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func recursiveImportAPIGroupFromAms(Tx *SQL.Tx, projectID, userID int, groupInfo entity.AmsGroupInfo, groupDepth, parentGroupID int, groupPath, now string) (bool, string, error) {
	// 插入分组信息
	result, err := Tx.Exec("INSERT INTO goku_gateway_api_group (projectID,groupName,groupDepth,parentGroupID) VALUES (?,?,?,?);", projectID, groupInfo.GroupName, groupDepth, parentGroupID)
	if err != nil {
		info := err.Error()
		log.Info(info)
		return false, err.Error(), err
	}
	groupID, err := result.LastInsertId()
	if err != nil {
		info := err.Error()
		log.Info(info)
		return false, err.Error(), err
	}
	if groupPath == "" {
		groupPath = strconv.Itoa(int(groupID))
	} else {
		groupPath = groupPath + "," + strconv.Itoa(int(groupID))
	}

	// 更新groupPath
	_, err = Tx.Exec("UPDATE goku_gateway_api_group SET groupPath = ? WHERE groupID = ?;", groupPath, groupID)
	if err != nil {
		info := err.Error()
		log.Info(info)
		return false, err.Error(), err
	}
	for _, childGroup := range groupInfo.APIGroupChildList {
		_, _, err := recursiveImportAPIGroupFromAms(Tx, projectID, userID, childGroup, groupDepth+1, int(groupID), groupPath, now)
		if err != nil {
			continue
		}
	}
	for _, childGroup := range groupInfo.ChildGroupList {
		_, _, err := recursiveImportAPIGroupFromAms(Tx, projectID, userID, childGroup, groupDepth+1, int(groupID), groupPath, now)
		if err != nil {
			continue
		}
	}
	for _, api := range groupInfo.APIList {
		flag := importAPIInfo(Tx, api, projectID, int(groupID), userID, now)
		if !flag {
			continue
		}
	}
	return true, "", nil
}

//ImportAPIGroupFromAms 导入分组
func ImportAPIGroupFromAms(projectID, userID int, groupInfo entity.AmsGroupInfo) (bool, string, error) {
	db := database2.GetConnection()
	Tx, _ := db.Begin()
	now := time.Now().Format("2006-01-02 15:04:05")
	_, errInfo, err := recursiveImportAPIGroupFromAms(Tx, projectID, userID, groupInfo, 1, 0, "", now)
	if err != nil {
		Tx.Rollback()
		return false, errInfo, err
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

//ImportProjectFromAms 导入项目
func ImportProjectFromAms(userID int, projectInfo entity.AmsProject) (bool, string, error) {
	db := database2.GetConnection()
	Tx, _ := db.Begin()
	now := time.Now().Format("2006-01-02 15:04:05")
	// 插入项目信息
	projectResult, err := Tx.Exec("INSERT INTO goku_gateway_project (projectName,updateTime,createTime) VALUES (?,?,?)", projectInfo.ProjectInfo.ProjectName, now, now)
	if err != nil {
		Tx.Rollback()
		log.Info(err.Error())
		return false, err.Error(), err
	}
	projectID, err := projectResult.LastInsertId()
	if err != nil {
		Tx.Rollback()
		log.Info(err.Error())
		return false, err.Error(), err
	}
	id := int(projectID)
	for _, groupInfo := range projectInfo.APIGroupList {
		_, _, err := recursiveImportAPIGroupFromAms(Tx, id, userID, groupInfo, 1, 0, "", now)
		if err != nil {
			continue
		}
	}
	Tx.Commit()
	return true, "", nil
}

//ImportAPIFromAms 从ams中导入接口
func ImportAPIFromAms(projectID, groupID, userID int, apiList []entity.AmsAPIInfo) (bool, string, error) {
	db := database2.GetConnection()
	Tx, _ := db.Begin()
	now := time.Now().Format("2006-01-02 15:04:05")
	for _, a := range apiList {
		flag := importAPIInfo(Tx, a, projectID, groupID, userID, now)
		if !flag {
			continue
		}
	}
	Tx.Commit()
	return true, "", nil
}

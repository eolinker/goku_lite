package consolemysql

import (
	SQL "database/sql"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"net/url"
	"strconv"
	"strings"
	"time"

	database2 "github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/console-entity"
)

var method = []string{"POST", "GET", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH"}

// importAPIGroupInfo 导入分组信息
func importAPIGroupInfo(Tx *SQL.Tx, groupName, groupPath string, parentGroupID, projectID, groupDepth int) (bool, int, string) {
	result, err := Tx.Exec("INSERT INTO goku_gateway_api_group (projectID,groupName,parentGroupID,groupDepth) VALUES (?,?,?,?);", projectID, groupName, parentGroupID, groupDepth)
	if err != nil {
		info := err.Error()
		log.Info(info)
		return false, 0, ""
	}
	groupID, err := result.LastInsertId()
	if err != nil {
		info := err.Error()
		log.Info(info)
		return false, 0, ""
	}
	id := int(groupID)
	groupPath = groupPath + "," + strconv.Itoa(id)
	// 更新groupPath
	_, err = Tx.Exec("UPDATE goku_gateway_api_group SET groupPath = ? WHERE groupID = ?;", groupPath, id)
	if err != nil {
		info := err.Error()
		log.Info(info)
		return false, 0, ""
	}
	return true, id, groupPath
}

// importAPIInfo 导入接口信息
func importAPIInfo(Tx *SQL.Tx, api entity.AmsAPIInfo, projectID, groupID, userID int) bool {
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
	now := time.Now().Format("2006-01-02 15:04:05")
	requestMethod := method[api.BaseInfo.APIRequestType]
	_, err = Tx.Exec("INSERT INTO goku_gateway_api (projectID,groupID,apiName,requestURL,targetURL,requestMethod,targetMethod,isFollow,stripPrefix,timeout,retryCount,createTime,updateTime,managerID,lastUpdateUserID,createUserID,protocol,balanceName,stripSlash) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);", projectID, groupID, api.BaseInfo.APIName, requestURL, requestURL, requestMethod, requestMethod, "true", "true", 2000, 0, now, now, userID, userID, userID, protocol, host, stripSlash)
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

//ImportAPIGroupFromAms 导入接口分组
func ImportAPIGroupFromAms(projectID, userID int, groupInfo entity.AmsGroupInfo) (bool, string, error) {
	db := database2.GetConnection()
	Tx, _ := db.Begin()
	now := time.Now().Format("2006-01-02 15:04:05")
	// 插入分组信息
	result, err := Tx.Exec("INSERT INTO goku_gateway_api_group (projectID,groupName,groupDepth) VALUES (?,?,1);", projectID, groupInfo.GroupName)
	if err != nil {
		Tx.Rollback()
		info := err.Error()
		log.Info(info)
	}
	groupID, err := result.LastInsertId()
	if err != nil {
		Tx.Rollback()
		info := err.Error()
		log.Info(info)
	}
	groupPath := strconv.Itoa(int(groupID))
	// 更新groupPath
	_, err = Tx.Exec("UPDATE goku_gateway_api_group SET groupPath = ? WHERE groupID = ?;", groupPath, groupID)
	if err != nil {
		Tx.Rollback()
		info := err.Error()
		log.Info(info)
	}
	// 插入子分组信息
	for _, value := range groupInfo.ChildGroupList {
		flag, secGroupID, secgroupPath := importAPIGroupInfo(Tx, value.GroupName, groupPath, int(groupID), projectID, 2)
		if !flag {
			continue
		}
		for _, v := range value.ChildGroupList {
			flag, thirdGroupID, thirdgroupPath := importAPIGroupInfo(Tx, v.GroupName, secgroupPath, secGroupID, projectID, 3)
			if !flag {
				continue
			}
			for _, vv := range v.ChildGroupList {
				flag, fourthGroupID, fourthgroupPath := importAPIGroupInfo(Tx, vv.GroupName, thirdgroupPath, thirdGroupID, projectID, 4)
				if !flag {
					continue
				}
				for _, vvv := range vv.ChildGroupList {
					flag, fifthGroupID, _ := importAPIGroupInfo(Tx, vvv.GroupName, fourthgroupPath, fourthGroupID, projectID, 5)
					if !flag {
						continue
					}
					for _, aaa := range vvv.APIList {
						flag = importAPIInfo(Tx, aaa, projectID, fifthGroupID, userID)
						if !flag {
							continue
						}
					}
				}
				for _, aa := range vv.APIList {
					flag = importAPIInfo(Tx, aa, projectID, fourthGroupID, userID)
					if !flag {
						continue
					}
				}
			}
			for _, a := range v.APIList {
				flag = importAPIInfo(Tx, a, projectID, thirdGroupID, userID)
				if !flag {
					continue
				}
			}
		}
		for _, api := range value.APIList {
			flag := importAPIInfo(Tx, api, projectID, secGroupID, userID)
			if !flag {
				continue
			}
		}
	}
	for _, apiInfo := range groupInfo.APIList {
		flag := importAPIInfo(Tx, apiInfo, projectID, int(groupID), userID)
		if !flag {
			continue
		}
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

// ImportProjectFromAms 导入项目
func ImportProjectFromAms(userID int, projectInfo entity.AmsProject) (bool, string, error) {
	db := database2.GetConnection()
	Tx, _ := db.Begin()
	now := time.Now().Format("2006-01-02 15:04:05")
	// 插入项目信息
	projectResult, err := Tx.Exec("INSERT INTO goku_gateway_project (projectName,updateTime,createTime) VALUES (?,?,?)", projectInfo.ProjectInfo.ProjectName, now, now)
	if err != nil {
		Tx.Rollback()
		log.Info(err.Error())
	}
	projectID, err := projectResult.LastInsertId()
	if err != nil {
		Tx.Rollback()
		log.Info(err.Error())
	}
	id := int(projectID)
	for _, groupInfo := range projectInfo.APIGroupList {
		// 插入分组信息
		result, err := Tx.Exec("INSERT INTO goku_gateway_api_group (projectID,groupName,groupDepth) VALUES (?,?,1);", projectID, groupInfo.GroupName)
		if err != nil {
			Tx.Rollback()
			log.Info(err.Error())
		}
		groupID, err := result.LastInsertId()
		if err != nil {
			Tx.Rollback()
			log.Info(err.Error())
		}
		groupPath := strconv.Itoa(int(groupID))
		// 更新groupPath
		_, err = Tx.Exec("UPDATE goku_gateway_api_group SET groupPath = ? WHERE groupID = ?;", groupPath, groupID)
		if err != nil {
			Tx.Rollback()
			info := err.Error()
			log.Info(info)
		}

		// 插入子分组信息
		for _, value := range groupInfo.APIGroupChildList {
			flag, secGroupID, secgroupPath := importAPIGroupInfo(Tx, value.GroupName, groupPath, int(groupID), id, 2)
			if !flag {
				continue
			}
			for _, v := range value.APIGroupChildList {
				flag, thirdGroupID, thirdgroupPath := importAPIGroupInfo(Tx, v.GroupName, secgroupPath, secGroupID, id, 3)
				if !flag {
					continue
				}
				for _, vv := range v.APIGroupChildList {
					flag, fourthGroupID, fourthgroupPath := importAPIGroupInfo(Tx, vv.GroupName, thirdgroupPath, thirdGroupID, id, 4)
					if !flag {
						continue
					}
					for _, vvv := range vv.APIGroupChildList {
						flag, fifthGroupID, _ := importAPIGroupInfo(Tx, vvv.GroupName, fourthgroupPath, fourthGroupID, id, 5)
						if !flag {
							continue
						}
						for _, aaa := range v.APIList {
							flag = importAPIInfo(Tx, aaa, id, fifthGroupID, userID)
							if !flag {
								continue
							}
						}
					}
					for _, aa := range v.APIList {
						flag = importAPIInfo(Tx, aa, id, fourthGroupID, userID)
						if !flag {
							continue
						}
					}
				}
				for _, a := range v.APIList {
					flag = importAPIInfo(Tx, a, id, thirdGroupID, userID)
					if !flag {
						continue
					}
				}
			}
			for _, api := range value.APIList {
				flag = importAPIInfo(Tx, api, id, secGroupID, userID)
				if !flag {
					continue
				}
			}
		}
		for _, apiInfo := range groupInfo.APIList {
			flag := importAPIInfo(Tx, apiInfo, id, int(groupID), userID)
			if !flag {
				continue
			}
		}
	}
	Tx.Commit()
	return true, "", nil
}

//ImportAPIFromAms 导入接口
func ImportAPIFromAms(projectID, groupID, userID int, apiList []entity.AmsAPIInfo) (bool, string, error) {
	db := database2.GetConnection()
	Tx, _ := db.Begin()
	now := time.Now().Format("2006-01-02 15:04:05")
	for _, apiInfo := range apiList {
		// 新增API
		requestURL := ""
		host := ""
		u, err := url.ParseRequestURI(apiInfo.BaseInfo.APIURI)
		if err == nil {
			requestURL = u.Path
			if u.Scheme != "" && u.Host != "" {
				host = u.Scheme + "://" + u.Host
			}
		} else {
			requestURL = apiInfo.BaseInfo.APIURI
		}
		now = time.Now().Format("2006-01-02 15:04:05")
		requestMethod := method[apiInfo.BaseInfo.APIRequestType]
		_, err = Tx.Exec("INSERT INTO goku_gateway_api (projectID,groupID,apiName,requestURL,targetURL,requestMethod,targetServer,targetMethod,isFollow,stripPrefix,timeout,retryCount,createTime,updateTime,managerID,lastUpdateUserID,createUserID) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);", projectID, groupID, apiInfo.BaseInfo.APIName, requestURL, requestURL, requestMethod, host, requestMethod, "true", "true", 2000, 0, now, now, userID, userID, userID)
		if err != nil {
			Tx.Rollback()
			log.Info(err.Error())
		}
	}
	// 更新项目更新时间
	_, err := Tx.Exec("UPDATE goku_gateway_project SET updateTime = ? WHERE projectID = ?;", now, projectID)
	if err != nil {
		Tx.Rollback()
		return false, "[ERROR]Fail to update data!", err
	}
	Tx.Commit()
	return true, "", nil
}

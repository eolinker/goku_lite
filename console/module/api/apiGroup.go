package api

//AddAPIGroup 新建接口分组
func AddAPIGroup(groupName string, projectID, parentGroupID int) (bool, interface{}, error) {
	return apiGroupDao.AddAPIGroup(groupName, projectID, parentGroupID)
}

//EditAPIGroup 修改接口分组
func EditAPIGroup(groupName string, groupID, projectID int) (bool, string, error) {
	return apiGroupDao.EditAPIGroup(groupName, groupID, projectID)
}

//DeleteAPIGroup 删除接口分组
func DeleteAPIGroup(projectID, groupID int) (bool, string, error) {
	flag, result, err := apiGroupDao.DeleteAPIGroup(projectID, groupID)

	return flag, result, err
}

//GetAPIGroupList 获取接口分组列表
func GetAPIGroupList(projectID int) (bool, []map[string]interface{}, error) {
	return apiGroupDao.GetAPIGroupList(projectID)
}

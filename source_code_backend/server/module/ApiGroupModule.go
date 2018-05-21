package module

import (
	"goku-ce/server/dao"
	"goku-ce/server/conf"
)

// 新增分组
func AddApiGroup(gatewayAlias,groupName string) (bool,int) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.AddApiGroup(gateway["apiGroupConfPath"],groupName)
	}else {
		return false,0
	}

}

// 修改分组
func EditApiGroup(gatewayAlias,groupName string,groupID int) (bool) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.EditApiGroup(gateway["apiGroupConfPath"],groupName,groupID)
	}else {
		return false
	}
}

// 删除分组
func DeleteApiGroup(gatewayAlias string,groupID int) (bool) {
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		flag = dao.DeleteApiGroup(gateway["apiGroupConfPath"],groupID)
		if flag{
			return dao.DeleteApiOfGroup(gateway["apiConfPath"],groupID)
		} else {
			return false
		}
	}else {
		return false
	}
}


// 获取分组列表
func GetApiGroupList(gatewayAlias string) []*conf.GroupInfo{
	// 获取网关配置路径
	flag,gateway := dao.GetGatewayConfPath(gatewayAlias)
	if flag {
		return dao.GetApiGroupList(gateway["apiGroupConfPath"])
	}else {
		return make([]*conf.GroupInfo,0)
	}
}
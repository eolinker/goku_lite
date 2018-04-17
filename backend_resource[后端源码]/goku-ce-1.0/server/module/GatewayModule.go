package module

import (
	"goku-ce-1.0/utils"
	"goku-ce-1.0/server/dao"
	"time"
)
// 新增网关
func Addgateway(gatewayName,gatewayDesc,gatewayAlias string,userID int) (bool,string){
	createTime := time.Now().Format("2006-01-02 15:04:05")
	hashKey := utils.GetHashKey(gatewayName,gatewayDesc,gatewayAlias)
	// token := utils.GetHashKey(gatewayName)
	if flag,_ :=dao.Addgateway(gatewayName,gatewayDesc,gatewayAlias,createTime,hashKey,userID);flag{
		return true,hashKey
	}else{
		return true,""
	}
}

// 修改网关
func EditGateway(gatewayName,gatewayAlias,gatewayDesc,gatewayHashKey string,userID int) (bool){
	flag,gatewayID := dao.GetIDFromHashKey(gatewayHashKey)
	if flag{
		if dao.CheckGatewayPermission(gatewayID,userID){
			return dao.EditGateway(gatewayName,gatewayAlias,gatewayDesc,gatewayHashKey)
		}else{
			return false
		}
	}else{
		return false
	}
}

// 删除网关
func DeleteGateway(gatewayHashKey string,userID int) (bool){
	flag,gatewayID := dao.GetIDFromHashKey(gatewayHashKey)
	if flag{
		if dao.CheckGatewayPermission(gatewayID,userID){
			return dao.DeleteGateway(gatewayHashKey)
		}else{
			return false
		}
	}else{
		return false
	}
}

// 获取网关信息
func GetGatewayInfo(gatewayHashKey string,userID int) (bool,interface{}){
	flag,gatewayID := dao.GetIDFromHashKey(gatewayHashKey)
	if flag{
		if dao.CheckGatewayPermission(gatewayID,userID){
			return dao.GetGatewayInfo(gatewayHashKey)
		}else{
			return false,nil
		}
	}else{
		return false,nil
	}
}

// 获取网关列表
func GetGatewayList(userID int) (bool,[]*utils.GatewayInfo){
	return dao.GetGatewayList(userID)
}
/**
 * 判断用户是否拥有对网关的操作权限
 * @param $hash_key 网关的hash_key
 * @param $user_id 用户的数字ID
 */
func CheckGatewayPermission(gatewayHashKey string,userID int) bool{
	flag,gatewayID := dao.GetIDFromHashKey(gatewayHashKey)
	
	if flag{
		return dao.CheckGatewayPermission(gatewayID,userID)
	}else{
		return false
	}
}

// 从hashKey获取ID
func GetIDFromHashKey(gatewayHashKey string) (bool,int){
	return dao.GetIDFromHashKey(gatewayHashKey)
}

// 通过hashKey获取网关别名
func GetGatewayAlias(gatewayHashKey string) (bool,string){
	return dao.GetGatewayAlias(gatewayHashKey)
}

// 查询网关别名是否存在
func CheckGatewayAliasIsExist(gatewayAlias string) (bool,string){
	return dao.CheckGatewayAliasIsExist(gatewayAlias)
}
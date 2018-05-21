package dao

import (
	"goku-ce/server/conf"
	"goku-ce/utils"
)
 
// 用户登录
func Login(loginName, loginPassword string) (bool) {
	if conf.GlobalConf.LoginName == loginName && conf.GlobalConf.LoginPassword == loginPassword {
		return true
	} else {
		return false
	}
}

// 检查用户是否登录
func CheckLogin(userToken, loginName string) (bool) {
	if utils.Md5(conf.GlobalConf.LoginName + conf.GlobalConf.LoginPassword) == userToken {
		return true
	} else {
		return false
	}
}
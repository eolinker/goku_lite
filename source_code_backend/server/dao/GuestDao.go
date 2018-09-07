package dao

import (
	"fmt"
	"goku-ce/server/conf"
	"goku-ce/utils"
)

// 用户登录
func Login(loginName, loginPassword string) bool {
	fmt.Println("conf LoginName", conf.GlobalConf.LoginName)
	fmt.Println("LoginName", loginName)
	fmt.Println("conf LoginPassword", conf.GlobalConf.LoginPassword)
	fmt.Println("LoginPassword", loginPassword)
	if conf.GlobalConf.LoginName == loginName && conf.GlobalConf.LoginPassword == loginPassword {
		return true
	} else {
		return false
	}
}

// 检查用户是否登录
func CheckLogin(userToken, loginName string) bool {
	if utils.Md5(conf.GlobalConf.LoginName+conf.GlobalConf.LoginPassword) == userToken {
		return true
	} else {
		return false
	}
}

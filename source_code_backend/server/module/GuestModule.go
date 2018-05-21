package module

import (
	"goku-ce/server/dao"
)

// 用户登录
func Login(loginName, loginPassword string) bool {
	return dao.Login(loginName,loginPassword)
}

// 检查用户登录
func CheckLogin(userToken, loginName string) bool {
	return dao.CheckLogin(userToken,loginName)
} 
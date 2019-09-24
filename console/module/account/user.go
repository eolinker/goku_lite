package account

import (
	"github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

// 获取具有编辑权限的用户列表
func GetUserListWithPermission(operationType, operation string) (bool, []map[string]interface{}, error) {
	return console_mysql.GetUserListWithPermission(operationType, operation)
}

// 修改账户信息
func EditPassword(oldPassword, newPassword string, userID int) (bool, string, error) {
	return console_mysql.EditPassword(oldPassword, newPassword, userID)
}

// 获取账户信息
func GetUserInfo(userID int) (bool, interface{}, error) {
	return console_mysql.GetUserInfo(userID)
}

// 获取用户类型
func GetUserType(userID int) (bool, interface{}, error) {
	return console_mysql.GetUserType(userID)
}

// 判断是否是管理员
func CheckUserIsAdmin(userID int) (bool, string, error) {
	return console_mysql.CheckUserIsAdmin(userID)
}

// 判断是否是超级管理员
func CheckUserIsSuperAdmin(userID int) (bool, string, error) {
	return console_mysql.CheckUserIsSuperAdmin(userID)
}

// 检查用户权限
func CheckUserPermission(operationType, operation string, userID int) (bool, string, error) {
	return console_mysql.CheckUserPermission(operationType, operation, userID)
}

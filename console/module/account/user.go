package account

import (
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
)

//GetUserListWithPermission 获取具有编辑权限的用户列表
func GetUserListWithPermission(operationType, operation string) (bool, []map[string]interface{}, error) {
	return console_sqlite3.GetUserListWithPermission(operationType, operation)
}

//EditPassword 修改账户信息
func EditPassword(oldPassword, newPassword string, userID int) (bool, string, error) {
	return console_sqlite3.EditPassword(oldPassword, newPassword, userID)
}

//GetUserInfo 获取账户信息
func GetUserInfo(userID int) (bool, interface{}, error) {
	return console_sqlite3.GetUserInfo(userID)
}

//GetUserType 获取用户类型
func GetUserType(userID int) (bool, interface{}, error) {
	return console_sqlite3.GetUserType(userID)
}

//CheckUserIsAdmin 判断是否是管理员
func CheckUserIsAdmin(userID int) (bool, string, error) {
	return console_sqlite3.CheckUserIsAdmin(userID)
}

//CheckUserIsSuperAdmin 判断是否是超级管理员
func CheckUserIsSuperAdmin(userID int) (bool, string, error) {
	return console_sqlite3.CheckUserIsSuperAdmin(userID)
}

//CheckUserPermission 检查用户权限
func CheckUserPermission(operationType, operation string, userID int) (bool, string, error) {
	return console_sqlite3.CheckUserPermission(operationType, operation, userID)
}

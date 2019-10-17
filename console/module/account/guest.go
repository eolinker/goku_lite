package account

import (
	console_sqlite3 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3"
)

//Login 登录
func Login(loginCall, loginPassword string) (bool, int) {
	return console_sqlite3.Login(loginCall, loginPassword)

}

//CheckLogin 检查用户是否登录
func CheckLogin(userToken string, userID int) bool {
	return console_sqlite3.CheckLogin(userToken, userID)
}

//Register 用户注册
func Register(loginCall, loginPassword string) bool {
	return console_sqlite3.Register(loginCall, loginPassword)
}

//CheckSuperAdminCount 获取超级管理员数量
func CheckSuperAdminCount() (int, error) {
	b, err := console_sqlite3.CheckSuperAdminCount()
	return b, err
}

package account

import (
	"github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

//Login 登录
func Login(loginCall, loginPassword string) (bool, int) {
	return consolemysql.Login(loginCall, loginPassword)

}

//CheckLogin 检查用户是否登录
func CheckLogin(userToken string, userID int) bool {
	return consolemysql.CheckLogin(userToken, userID)
}

//Register 用户注册
func Register(loginCall, loginPassword string) bool {
	return consolemysql.Register(loginCall, loginPassword)
}

//CheckSuperAdminCount 获取超级管理员数量
func CheckSuperAdminCount() (int, error) {
	b, err := consolemysql.CheckSuperAdminCount()
	return b, err
}

package account

import (
	"github.com/eolinker/goku/server/dao/console-mysql"
)

func Login(loginCall, loginPassword string) (bool, int) {
	return console_mysql.Login(loginCall, loginPassword)

}

// 检查用户是否登录
func CheckLogin(userToken string, userID int) bool {
	return console_mysql.CheckLogin(userToken, userID)
}

// 用户注册
func Register(loginCall, loginPassword string) bool {
	return console_mysql.Register(loginCall, loginPassword)
}

func CheckSuperAdminCount()(int,error){
	b,err:=console_mysql.CheckSuperAdminCount()
	return b,err
}
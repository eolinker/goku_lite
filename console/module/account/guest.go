package account

import (
	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"

)

var(
	guestDao dao.GuestDao
	userDao dao.UserDao
)

func init() {
	pdao.Need(&guestDao,&userDao)
}
 

//Login 登录
func  Login(loginCall, loginPassword string) (bool, int) {
	return guestDao.Login(loginCall, loginPassword)

}

//CheckLogin 检查用户是否登录
func  CheckLogin(userToken string, userID int) bool {

	return  guestDao.CheckLogin(userToken, userID)
}

//Register 用户注册
func  Register(loginCall, loginPassword string) bool {
	return  guestDao.Register(loginCall, loginPassword)
}

//CheckSuperAdminCount 获取超级管理员数量
func  CheckSuperAdminCount() (int, error) {
	b, err :=  userDao.CheckSuperAdminCount()
	return b, err
}



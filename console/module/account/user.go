package account

//EditPassword 修改账户信息
func EditPassword(oldPassword, newPassword string, userID int) (bool, string, error) {
	return userDao.EditPassword(oldPassword, newPassword, userID)
}

//GetUserInfo 获取账户信息
func GetUserInfo(userID int) (bool, interface{}, error) {
	return userDao.GetUserInfo(userID)
}

//GetUserType 获取用户类型
func GetUserType(userID int) (bool, interface{}, error) {
	return userDao.GetUserType(userID)
}

//CheckUserIsAdmin 判断是否是管理员
func CheckUserIsAdmin(userID int) (bool, string, error) {
	return userDao.CheckUserIsAdmin(userID)
}

//CheckUserIsSuperAdmin 判断是否是超级管理员
func CheckUserIsSuperAdmin(userID int) (bool, string, error) {
	return userDao.CheckUserIsSuperAdmin(userID)
}

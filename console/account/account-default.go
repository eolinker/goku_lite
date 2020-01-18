package account

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/module/account"
)

//DefaultAccount default
type DefaultAccount struct {
}

//NewDefaultAccount new defaultAccount
func NewDefaultAccount() *DefaultAccount {
	return &DefaultAccount{}
}

//CheckLogin 判断是否登录
func (d *DefaultAccount) CheckLogin(r *http.Request) (int, error) {

	userIDCookie, idErr := r.Cookie("userID")
	userCookie, userErr := r.Cookie("userToken")
	if idErr != nil || userErr != nil {
		e := errors.New("user not logged in")
		return 0, e
	}
	userID, err := strconv.Atoi(userIDCookie.Value)
	if err != nil {
		return 0, err
	}
	flag := account.CheckLogin(userCookie.Value, userID)
	if !flag {
		e := errors.New("illegal users")
		return userID, e
	}

	return userID, nil
}

//CheckPermission 检查操作权限
func (d *DefaultAccount) CheckPermission(pre string, isEdit bool, userID int) (bool, error) {
	if isEdit {
		return true, nil
	}
	return true, nil
}

package account

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/account"
	"github.com/eolinker/goku-api-gateway/utils"
)

//Logout 用户注销
func Logout(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	userIDCookie := http.Cookie{Name: "userID", Path: "/", MaxAge: -1}
	userCookie := http.Cookie{Name: "userToken", Path: "/", MaxAge: -1}
	http.SetCookie(httpResponse, &userIDCookie)
	http.SetCookie(httpResponse, &userCookie)

	controller.WriteResultInfo(httpResponse, "user", "", nil)
	return
}

//EditPassword 修改账户信息
func EditPassword(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	userID, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationEDIT)
	if e != nil {
		return
	}

	oldPassword := httpRequest.PostFormValue("oldPassword")
	newPassword := httpRequest.PostFormValue("newPassword")
	if flag, _ := regexp.MatchString("^[0-9a-zA-Z]{32}$", oldPassword); !flag {

		controller.WriteError(httpResponse,
			"110005",
			"user",
			"[error]illegal oldPassword!",
			errors.New("[error]illegal oldPassword"))
		return
	}
	if flag, _ := regexp.MatchString("^[0-9a-zA-Z]{32}$", newPassword); !flag {

		controller.WriteError(httpResponse,
			"110006",
			"user",
			"[error]illegal newPassword!",
			errors.New("[error]illegal newPassword"))
		return
	}
	flag, result, err := account.EditPassword(oldPassword, newPassword, userID)
	if !flag {
		controller.WriteError(httpResponse,
			"120000",
			"user",
			result,
			err)
		return
	}

	userCookie := &http.Cookie{Name: "userToken", Value: utils.Md5(result + utils.Md5(newPassword)), Path: "/", MaxAge: 86400}
	nameCookie := &http.Cookie{Name: "userID", Value: strconv.Itoa(userID), Path: "/", MaxAge: 86400}
	http.SetCookie(httpResponse, userCookie)
	http.SetCookie(httpResponse, nameCookie)

	controller.WriteResultInfo(httpResponse, "user", "", nil)

	return
}

//GetUserInfo 获取用户信息
func GetUserInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	userID, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationEDIT)
	if e != nil {
		return
	}

	flag, result, err := account.GetUserInfo(userID)
	if !flag {

		controller.WriteError(httpResponse, "110000", "user", result.(string), err)
		return
	}

	controller.WriteResultInfo(httpResponse, "user", "userInfo", result)

	return
}

//GetUserType 获取用户类型
func GetUserType(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	userID, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationEDIT)
	if e != nil {
		return
	}

	flag, result, err := account.GetUserType(userID)
	if !flag {

		controller.WriteError(httpResponse, "110000", "user", result.(string), err)
		return
	}
	controller.WriteResultInfo(httpResponse, "user", "userType", result)

	return
}

//CheckUserIsAdmin 判断是否是管理员
func CheckUserIsAdmin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	userID, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationEDIT)
	if e != nil {
		return
	}

	flag, _, err := account.CheckUserIsAdmin(userID)
	if !flag {

		controller.WriteError(httpResponse,
			"110000",
			"user",
			"This is not administrator",
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "user", "", nil)
	return
}

//CheckUserIsSuperAdmin 判断是否是超级管理员
func CheckUserIsSuperAdmin(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	userID, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationEDIT)
	if e != nil {
		return
	}

	flag, _, err := account.CheckUserIsSuperAdmin(userID)

	if !flag {

		controller.WriteError(httpResponse,
			"110000",
			"user",
			"This is not administrator",
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "user", "", nil)
	return
}

//CheckUserPermission 检查用户权限
func CheckUserPermission(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	userID, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationEDIT)
	if e != nil {
		return
	}

	operationType := httpRequest.PostFormValue("operationType")
	operation := httpRequest.PostFormValue("operation")
	flag, result, err := account.CheckUserPermission(operationType, operation, userID)

	if !flag {

		controller.WriteError(httpResponse,
			"110000",
			"user",
			result,
			err)
		return

	}

	controller.WriteResultInfo(httpResponse, "user", "", nil)
	return
}

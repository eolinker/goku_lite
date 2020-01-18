package account

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/account"
	"github.com/eolinker/goku-api-gateway/utils"
)

//OperationUser 用户权限
const OperationUser = "user"

//UserController 用户控制器
type UserController struct {
}

//NewUserController 新建用户控制器
func NewUserController() *UserController {
	return &UserController{}
}

//Handlers 处理类
func (u *UserController) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {

	return map[string]http.Handler{
		"/logout":            factory.NewAccountHandleFunction(OperationUser, false, Logout),
		"/password/edit":     factory.NewAccountHandleFunction(OperationUser, false, EditPassword),
		"/getInfo":           factory.NewAccountHandleFunction(OperationUser, false, GetUserInfo),
		"/getUserType":       factory.NewAccountHandleFunction(OperationUser, false, GetUserType),
		"/checkIsAdmin":      factory.NewAccountHandleFunction(OperationUser, false, CheckUserIsAdmin),
		"/checkIsSuperAdmin": factory.NewAccountHandleFunction(OperationUser, false, CheckUserIsSuperAdmin),
	}
}

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

	oldPassword := httpRequest.PostFormValue("oldPassword")
	newPassword := httpRequest.PostFormValue("newPassword")
	if flag, _ := regexp.MatchString("^[0-9a-zA-Z]{32}$", oldPassword); !flag {

		controller.WriteError(httpResponse,
			"110005",
			"user",
			"[ERROR]Illegal oldPassword!",
			errors.New("[ERROR]Illegal oldPassword"))
		return
	}
	if flag, _ := regexp.MatchString("^[0-9a-zA-Z]{32}$", newPassword); !flag {

		controller.WriteError(httpResponse,
			"110006",
			"user",
			"[ERROR]Illegal newPassword!",
			errors.New("[ERROR]Illegal newPassword"))
		return
	}
	userID := goku_handler.UserIDFromRequest(httpRequest)
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
	userID := goku_handler.UserIDFromRequest(httpRequest)

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

	userID := goku_handler.UserIDFromRequest(httpRequest)
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
	userID := goku_handler.UserIDFromRequest(httpRequest)

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
	userID := goku_handler.UserIDFromRequest(httpRequest)

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

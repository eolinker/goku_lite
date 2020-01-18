package account

import (
	"errors"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"net/http"
	"strconv"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/account"
	"github.com/eolinker/goku-api-gateway/utils"
)

//Account 账号类
type Account struct {
}

//Handlers 处理器
func (c *Account) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/login": http.HandlerFunc(Login),
	}
}

//NewAccountController 新建账号控制类
func NewAccountController() *Account {
	return &Account{}
}

//Login 用户登录
func Login(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	loginCall := httpRequest.PostFormValue("loginCall")
	loginPassword := httpRequest.PostFormValue("loginPassword")

	loginPassword = utils.Md5(loginPassword)
	flag, userID := account.Login(loginCall, loginPassword)
	if !flag {

		controller.WriteError(httpResponse,
			"100000",
			"guest",
			"[ERROR]Wrong username or password!",
			errors.New("wrong username or password"))
		return
	}

	userCookie := &http.Cookie{Name: "userToken", Value: utils.Md5(loginCall + loginPassword), Path: "/", MaxAge: 86400}
	nameCookie := &http.Cookie{Name: "userID", Value: strconv.Itoa(userID), Path: "/", MaxAge: 86400}
	http.SetCookie(httpResponse, userCookie)
	http.SetCookie(httpResponse, nameCookie)

	controller.WriteResultInfo(httpResponse, "guest", "userID", userID)
	return
}

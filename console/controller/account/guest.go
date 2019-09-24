package account

import (
	"errors"
	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/account"
	"github.com/eolinker/goku-api-gateway/utils"
	"net/http"
	"strconv"
)

// 用户登录
func Login(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//resultInfo := entity2.ResultInfo{}

	loginCall := httpRequest.PostFormValue("loginCall")
	loginPassword := httpRequest.PostFormValue("loginPassword")

	loginPassword = utils.Md5(loginPassword)
	flag, userID := account.Login(loginCall, loginPassword)
	if !flag {

		controller.WriteError(httpResponse,
			"100000",
			"guest",
			"[ERROR]Wrong username or password!",
			errors.New("Wrong username or password"))
		return
	}

	userCookie := &http.Cookie{Name: "userToken", Value: utils.Md5(loginCall + loginPassword), Path: "/", MaxAge: 86400}
	nameCookie := &http.Cookie{Name: "userID", Value: strconv.Itoa(userID), Path: "/", MaxAge: 86400}
	http.SetCookie(httpResponse, userCookie)
	http.SetCookie(httpResponse, nameCookie)

	controller.WriteResultInfo(httpResponse, "guest", "userID", userID)
	return
}

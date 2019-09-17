package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/eolinker/goku/console/controller"
	"github.com/eolinker/goku/console/module/auth"
	log "github.com/eolinker/goku/goku-log"
)

// 获取认证状态
func GetAuthStatus(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}

	strategyID := httpRequest.PostFormValue("strategyID")

	flag, result, err := auth.GetAuthStatus(strategyID)
	if !flag {

		controller.WriteError(httpResponse, "250000", "auth", "[ERROR]The auth info of the strategy does not exist!", err)
		return
	}

	result["statusCode"] = "000000"
	result["type"] = "auth"
	res, _ := json.Marshal(result)

	httpResponse.Write(res)

	return
}

// 获取认证信息
func GetAuthInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationREAD)
	if e != nil {
		return
	}

	strategyID := httpRequest.PostFormValue("strategyID")

	flag, result, err := auth.GetAuthInfo(strategyID)
	if !flag {

		log.Debug(err)
		controller.WriteError(httpResponse, "250000", "auth", "[ERROR]The auth info of the strategy does not exist!", err)

		return
	}

	result["statusCode"] = "000000"
	result["type"] = "auth"
	res, _ := json.Marshal(result)

	httpResponse.Write(res)
	return

}

// 编辑认证信息
func EditAuthInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	_, e := controller.CheckLogin(httpResponse, httpRequest, controller.OperationNone, controller.OperationEDIT)
	if e != nil {
		return
	}

	strategyID := httpRequest.PostFormValue("strategyID")
	strategyName := httpRequest.PostFormValue("strategyName")
	basicAuthList := httpRequest.PostFormValue("basicAuthList")
	apikeyList := httpRequest.PostFormValue("apiKeyList")
	jwtCredentialList := httpRequest.PostFormValue("jwtCredentialList")
	oauth2CredentialList := httpRequest.PostFormValue("oauth2CredentialList")
	delClientIDList := httpRequest.PostFormValue("deleteClientIDList")

	idList := strings.Split(delClientIDList, ",")
	flag, err := auth.EditAuthInfo(strategyID, strategyName, basicAuthList, apikeyList, jwtCredentialList, oauth2CredentialList, idList)
	if !flag {
		controller.WriteError(httpResponse, "250000", "auth", "[ERROR]Fail to edit auth!", err)

		return
	}
	controller.WriteResultInfo(httpResponse, "auth", "", nil)
}

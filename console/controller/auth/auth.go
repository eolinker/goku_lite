package auth

import (
	"encoding/json"
	"net/http"
	"strings"

	goku_handler "github.com/eolinker/goku-api-gateway/goku-handler"

	"github.com/eolinker/goku-api-gateway/console/controller"
	"github.com/eolinker/goku-api-gateway/console/module/auth"
	log "github.com/eolinker/goku-api-gateway/goku-log"
)

const operationAuth = "strategyManagement"

//Handlers handlers
type Handlers struct {
}

//Handlers handlers
func (h *Handlers) Handlers(factory *goku_handler.AccountHandlerFactory) map[string]http.Handler {
	return map[string]http.Handler{
		"/getStatus": factory.NewAccountHandleFunction(operationAuth, true, GetAuthStatus),
		"/getInfo":   factory.NewAccountHandleFunction(operationAuth, true, GetAuthInfo),
		"/editInfo":  factory.NewAccountHandleFunction(operationAuth, true, EditAuthInfo),
	}
}

//NewHandlers new handlers
func NewHandlers() *Handlers {
	return &Handlers{}
}

//GetAuthStatus 获取认证状态
func GetAuthStatus(httpResponse http.ResponseWriter, httpRequest *http.Request) {

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

//GetAuthInfo 获取认证信息
func GetAuthInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {

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

//EditAuthInfo 编辑认证信息
func EditAuthInfo(httpResponse http.ResponseWriter, httpRequest *http.Request) {

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

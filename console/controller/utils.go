package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	"github.com/eolinker/goku-api-gateway/console/module/account"
)

const (
	OperationEDIT = "edit"
	OperationREAD = "read"
)
const (
	OperationNone          = ""
	OperationAPI           = "apiManagement"
	OperationADMIN         = "adminManagement"
	OperationLoadBalance   = "loadBalance"
	OperationStrategy      = "strategyManagement"
	OperationNode          = "nodeManagement"
	OperationPlugin        = "pluginManagement"
	OperationGatewayConfig = "gatewayConfig"
	OperationAlert         = "alertManagement"
)

type PageInfo struct {
	ItemNum  int `json:"itemNum,"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
	TotalNum int `json:"totalNum,"`
}

func (p *PageInfo) SetPage(page, size, total int) *PageInfo {
	p.Page = page
	p.PageSize = size
	p.TotalNum = total
	return p
}
func NewItemNum(num int) *PageInfo {
	return &PageInfo{
		ItemNum: num,
	}
}
func WriteResult(w http.ResponseWriter, code string, resultType, resultKey string, result interface{}, pageInfo *PageInfo) {
	ret := map[string]interface{}{
		"statusCode": code,
	}
	if resultType != "" {
		ret["type"] = resultType
	}
	if result != nil {
		ret[resultKey] = result
	}
	if pageInfo != nil {
		ret["page"] = pageInfo
	}

	data, err := json.Marshal(ret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithFields(ret).Debug(err)
		return
	}

	i, err := w.Write(data)
	if err != nil {
		log.WithFields(ret).Debug("write error:", err)
	} else {
		log.WithFields(ret).Debug("write :", i)
	}
}
func WriteResultInfoWithPage(w http.ResponseWriter, resultType, resultKey string, result interface{}, pageInfo *PageInfo) {
	WriteResult(w, "000000", resultType, resultKey, result, pageInfo)
}
func WriteResultInfo(w http.ResponseWriter, resultType string, resultKey string, result interface{}) {

	if result != nil {
		if t := reflect.TypeOf(result); t.Kind() == reflect.Slice {

			WriteResultInfoWithPage(w, resultType, resultKey, result, NewItemNum(reflect.ValueOf(result).Len()))
			return
		}
	}
	WriteResultInfoWithPage(w, resultType, resultKey, result, nil)
}
func WriteResultInfoWithCode(w http.ResponseWriter, code string, resultType, resultKey string, result interface{}) {
	WriteResult(w, code, resultType, resultKey, result, nil)
}
func WriteError(w http.ResponseWriter, statusCode, resultType, resultDesc string, resuleErr error) {

	ret := map[string]interface{}{
		"type":       resultType,
		"statusCode": statusCode,
		"resultDesc": resultDesc,
	}

	if resuleErr != nil {
		log.Info(resuleErr)
	}

	data, err := json.Marshal(ret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithFields(ret).Debug("write error:", err)
	}

	w.Write(data)
	log.WithFields(ret).Debug("write error:", err)

}

func CheckLogin(w http.ResponseWriter, r *http.Request, operationType, operation string) (int, error) {

	userIDCookie, idErr := r.Cookie("userID")
	userCookie, userErr := r.Cookie("userToken")
	if idErr != nil || userErr != nil {
		e := errors.New("User not logged in!")
		WriteError(w, "100001", "user", e.Error(), e)
		return 0, e
	}
	userID, err := strconv.Atoi(userIDCookie.Value)
	if err != nil {
		WriteError(w, "100001", "user", "Illegal user Id!", err)
		return 0, err
	}
	flag := account.CheckLogin(userCookie.Value, userID)
	if !flag {
		e := errors.New("Illegal users!")
		WriteError(w, "100001", "user", "Illegal users!", e)
		return userID, e
	}
	if operation == OperationEDIT && OperationNone != operationType {
		if operationType == OperationADMIN {

			flag, desc, err := account.CheckUserIsAdmin(userID)
			if !flag {
				WriteError(w, "100002", "user", desc, err)
				return userID, errors.New(desc)
			}
		} else {
			flag, desc, err := account.CheckUserPermission(operationType, "edit", userID)
			if !flag {

				WriteError(w, "100002", "user", desc, err)
				return userID, errors.New(desc)
			}
		}

	}

	return userID, nil
}

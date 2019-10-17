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
	//OperationEDIT 编辑
	OperationEDIT = "edit"
	//OperationREAD 读取
	OperationREAD = "read"
)
const (
	//OperationNone 无操作
	OperationNone = ""
	//OperationAPI 接口操作
	OperationAPI = "apiManagement"
	//OperationADMIN 管理员
	OperationADMIN = "adminManagement"
	//OperationLoadBalance 负载操作
	OperationLoadBalance = "loadBalance"
	//OperationStrategy 策略操作
	OperationStrategy = "strategyManagement"
	//OperationNode 节点操作
	OperationNode = "nodeManagement"
	//OperationPlugin 插件操作
	OperationPlugin = "pluginManagement"
	//OperationGatewayConfig 网关配置操作
	OperationGatewayConfig = "gatewayConfig"
)

//PageInfo 页码信息
type PageInfo struct {
	ItemNum  int `json:"itemNum,"`
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pageSize,omitempty"`
	TotalNum int `json:"totalNum,"`
}

//SetPage 设置页码
func (p *PageInfo) SetPage(page, size, total int) *PageInfo {
	p.Page = page
	p.PageSize = size
	p.TotalNum = total
	return p
}

//NewItemNum 创建新的item
func NewItemNum(num int) *PageInfo {
	return &PageInfo{
		ItemNum: num,
	}
}

//WriteResult 写响应
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

//WriteResultInfoWithPage 返回带页码信息的数据
func WriteResultInfoWithPage(w http.ResponseWriter, resultType, resultKey string, result interface{}, pageInfo *PageInfo) {
	WriteResult(w, "000000", resultType, resultKey, result, pageInfo)
}

//WriteResultInfo 返回信息
func WriteResultInfo(w http.ResponseWriter, resultType string, resultKey string, result interface{}) {

	if result != nil {
		if t := reflect.TypeOf(result); t.Kind() == reflect.Slice {

			WriteResultInfoWithPage(w, resultType, resultKey, result, NewItemNum(reflect.ValueOf(result).Len()))
			return
		}
	}
	WriteResultInfoWithPage(w, resultType, resultKey, result, nil)
}

//WriteResultInfoWithCode 返回带状态信息的响应
func WriteResultInfoWithCode(w http.ResponseWriter, code string, resultType, resultKey string, result interface{}) {
	WriteResult(w, code, resultType, resultKey, result, nil)
}

//WriteError 返回错误
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

//CheckLogin 判断是否登录
func CheckLogin(w http.ResponseWriter, r *http.Request, operationType, operation string) (int, error) {

	userIDCookie, idErr := r.Cookie("userID")
	userCookie, userErr := r.Cookie("userToken")
	if idErr != nil || userErr != nil {
		e := errors.New("user not logged in")
		WriteError(w, "100001", "user", e.Error(), e)
		return 0, e
	}
	userID, err := strconv.Atoi(userIDCookie.Value)
	if err != nil {
		WriteError(w, "100001", "user", "Illegal user id!", err)
		return 0, err
	}
	flag := account.CheckLogin(userCookie.Value, userID)
	if !flag {
		e := errors.New("illegal users")
		WriteError(w, "100001", "user", "Illegal users!", e)
		return userID, e
	}

	return userID, nil
}

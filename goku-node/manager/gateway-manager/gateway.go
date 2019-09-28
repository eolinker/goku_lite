package gatewaymanager

import (
	"strconv"
	"strings"
	"sync"

	dao_gateway "github.com/eolinker/goku-api-gateway/server/dao/node-mysql/dao-gateway"
)

var (
	locker        = sync.RWMutex{}
	_SuccessCode  = "200"
	_UpdatePeriod = 5
	_AlertStatus  = 0
	_AlertInfo    = "{\"alertAddr\":\"\",\"alertPeriodType\":0,\"logPath\":\"./log/apiAlert\",\"receiverList\":\"\"}"
)

//GetUpdatePeriod 获取更新周期
func GetUpdatePeriod() int {
	locker.RLock()
	defer locker.RUnlock()
	return _UpdatePeriod
}

//GetAlertStatus 获取告警状态
func GetAlertStatus() int {
	locker.RLock()
	defer locker.RUnlock()
	return _AlertStatus
}

//GetAlertInfo 获取告警信息
func GetAlertInfo() string {
	locker.RLock()
	defer locker.RUnlock()
	return _AlertInfo
}

//LoadGatewayConfig 加载网关配置
func LoadGatewayConfig() {
	code, period := dao_gateway.GetGatewayBaseInfo()
	alertInfo, alrtStatus := dao_gateway.GetGatewayAlertInfo()
	locker.Lock()
	defer locker.Unlock()
	_SuccessCode = code
	_UpdatePeriod = period
	_AlertStatus = alrtStatus
	_AlertInfo = alertInfo
}

//IsSucess 是否成功
func IsSucess(statusCode int) bool {
	locker.RLock()
	defer locker.RUnlock()
	return strings.Contains(_SuccessCode, strconv.Itoa(statusCode))
}

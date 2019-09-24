package dao_gateway

import "github.com/eolinker/goku-api-gateway/common/database"

// 获取网关成功状态码
func GetGatewayBaseInfo() (string, int) {
	db := database.GetConnection()
	var successCode string
	var updatePeriod int
	sql := "SELECT successCode,nodeUpdatePeriod FROM goku_gateway WHERE id = 1;"
	err := db.QueryRow(sql).Scan(&successCode, &updatePeriod)
	if err != nil {
		return "200", 5
	}
	return successCode, updatePeriod
}

// 获取节点告警信息
func GetGatewayAlertInfo() (string, int) {
	db := database.GetConnection()
	var apiAlertInfo string
	var alertStatus int
	sql := "SELECT alertStatus,apiAlertInfo FROM goku_gateway WHERE id = 1;"
	err := db.QueryRow(sql).Scan(&alertStatus, &apiAlertInfo)
	if err != nil {
		return "{\"alertAddr\":\"\",\"alertPeriodType\":0,\"logPath\":\"./log/apiAlert\",\"receiverList\":\"\"}", 0
	}
	return apiAlertInfo, alertStatus
}

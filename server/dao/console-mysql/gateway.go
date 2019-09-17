package console_mysql

import (
	database2 "github.com/eolinker/goku/common/database"
)

func GetGatewayConfig() (map[string]interface{}, error) {
	db := database2.GetConnection()
	var successCode string
	var nodeUpdatePeriod, monitorUpdatePeriod, monitorTimeout int
	sql := `SELECT successCode,nodeUpdatePeriod,monitorUpdatePeriod,monitorTimeout FROM goku_gateway WHERE id = 1;`
	err := db.QueryRow(sql).Scan(&successCode, &nodeUpdatePeriod, &monitorUpdatePeriod, &monitorTimeout)
	if err != nil {
		return nil, err
	}
	gatewayConfig := map[string]interface{}{
		"successCode":         successCode,
		"nodeUpdatePeriod":    nodeUpdatePeriod,
		"monitorUpdatePeriod": monitorUpdatePeriod,
		"monitorTimeout":      monitorTimeout,
	}
	return gatewayConfig, nil
}

// 编辑网关基本配置
func EditGatewayBaseConfig(successCode string, nodeUpdatePeriod, monitorUpdatePeriod, monitorTimeout int) (bool, string, error) {
	db := database2.GetConnection()
	sql := "SELECT successCode FROM goku_gateway WHERE id = 1;"
	code := ""
	err := db.QueryRow(sql).Scan(&code)
	if err != nil {
		sql = "INSERT INTO goku_gateway (id,successCode,nodeUpdatePeriod,monitorUpdatePeriod,monitorTimeout) VALUES (1,?,?,?,?)"
	} else {
		sql = "UPDATE goku_gateway SET successCode = ?,nodeUpdatePeriod = ?,monitorUpdatePeriod = ?,monitorTimeout = ? WHERE id = 1;"
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, "[ERROR]Illegal SQL Statement!", err
	}
	_, err = stmt.Exec(successCode, nodeUpdatePeriod, monitorUpdatePeriod, monitorTimeout)
	if err != nil {
		return false, "[ERROR]Fail to excute SQL Statement!", err
	}
	return true, "", nil
}

// 编辑网关告警配置
func EditGatewayAlarmConfig(apiAlertInfo, sender, senderPassword, smtpAddress string, alertStatus, smtpPort, smtpProtocol int) (bool, string, error) {
	db := database2.GetConnection()
	sql := "SELECT successCode FROM goku_gateway WHERE id = 1;"
	code := ""
	err := db.QueryRow(sql).Scan(&code)
	if err != nil {
		sql = "INSERT INTO goku_gateway (id,apiAlertInfo,alertStatus,sender,senderPassword,smtpAddress,smtpPort,smtpProtocol) VALUES (1,?,?,?,?,?,?,?)"
	} else {
		sql = "UPDATE goku_gateway SET apiAlertInfo = ?,alertStatus = ?,sender = ?,senderPassword = ?,smtpAddress = ?,smtpPort = ?,smtpProtocol = ? WHERE id = 1;"
	}
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, "[ERROR]Illegal SQL Statement!", err
	}
	_, err = stmt.Exec(apiAlertInfo, alertStatus, sender, senderPassword, smtpAddress, smtpPort, smtpProtocol)
	if err != nil {
		return false, "[ERROR]Fail to excute SQL Statement!", err
	}
	return true, "", nil
}

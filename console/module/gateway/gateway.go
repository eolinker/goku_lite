package gateway

import (
	"github.com/eolinker/goku-api-gateway/server/dao"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
)

func GetGatewayConfig() (map[string]interface{}, error) {
	return console_mysql.GetGatewayConfig()
}

// 编辑网关基本配置
func EditGatewayBaseConfig(successCode string, nodeUpdatePeriod, monitorUpdatePeriod, timeout int) (bool, string, error) {
	tableName := "goku_gateway"
	flag, result, err := console_mysql.EditGatewayBaseConfig(successCode, nodeUpdatePeriod, monitorUpdatePeriod, timeout)
	if flag {
		dao.UpdateTable(tableName)
	}
	return flag, result, err
}

// 编辑网关告警配置
func EditGatewayAlarmConfig(apiAlertInfo, sender, senderPassword, smtpAddress string, alertStatus, smtpPort, smtpProtocol int) (bool, string, error) {
	tableName := "goku_gateway"
	flag, result, err := console_mysql.EditGatewayAlarmConfig(apiAlertInfo, sender, senderPassword, smtpAddress, alertStatus, smtpPort, smtpProtocol)
	if flag {
		dao.UpdateTable(tableName)
	}
	return flag, result, err
}

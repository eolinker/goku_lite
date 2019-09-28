package consolemysql

import (
	"encoding/json"
	database2 "github.com/eolinker/goku-api-gateway/common/database"
	log "github.com/eolinker/goku-api-gateway/goku-log"
)

type balanceConfig struct {
	LoadBalancingServer []server `json:"loadBalancingServer"`
}

type server struct {
	Server string `json:"server"`
	Weight int    `json:"weight"`
}

// RefreshGatewayAlertConfig 新建项目
func RefreshGatewayAlertConfig() bool {
	db := database2.GetConnection()
	var (
		alertAddr       string
		alertLogPath    string
		alertPeriodType int
		receiverList    string
	)
	// 获取网关告警配置
	sqlCode := `SELECT alertAddress,alertLogPath,alertPeriodType,receiverList FROM goku_gateway;`
	err := db.QueryRow(sqlCode).Scan(&alertAddr, &alertLogPath, &alertPeriodType, &receiverList)
	if err != nil {
		log.Error(err)
		return false
	}
	if alertLogPath == "" {
		alertLogPath = "./log/apiAlert"
	}
	// 构造告警需要的信息
	apiAlertInfo := map[string]interface{}{
		"alertPeriodType": alertPeriodType,
		"receiverList":    receiverList,
		"alertAddr":       alertAddr,
	}
	nodeAlertInfo := map[string]interface{}{
		"receiverList": receiverList,
		"alertAddr":    alertAddr,
	}
	redisAlertInfo := map[string]interface{}{
		"receiverList": receiverList,
		"alertAddr":    alertAddr,
	}
	apiAlertInfoByte, err := json.Marshal(apiAlertInfo)
	if err != nil {
		log.Error(err)
		return false
	}
	nodeAlertInfoByte, err := json.Marshal(nodeAlertInfo)
	if err != nil {
		log.Error(err)
		return false
	}
	redisAlertInfoByte, err := json.Marshal(redisAlertInfo)
	if err != nil {
		log.Error(err)
		return false
	}
	Tx, _ := db.Begin()

	_, err = Tx.Exec("UPDATE goku_gateway SET apiAlertInfo = ?,nodeAlertInfo = ?,redisAlertInfo = ?;", string(apiAlertInfoByte), string(nodeAlertInfoByte), string(redisAlertInfoByte))
	if err != nil {
		Tx.Rollback()
		log.Error(err)
		return false
	}
	dropColomn := []string{"alertPeriodType", "alertAddress", "alertLogPath", "receiverList"}
	for _, colomn := range dropColomn {
		sql := "ALTER TABLE goku_gateway DROP COLUMN " + colomn + ";"
		log.Debug("RefreshGatewayAlertConfig-sql:", sql)
		_, err = Tx.Exec(sql)
		if err != nil {
			Tx.Rollback()
			log.Error(err)
			return false
		}
	}
	Tx.Commit()
	return true
}

type monitorRecord struct {
	gatewayRequestCount   int
	gatewaySuccessCount   int
	gatewayStatus2xxCount int
	gatewayStatus4xxCount int
	gatewayStatus5xxCount int
	proxyRequestCount     int
	proxySuccessCount     int
	proxyStatus2xxCount   int
	proxyStatus4xxCount   int
	proxyStatus5xxCount   int
	proxyTimeoutCount     int
	updateTime            string
	hour                  string
	strategyID            string
	apiID                 int
}

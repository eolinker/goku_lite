package console_mysql

import (
	"encoding/json"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"strconv"
	"strings"
	"time"

	entity2 "github.com/eolinker/goku-api-gateway/server/entity/balance-entity"

	database2 "github.com/eolinker/goku-api-gateway/common/database"
)

type balanceConfig struct {
	LoadBalancingServer []server `json:"loadBalancingServer"`
}

type server struct {
	Server string `json:"server"`
	Weight int    `json:"weight"`
}

func RefreshApiInfo() bool {
	db := database2.GetConnection()
	// 随机生成字符串
	sqlCode := `SELECT apiID,IFNULL(targetServer,"http://") FROM goku_gateway_api;`
	rows, err := db.Query(sqlCode)
	if err != nil {
		log.Error(err)
		return false
	}
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
		log.Error(err)
		return false
	} else {
		Tx, _ := db.Begin()
		for rows.Next() {
			var apiID int
			var targetServer string
			err = rows.Scan(&apiID, &targetServer)
			if err != nil {
				Tx.Rollback()
				log.Error(err)
				return false
			}
			protocol := "http"
			balanceName := ""
			if len(targetServer) > 4 {
				arr := strings.Split(targetServer, "://")
				arrLen := len(arr)
				if arrLen > 1 {
					protocol, balanceName = arr[0], arr[1]
				} else {
					if arrLen == 1 {
						if arr[0] == "http" || arr[0] == "https" {
							protocol = arr[0]
						}
					}
				}

			}
			stripSlash := true
			_, err = Tx.Exec("UPDATE goku_gateway_api SET stripSlash = ?,protocol=?,balanceName=? WHERE apiID = ?", stripSlash, protocol, balanceName, apiID)
			if err != nil {
				Tx.Rollback()
				log.Error(err)
				return false
			}
		}

		Tx.Commit()
		return true
	}
}

// 新建项目
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

// 新建项目
func RefreshBalance(name string) bool {
	db := database2.GetConnection()
	// 获取网关告警配置
	sqlCode := "SELECT balanceID,IFNULL(balanceConfig,'') FROM goku_balance;"
	rows, err := db.Query(sqlCode)
	if err != nil {
		log.Error(err)
		return false
	}
	defer rows.Close()
	Tx, _ := db.Begin()
	for rows.Next() {
		var balanceConfigStr string
		var balanceID int
		err = rows.Scan(&balanceID, &balanceConfigStr)
		if err != nil {
			Tx.Rollback()
			log.Error(err)
			return false
		}
		var configs balanceConfig
		err = json.Unmarshal([]byte(balanceConfigStr), &configs)
		if err != nil {
			Tx.Rollback()
			log.Error(err)
			return false
		}
		staticOrg := ""
		for _, config := range configs.LoadBalancingServer {
			config.Server = strings.Replace(config.Server, " ", "", -1)
			if config.Server == "" {
				continue
			}
			weight := 1
			if config.Weight > 0 {
				weight = config.Weight
			}

			staticOrg += config.Server + " " + strconv.Itoa(weight) + ";"
		}
		defaultConfig := &entity2.BalanceConfig{
			ServersConfigOrg: staticOrg,
		}
		clusterConfig := make(map[string]*entity2.BalanceConfig)
		clusterConfig[name] = &entity2.BalanceConfig{}
		dcStr, _ := json.Marshal(defaultConfig)
		ccStr, _ := json.Marshal(clusterConfig)
		_, err = Tx.Exec("UPDATE goku_balance SET defaultConfig = ?,clusterConfig = ? WHERE balanceID = ?;", dcStr, ccStr, balanceID)
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

func RefreshMonitorRecord(clusterID int) bool {
	db := database2.GetConnection()
	// 获取网关告警配置
	// gatewayRecordSQL := "SELECT gatewayRequestCount,gatewaySuccessCount,gatewayStatus2xxCount,gatewayStatus4xxCount,gatewayStatus5xxCount,gatewayStatus5xxCount,proxySuccessCount,proxyStatus2xxCount,proxyStatus4xxCount,proxyStatus5xxCount,proxyTimeoutCount,updateTime FROM goku_gateway_request_record WHERE gatewayRequestCount > 0;"
	// apiRecordSQL := "SELECT apiID,gatewayRequestCount,gatewaySuccessCount,gatewayStatus2xxCount,gatewayStatus4xxCount,gatewayStatus5xxCount,gatewayStatus5xxCount,proxySuccessCount,proxyStatus2xxCount,proxyStatus4xxCount,proxyStatus5xxCount,proxyTimeoutCount,updateTime FROM goku_gateway_api_request_record WHERE gatewayRequestCount > 0;"
	strategyRecordSQL := "SELECT strategyID,gatewayRequestCount,gatewaySuccessCount,gatewayStatus2xxCount,gatewayStatus4xxCount,gatewayStatus5xxCount,proxyRequestCount,proxySuccessCount,proxyStatus2xxCount,proxyStatus4xxCount,proxyStatus5xxCount,proxyTimeoutCount,updateTime FROM goku_gateway_strategy_request_record WHERE gatewayRequestCount > 0;"
	monitorRecordSQL := "SELECT strategyID,apiID,gatewayRequestCount,gatewaySuccessCount,gatewayStatus2xxCount,gatewayStatus4xxCount,gatewayStatus5xxCount,proxyRequestCount,proxySuccessCount,proxyStatus2xxCount,proxyStatus4xxCount,proxyStatus5xxCount,proxyTimeoutCount,updateTime FROM goku_gateway_monitor_request_record WHERE gatewayRequestCount > 0 AND monitorType = 0;"
	rows, err := db.Query(strategyRecordSQL)
	if err != nil {
		log.Error(err)
		return false
	}
	timeTemplate1 := "2006-01-02 15:04:05"
	timeTemplate2 := "2006010215"
	defer rows.Close()
	records := make([]monitorRecord, 0)
	for rows.Next() {
		record := monitorRecord{}
		err := rows.Scan(&record.strategyID, &record.gatewayRequestCount, &record.gatewaySuccessCount, &record.gatewayStatus2xxCount, &record.gatewayStatus4xxCount, &record.gatewayStatus5xxCount, &record.proxyRequestCount, &record.proxySuccessCount, &record.proxyStatus2xxCount, &record.proxyStatus4xxCount, &record.proxyStatus5xxCount, &record.proxyTimeoutCount, &record.updateTime)
		if err != nil {

			log.Error(err)
			continue
		}
		stamp, _ := time.ParseInLocation(timeTemplate1, record.updateTime, time.Local)
		record.hour = stamp.Format(timeTemplate2)
		record.apiID = 0
		records = append(records, record)
	}

	rows, err = db.Query(monitorRecordSQL)
	if err != nil {
		log.Error(err)
		return false
	}
	for rows.Next() {
		record := monitorRecord{}
		err := rows.Scan(&record.strategyID, &record.apiID, &record.gatewayRequestCount, &record.gatewaySuccessCount, &record.gatewayStatus2xxCount, &record.gatewayStatus4xxCount, &record.gatewayStatus5xxCount, &record.proxyRequestCount, &record.proxySuccessCount, &record.proxyStatus2xxCount, &record.proxyStatus4xxCount, &record.proxyStatus5xxCount, &record.proxyTimeoutCount, &record.updateTime)
		if err != nil {
			log.Error(err)
			continue
		}
		stamp, _ := time.ParseInLocation(timeTemplate1, record.updateTime, time.Local)
		record.hour = stamp.Format(timeTemplate2)
		records = append(records, record)
	}
	Tx, _ := db.Begin()
	refreshSQL := "INSERT INTO goku_monitor_cluster (strategyID,apiID,gatewayRequestCount,gatewaySuccessCount,gatewayStatus2xxCount,gatewayStatus4xxCount,gatewayStatus5xxCount,proxyRequestCount,proxySuccessCount,proxyStatus2xxCount,proxyStatus4xxCount,proxyStatus5xxCount,proxyTimeoutCount,updateTime,clusterID,hour) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
	for _, record := range records {
		Tx.Exec(refreshSQL, record.strategyID, record.apiID, record.gatewayRequestCount, record.gatewaySuccessCount, record.gatewayStatus2xxCount, record.gatewayStatus4xxCount, record.gatewayStatus5xxCount, record.proxyRequestCount, record.proxySuccessCount, record.proxyStatus2xxCount, record.proxyStatus4xxCount, record.proxyStatus5xxCount, record.proxyTimeoutCount, record.updateTime, clusterID, record.hour)
	}
	Tx.Commit()
	return true
}

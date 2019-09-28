package consolemysql

import (
	"encoding/json"
	"strconv"
	"time"

	database2 "github.com/eolinker/goku-api-gateway/common/database"
)

// GetAlertMsgList 获取告警信息列表
func GetAlertMsgList(page, pageSize int) (bool, []map[string]interface{}, int, error) {
	db := database2.GetConnection()
	var count int
	sql := `SELECT COUNT(*) FROM goku_gateway_alert;`
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		return false, make([]map[string]interface{}, 0), 0, err
	}
	sql = "SELECT A.alertID,A.requestURL,A.targetServer,A.targetURL,A.alertPeriodType,A.alertCount,A.updateTime,IFNULL(A.clusterName,''),IFNULL(A.nodeIP,''),IFNULL(B.title,'') FROM goku_gateway_alert A LEFT JOIN goku_cluster B ON B.`name` = A.clusterName ORDER BY updateTime DESC LIMIT ?,?;"
	rows, err := db.Query(sql, (page-1)*pageSize, pageSize)
	if err != nil {
		return false, make([]map[string]interface{}, 0), 0, err
	}
	defer rows.Close()
	alertList := make([]map[string]interface{}, 0)
	for rows.Next() {
		var requestURL, targetServer, targetURL, updateTime, nodeIP, clusterName, title string
		var alertID, alertPeriodType, alertCount int
		err = rows.Scan(&alertID, &requestURL, &targetServer, &targetURL, &alertPeriodType, &alertCount, &updateTime, &clusterName, &nodeIP, &title)
		if err != nil {
			return false, make([]map[string]interface{}, 0), 0, err
		}
		period := "1"
		if alertPeriodType == 1 {
			period = "5"
		} else if alertPeriodType == 2 {
			period = "15"
		} else if alertPeriodType == 3 {
			period = "30"
		} else if alertPeriodType == 4 {
			period = "60"
		}
		var msg  = "网关转发失败" + period + "分钟达到" + strconv.Itoa(alertCount) + "次"
		alertInfo := map[string]interface{}{
			"alertID":      alertID,
			"requestURL":   requestURL,
			"targetServer": targetServer,
			"addr":         nodeIP,
			"name":         clusterName,
			"title":        title,
			"msg":          msg,
			"updateTime":   updateTime,
			"targetURL":    targetURL,
		}
		alertList = append(alertList, alertInfo)
	}
	return true, alertList, count, nil
}

// ClearAlertMsg 清空告警信息列表
func ClearAlertMsg() (bool, string, error) {
	db := database2.GetConnection()
	sql := "DELETE FROM goku_gateway_alert;"
	_, err := db.Exec(sql)
	if err != nil {
		return false, "[ERROR]Illegal SQL Statement!", err
	}
	return true, "", nil
}

// DeleteAlertMsg 删除告警信息
func DeleteAlertMsg(alertID int) (bool, string, error) {
	db := database2.GetConnection()
	sql := "DELETE FROM goku_gateway_alert WHERE alertID = ?;"
	_, err := db.Exec(sql, alertID)
	if err != nil {
		return false, "[ERROR]Illegal SQL Statement!", err
	}
	return true, "", nil
}

// AddAlertMsg 新增告警信息
func AddAlertMsg(requestURL, targetServer, targetURL, ip, clusterName string, alertPeriodType, alertCount int) (bool, string, error) {
	db := database2.GetConnection()
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "INSERT INTO goku_gateway_alert (requestURL,targetServer,targetURL,alertPeriodType,alertCount,nodeIP,clusterName,updateTime) VALUES (?,?,?,?,?,?,?,?);"
	_, err := db.Exec(sql, requestURL, targetServer, targetURL, alertPeriodType, alertCount, ip, clusterName, now)
	if err != nil {
		return false, "[ERROR]Illegal SQL Statement!", err
	}
	return true, "", nil
}

//GetAlertConfig 获取告警配置
func GetAlertConfig() (bool, map[string]interface{}, error) {
	db := database2.GetConnection()
	var apiAlertInfo, sender, senderPassword, smtpAddress string
	var alertStatus, smtpPort, smtpProtocol int
	sql := `SELECT alertStatus,IFNULL(apiAlertInfo,"{}"),IFNULL(sender,""),IFNULL(senderPassword,""),IFNULL(smtpAddress,""),IFNULL(smtpPort,25),IFNULL(smtpProtocol,0) FROM goku_gateway WHERE id = 1;`
	err := db.QueryRow(sql).Scan(&alertStatus, &apiAlertInfo, &sender, &senderPassword, &smtpAddress, &smtpPort, &smtpProtocol)
	if err != nil {
		return false, nil, err
	}

	apiAlertInfoJSON := map[string]interface{}{}

	if apiAlertInfo == "" || apiAlertInfo == "{}" {
		apiAlertInfoJSON["alertPeriodType"] = 0
		apiAlertInfoJSON["alertAddr"] = ""
		apiAlertInfoJSON["receiverList"] = ""
	} else {
		err = json.Unmarshal([]byte(apiAlertInfo), &apiAlertInfoJSON)
		if err != nil {
			return false, nil, err
		}
	}

	gatewayConfig := map[string]interface{}{
		"alertStatus":    alertStatus,
		"sender":         sender,
		"senderPassword": senderPassword,
		"smtpAddress":    smtpAddress,
		"smtpPort":       smtpPort,
		"smtpProtocol":   smtpProtocol,
		"apiAlertInfo":   apiAlertInfoJSON,
	}
	return true, gatewayConfig, nil
}

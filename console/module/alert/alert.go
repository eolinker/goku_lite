package alert

import (
	"strconv"

	config_log "github.com/eolinker/goku-api-gateway/console/module/config-log"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	console_mysql "github.com/eolinker/goku-api-gateway/server/dao/console-mysql"
	"github.com/eolinker/goku-api-gateway/utils"
)

// 获取告警信息列表
func GetAlertMsgList(page, pageSize int) (bool, []map[string]interface{}, int, error) {
	return console_mysql.GetAlertMsgList(page, pageSize)
}

// 清空告警信息列表
func ClearAlertMsg() (bool, string, error) {
	return console_mysql.ClearAlertMsg()
}

// 删除告警信息
func DeleteAlertMsg(alertID int) (bool, string, error) {
	return console_mysql.DeleteAlertMsg(alertID)
}

// 新增告警信息
func AddAlertMsg(apiID, apiName, requestURL, targetServer, targetURL, requestMethod, proxyMethod, headerList, queryParamList, formParamList, responseHeaderList, strategyID, strategyName, requestID string, alertPeriodType, alertCount, responseStatus int, isAlert bool, ip, clusterName string) (bool, string, error) {

	flag, result, err := console_mysql.GetAlertConfig()
	if !flag {
		return false, "", err
	}

	apiAlertInfo := result["apiAlertInfo"].(map[string]interface{})
	AlertLog(requestURL, targetServer, targetURL, requestMethod, proxyMethod, headerList, queryParamList, formParamList, responseHeaderList, responseStatus, ip, strategyID, strategyName, requestID)
	if isAlert {
		// 发送邮件
		log.WithFields(log.Fields(result)).Debug("AddAlertMsg: apiAlertInfo[\"receiverList\"]", apiAlertInfo["receiverList"])

		logConfig, _ := config_log.Get("console")
		go utils.SendAlertMail(result["sender"].(string), result["senderPassword"].(string), result["smtpAddress"].(string), strconv.Itoa(result["smtpPort"].(int)), strconv.Itoa(result["smtpProtocol"].(int)), apiAlertInfo["receiverList"].(string), requestURL, logConfig.Dir, strconv.Itoa(alertPeriodType), strconv.Itoa(alertCount), apiName, apiID, targetServer, targetURL)

		return console_mysql.AddAlertMsg(requestURL, targetServer, targetURL, ip, clusterName, alertPeriodType, alertCount)
	} else {
		return true, "", nil
	}
}

func GetAlertConfig() (bool, map[string]interface{}, error) {
	return console_mysql.GetAlertConfig()
}

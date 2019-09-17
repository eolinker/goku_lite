package utils

import (
	log "github.com/eolinker/goku/goku-log"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
)

func SendAlertMail(sender, senderPassword, smtpAddress, smtpPort, smtpProtocol, receiverMail, requestURL, alertLogPath, alertPeriod, alertCount, apiName, apiID, targetServer, proxyURL string) (bool, error) {
	alertTime := time.Now().Format("2006-01-02 15:04:05")
	f, err := os.Open("html/currentAlert.html")
	if err != nil {
		log.Warn(err)
	}
	body, err := ioutil.ReadAll(f)
	if err != nil {
		log.Warn(err)
	}
	bodyStr := string(body)
	bodyStr = strings.Replace(bodyStr, "$requestURL", requestURL, -1)
	bodyStr = strings.Replace(bodyStr, "$alertTime", alertTime, -1)
	bodyStr = strings.Replace(bodyStr, "$alertLogPath", alertLogPath, -1)
	bodyStr = strings.Replace(bodyStr, "$alertPeriod", period[alertPeriod], -1)
	bodyStr = strings.Replace(bodyStr, "$alertCount", alertCount, -1)
	bodyStr = strings.Replace(bodyStr, "$apiID", apiID, -1)
	bodyStr = strings.Replace(bodyStr, "$targetServer", targetServer, -1)
	bodyStr = strings.Replace(bodyStr, "$proxyURL", proxyURL, -1)
	bodyStr = strings.Replace(bodyStr, "$apiName", apiName, -1)
	host := net.JoinHostPort(smtpAddress, smtpPort)
	subject := "GoKu告警：" + requestURL + "接口在" + period[alertPeriod] + "分钟内转发失败" + alertCount + "次"
	err = SendToMail(sender, senderPassword, host, receiverMail, subject, bodyStr, "html", smtpProtocol)
	if err != nil {
		log.Warn("SendAlertMail:",err)
	}
	return true, nil
}

func SendMonitorAlertMail(sender, senderPassword, smtpAddress, smtpPort, smtpProtocol, receiverMail, bodyStr string) (bool, error) {
	host := net.JoinHostPort(smtpAddress, smtpPort)
	subject := "EOLINKER AGW节点自动重启失败告警"
	err := SendToMail(sender, senderPassword, host, receiverMail, subject, bodyStr, "html", smtpProtocol)
	if err != nil {
		log.Warn("SendMonitorAlertMail:",err)
	}
	return true, nil
}

func ReplaceMonitorBody(nodeList []map[string]string) (bool, string, error) {
	f, err := os.Open("html/monitorAlert.html")
	if err != nil {
		log.Warn(err)
	}
	body, err := ioutil.ReadAll(f)
	if err != nil {
		log.Warn(err)
	}
	bodyStr := string(body)
	nodeStr := ""
	for _, nodeInfo := range nodeList {
		nodeStr += "<p>节点名称：" + nodeInfo["nodeName"] + "</p>"
		nodeStr += "<p>节点IP：" + nodeInfo["nodeIP"] + ":" + nodeInfo["nodePort"] + "</p>"
		nodeStr += "<p></p>"
	}
	bodyStr = strings.Replace(bodyStr, "$alertTime", time.Now().Format("2006-01-02 15:04:05"), -1)
	bodyStr = strings.Replace(bodyStr, "$nodeList", nodeStr, -1)
	return true, bodyStr, nil
}

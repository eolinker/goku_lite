package utils

import (
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"io/ioutil"
	"net"
	"strings"
	"time"
)

var (
	currentAlertBody string
)

func init() {
	body, err := ioutil.ReadFile("html/currentAlert.html")
	if err != nil {
		log.Panic(err)
	}
	currentAlertBody = string(body)
}

func SendAlertMail(sender, senderPassword, smtpAddress, smtpPort, smtpProtocol, receiverMail, requestURL, alertLogPath, alertPeriod, alertCount, apiName, apiID, targetServer, proxyURL string) (bool, error) {
	alertTime := time.Now().Format("2006-01-02 15:04:05")

	bodyStr := currentAlertBody
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
	err := SendToMail(sender, senderPassword, host, receiverMail, subject, bodyStr, "html", smtpProtocol)
	if err != nil {
		log.Warn("SendAlertMail:", err)
		return false, err
	}
	return true, nil
}

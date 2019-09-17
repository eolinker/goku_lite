package cmd

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	node_common "github.com/eolinker/goku/goku-node/node-common"
)
// 新增报警信息
func AddAlertMessage(apiID int, apiName, requestURL, targetServer, targetURL, requestMethod, proxyMethod, headerList, queryParamList, formParamList, responseHeaderList string, alertPeriodType, alertCount, responseStatus int, isAlert string, strategyID string, strategyName, requestID string) (bool, string, error) {
	client := &http.Client{
		Timeout: time.Millisecond * 700,
	}
	var data url.Values = url.Values{}
	data.Add("requestID", requestID)
	data.Add("apiID", strconv.Itoa(apiID))
	data.Add("apiName", apiName)
	data.Add("requestURL", requestURL)
	data.Add("targetServer", targetServer)
	data.Add("targetURL", targetURL)
	data.Add("clusterName", node_common.ClusterName())
	data.Add("requestMethod", requestMethod)
	data.Add("proxyMethod", proxyMethod)
	data.Add("headerList", headerList)
	data.Add("queryParamList", queryParamList)
	data.Add("formParamList", formParamList)
	data.Add("responseHeaderList", responseHeaderList)
	data.Add("alertPeriodType", strconv.Itoa(alertPeriodType))
	data.Add("alertCount", strconv.Itoa(alertCount))
	data.Add("responseStatus", strconv.Itoa(responseStatus))
	data.Add("isAlert", isAlert)
	data.Add("strategyID", strategyID)
	data.Add("strategyName", strategyName)
	data.Add("nodePort", strconv.Itoa(node_common.ListenPort))
	request, err := http.NewRequest("POST", node_common.GetAdminUrl("/alert/msg/add"), strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {

		return false, "[ERROR]Fail to create request!", err
	}
	response, err := client.Do(request)
	if err != nil {

		return false, "[ERROR]Fail to get response!", err
	}
	defer response.Body.Close()
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {

		return false, "[ERROR]Fail to get body!", err
	}

	return true, "", nil
}

func SendRequestToAlertAddress(alertAddress, requestURL, targetServer, proxyURL, msg, apiName string, apiID int) (bool, string, error) {
	if alertAddress == "" {
		return false, "[ERROR] Illegal alertAddress!", errors.New("[ERROR] Illegal alertAddress!")
	} else {
		_, err := url.Parse(alertAddress)
		if err != nil {
			return false, err.Error(), err
		}
	}
	client := &http.Client{
		Timeout: time.Millisecond * 700,
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	var data url.Values = url.Values{}
	data.Add("requestURL", requestURL)
	data.Add("targetServer", targetServer)
	data.Add("targetURL", proxyURL)
	data.Add("alertTime", now)
	data.Add("apiName", apiName)
	data.Add("apiID", strconv.Itoa(apiID))
	data.Add("msg", msg)
	request, err := http.NewRequest("POST", alertAddress, strings.NewReader(data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return false, "[ERROR]Fail to create request!", err
	}
	_, err = client.Do(request)
	if err != nil {
		return false, "[ERROR]Fail to get response!", err
	}
	return true, "", nil
}

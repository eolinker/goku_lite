package visit

import (
	"encoding/json"
	log "github.com/eolinker/goku/goku-log"
	"github.com/eolinker/goku/goku-node/common"
	"strconv"
	"time"

	node_common "github.com/eolinker/goku/goku-node/node-common"

	"github.com/eolinker/goku/common/redis-manager"
	cmd2 "github.com/eolinker/goku/goku-node/cmd"
	gateway_manager "github.com/eolinker/goku/goku-node/manager/gateway-manager"
	entity "github.com/eolinker/goku/server/entity/node-entity"
	jsoniter "github.com/json-iterator/go"
)

type alertInfo struct {
	AlertAddr       string `json:"alertAddr"`
	AlertPeriodType int    `json:"alertPeriodType"`
	ReceiveList     string `json:"receiveList"`
}

// 更新网关转发失败次数
func UpdateProxyFailureCount(apiInfo *entity.ApiExtend,
	requestMethod string,
	proxyMethod string,
	headers map[string][]string,
	queryParams map[string][]string,
	formParams map[string][]string,
	responseStatus int,
	responseHeader map[string][]string,

	ctx *common.Context) (bool, string, error) {
	redisConn := redis_manager.GetConnection()
	var (
		alertAddress    string
		alertStatus     int
		alertPeriodType int
	)
	alertStatus = gateway_manager.GetAlertStatus()

	if alertStatus == 0 || apiInfo.AlertValve == 0 {
		return true, "", nil
	} else {
		info := gateway_manager.GetAlertInfo()
		alertJson := alertInfo{}
		err := json.Unmarshal([]byte(info), &alertJson)
		if err != nil {
			return false, err.Error(), err
		}

		clusterName := node_common.ClusterName()
		alertPeriodType = alertJson.AlertPeriodType
		alertAddress = alertJson.AlertAddr
		for {
			rep, err := redisConn.SetNX(clusterName+":lock", 1, 3*time.Second).Result()
			if err != nil {
				log.Info(err)
				return false, "", err
			}
			if rep {
				break
			}
		}

		// 获取当前告警时间信息
		redisKey := clusterName + ":gokuFailureCount:" + strconv.Itoa(apiInfo.ApiID)
		// 获取key信息
		result, err := redisConn.Get(redisKey).Result()
		if err != nil {
			if err.Error() != "redis: nil" {
				log.Info(err)
				return false, "", err
			}
		}
		var expire int = 60
		if alertPeriodType == 1 {
			expire = 5 * 60
		} else if alertPeriodType == 2 {
			expire = 15 * 60
		} else if alertPeriodType == 3 {
			expire = 30 * 60
		} else if alertPeriodType == 4 {
			expire = 60 * 60
		}

		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		headerList, err := json.Marshal(headers)
		if err != nil {
			log.Info(err)
			return false, "", err
		}
		queryParamList, err := json.Marshal(queryParams)
		if err != nil {
			log.Info(err)
			return false, "", err
		}
		formParamsList, err := json.Marshal(formParams)
		if err != nil {
			log.Info(err)
			return false, "", err
		}
		responseHeaderList, err := json.Marshal(responseHeader)
		if err != nil {
			log.Info(err)
			return false, "", err
		}

		var count int = 0

		if result == "" {
			_, err = redisConn.SetNX(redisKey, 1, time.Duration(expire)*time.Second).Result()
			if err != nil {
				log.Info(err)

				return false, "", err
			}
			count = 1
		} else {
			count, err = strconv.Atoi(result)
			if err != nil {
				log.Info(err)

				return false, "", err
			}
		}

		if count >= apiInfo.AlertValve {

			_, _, _ = cmd2.AddAlertMessage(apiInfo.ApiID,
				apiInfo.ApiName,
				apiInfo.RequestURL,
				ctx.FinalTargetServer(),
				apiInfo.TargetURL,
				requestMethod,
				proxyMethod,
				string(headerList),
				string(queryParamList),
				string(formParamsList),
				string(responseHeaderList),
				alertPeriodType,
				apiInfo.AlertValve,
				responseStatus,
				"true",
				ctx.StrategyId(),
				ctx.StrategyName(),
				ctx.RequestId())
			AlertLog(
				apiInfo.RequestURL,
				ctx.FinalTargetServer(),
				apiInfo.TargetURL,
				requestMethod,
				proxyMethod,
				string(headerList),
				string(queryParamList),
				string(formParamsList),
				string(responseHeaderList),
				responseStatus,
				ctx.StrategyId(),
				ctx.StrategyName(),
				ctx.RequestId())

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
			msg := "[Alert] GoKu Gateway failed to proxy requests for " + period + " minute " + strconv.Itoa(apiInfo.AlertValve) + " times"
			cmd2.SendRequestToAlertAddress(alertAddress, apiInfo.RequestURL, ctx.FinalTargetServer(), apiInfo.TargetURL, msg, apiInfo.ApiName, apiInfo.ApiID)

			_, err = redisConn.Set(redisKey, 1, time.Duration(expire)*time.Second).Result()
			if err != nil {
				log.Info(err)
			}
		} else {
			_, _, _ = cmd2.AddAlertMessage(
				apiInfo.ApiID,
				apiInfo.ApiName,
				apiInfo.RequestURL,
				ctx.FinalTargetServer(),
				apiInfo.TargetURL,
				requestMethod,
				proxyMethod,
				string(headerList),
				string(queryParamList),
				string(formParamsList),
				string(responseHeaderList),
				alertPeriodType,
				apiInfo.AlertValve,
				responseStatus,
				"false",
				ctx.StrategyId(),
				ctx.StrategyName(),
				ctx.RequestId())
			AlertLog(
				apiInfo.RequestURL,
				ctx.FinalTargetServer(),
				apiInfo.TargetURL,
				apiInfo.RequestMethod,
				apiInfo.TargetMethod,
				string(headerList),
				string(queryParamList),
				string(formParamsList),
				string(responseHeaderList),
				responseStatus,
				ctx.StrategyId(),
				ctx.StrategyName(),
				ctx.RequestId())
			if _, err := redisConn.Incr(redisKey).Result(); err != nil {

				return false, err.Error(), err
			}
		}
		redisConn.Del("lock")
		return true, "", nil
	}
}

// 记录告警日志
func AlertLog(requestURL, targetServer, targetURL, requestMethod, proxyMethod, headerList, queryParamList, formParamList, responseHeaderList string, responseStatus int, strategyID string, strategyName, requestID string) {
	  log.WithFields(log.Fields{
		"request_id":requestID,
		"strategy_name":strategyName,
		"strategy_id":strategyID,
		"request_method":requestMethod,
		"request_url":requestURL,
		"target_method":proxyMethod,
		"target_server":targetServer,
		"target_url":targetURL,
		"request_query":queryParamList,
		"request_header":headerList,
		"request_form_param":formParamList,
		"response_statusCode":strconv.Itoa(responseStatus),
		"response_header":responseHeaderList,
	}).Warning("alert")

	 //_= logutils.Log("log/alertLog", "alert", log.PeriodDay, logInfo)

	return
}

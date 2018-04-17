package middleware

import (
	"goku-ce-1.0/dao"
	"net/http"
	"strings"
	"time"
	"fmt"
	"github.com/farseer810/yawf"
)

func GatewayValve(httpRequest *http.Request,context yawf.Context, httpResponse http.ResponseWriter) (bool, string) {
	t1:= time.Now()
	var gatewayHashkey string
	// 获取请求路径中的网关别名
	requestInfo := strings.Split(httpRequest.RequestURI,"/")
	gatewayAlias := requestInfo[1]
	strategyKey := requestInfo[2]

	// 通过网关别名获取网关hashKey
	gatewayHashkey = dao.GetGatewayHashKey(context,gatewayAlias)
	if gatewayHashkey == "" {
		httpResponse.WriteHeader(404)
		return false, "error gatewayAlias"
	}

	remoteAddr := httpRequest.RemoteAddr
	remoteIP := InterceptIP(remoteAddr, ":")

	strategySecondValve := dao.GetStrategySecondValve(context,gatewayHashkey,strategyKey)
	if strategySecondValve.LimitCount != -1{
		secondCount := dao.GetGatewayStrategyIPSecondCount(context,gatewayHashkey,strategyKey,remoteIP)
		if secondCount >= strategySecondValve.LimitCount {
			httpResponse.WriteHeader(408)
			return false, "second visit limit exceeded"
		}
	}else if strategySecondValve.ViewType == 1 {
		httpResponse.WriteHeader(408)
		return false, "Don't allow visit"
	}

	strategyMinuteValve := dao.GetStrategyMinuteValve(context,gatewayHashkey,strategyKey)
	if strategyMinuteValve.LimitCount != -1{
		minuteCount := dao.GetGatewayStrategyIPMinuteCount(context,gatewayHashkey,strategyKey,remoteIP)
		if minuteCount >= strategyMinuteValve.LimitCount {
			httpResponse.WriteHeader(408)
			return false, "minute visit limit exceeded"
		}
	}else if strategyMinuteValve.ViewType == 1 {
		httpResponse.WriteHeader(408)
		return false, "Don't allow visit"
	}

	strategyHourValve := dao.GetStrategyHourValve(context,gatewayHashkey,strategyKey)
	fmt.Println(strategyHourValve.LimitCount)
	if strategyHourValve.LimitCount != -1{
		hourCount := dao.GetGatewayStrategyIPHourCount(context,gatewayHashkey,strategyKey,remoteIP)
		fmt.Println(hourCount)
		if hourCount >= strategyHourValve.LimitCount {
			httpResponse.WriteHeader(408)
			return false, "hour visit limit exceeded"
		}
	}else if strategyHourValve.ViewType == 1 {
		httpResponse.WriteHeader(408)
		return false, "Don't allow visit"
	}
	
	strategyDayValve := dao.GetStrategyDayValve(context,gatewayHashkey,strategyKey)
	if strategyDayValve.LimitCount != -1{
		// 获取某一时段访问次数
		if strategyDayValve.StartTime != strategyDayValve.EndTime{
			dayCount := dao.GetStrategyPeriodCount(context,gatewayHashkey,strategyKey,strategyDayValve.StartTime,strategyDayValve.EndTime)
			fmt.Println(dayCount)
			if dayCount >= strategyDayValve.LimitCount {
				httpResponse.WriteHeader(408)
				return false, "day visit limit exceeded"
			}
		}

	}else if strategyDayValve.ViewType == 1 {
		httpResponse.WriteHeader(408)
		return false, "Don't allow visit"
	}
	
	fmt.Println("valve time:",time.Since(t1))
	return true,""
}

package middleware

import (
	"goku-ce/goku"
	"net/http"
	"strings"
)


func IPLimit(g *goku.Context,res http.ResponseWriter, req *http.Request) (bool,string) {
    remoteAddr := req.RemoteAddr
	remoteIP := InterceptIP(remoteAddr, ":")
	if !globalIPLimit(g,remoteIP){
		return false,"[Global] Illegal IP"
	} else if !strategyIPLimit(g,remoteIP) {
		return false,"[Strategy] Illegal IP"
	}
	return true,""
}

func globalIPLimit(g *goku.Context,remoteIP string) bool{
	if g.GatewayInfo.IPLimitType == "black"{
		for _,ip := range g.GatewayInfo.IPBlackList{
			if ip == remoteIP {
				return false
			}
		}
		return true
	} else if g.GatewayInfo.IPLimitType == "white" {
		for _,ip := range g.GatewayInfo.IPWhiteList{
			if ip == remoteIP {
				return true
			}
		}
		return false
	}
	return true
}

func strategyIPLimit(g *goku.Context,remoteIP string) bool {
	if g.StrategyInfo.IPLimitType == "black" {
		for _,ip := range g.StrategyInfo.IPBlackList{
			if ip == remoteIP {
				return false
			}
		}
		return true
	} else if g.StrategyInfo.IPLimitType == "white" {
		for _,ip := range g.StrategyInfo.IPWhiteList{
			if ip == remoteIP {
				return true
			}
		}
		return false
	}
	return true
}

func InterceptIP(str, substr string) string {
	result := strings.Index(str, substr)
	var rs string
	if result > 7 {
		rs = str[:result]
	}
	return rs
}
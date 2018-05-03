package middleware

import (
	"goku-ce/conf"
	"net/http"
	"strings"
)


func IPLimit(g conf.GatewayInfo,d conf.StrategyInfo,res http.ResponseWriter, req *http.Request) (bool,string) {
    remoteAddr := req.RemoteAddr
	remoteIP := InterceptIP(remoteAddr, ":")
	if !globalIPLimit(g,remoteIP){
		res.WriteHeader(404)
		return false,"[global] Illegal ip"
	} else if globalIPLimit(g,remoteIP) && !strategyIPLimit(d,remoteIP) {
		res.WriteHeader(404)
		return false,"[strategy] Illegal ip"
	}
	return true,""
}

func globalIPLimit(g conf.GatewayInfo,remoteIP string) bool{
	if g.IPLimitType == "black"{
		for _,ip := range g.IPBlackList{
			if ip == remoteIP {
				return false
			}
		}
		return true
	} else if g.IPLimitType == "white" {
		for _,ip := range g.IPWhiteList{
			if ip == remoteIP {
				return true
			}
		}
		return false
	}
	return true
}

func strategyIPLimit(d conf.StrategyInfo,remoteIP string) bool {
	if d.IPLimitType == "black" {
		for _,ip := range d.IPBlackList{
			if ip == remoteIP {
				return false
			}
		}
		return true
	} else if d.IPLimitType == "white" {
		for _,ip := range d.IPWhiteList{
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
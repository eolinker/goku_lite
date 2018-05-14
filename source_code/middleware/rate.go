package middleware

import (
	"fmt"
	"time"
	"goku-ce/conf"
	"goku-ce/goku"
)

func getStrategyRate(context *goku.Context) (bool,[]conf.RateLimitInfo) {
	c := context.StrategyInfo
	now := time.Now()
	flag := false
	rateLimitList := make([]conf.RateLimitInfo,0)
	var secRate,minRate,hourRate,dayRate,banRate conf.RateLimitInfo
	for _,i := range c.RateLimitList{
		if timeInPeriod(i,now.Hour()) {
			if i.Allow {
				if i.Period == "sec" && i.Priority > secRate.Priority {
					secRate = i
				} else if i.Period == "min" && i.Priority > minRate.Priority {
					minRate = i
				} else if i.Period == "hour" && i.Priority > hourRate.Priority {
					hourRate = i
				} else if i.Period == "day" && i.Priority > dayRate.Priority {
					dayRate = i
				}
			}else if i.Allow == false && i.Priority > banRate.Priority {
				banRate = i
			}
		}
	}
	rateLimitList = append(rateLimitList,secRate)
	rateLimitList = append(rateLimitList,minRate)
	rateLimitList = append(rateLimitList,hourRate)
	rateLimitList = append(rateLimitList,dayRate)
	rateLimitList = append(rateLimitList,banRate)
	priority := 0
	allow := true 
	for _,i := range rateLimitList {
		if i.Priority > priority {
			priority = i.Priority
			allow = i.Allow
		}
		if i.Priority == priority && !i.Allow && i.Priority != 0{
			allow = i.Allow
		}
	}
	if allow {
		flag = true 
	}
	return flag,rateLimitList
}

func timeInPeriod(c conf.RateLimitInfo,now int) bool {
	if c.StartTime <= now && now < c.EndTime {
		return true
	} else if c.StartTime > c.EndTime {
		if now >= c.StartTime && now < c.EndTime + 24 {
			return true
		} else if now < c.StartTime && now < c.EndTime {
			return true
		}
	}
	return false
}

func RateLimit(context *goku.Context) (bool,string) {
	c := context.StrategyInfo
	g := context.Rate
	value, ok := g[c.StrategyID]
	if !ok {
		var w goku.Rate
		g[c.StrategyID] = w
	} 
	if !value.IsInit{
		flag,r := getStrategyRate(context)
		if flag == false {
			return false,"Forbidden Request"
		}
		for _,i := range r {
			if i.Period == "sec" {
				value.SecLimit.SetRate(i.Limit,i.EndTime,"sec")
			} else if i.Period == "min" {
				value.MinuteLimit.SetRate(i.Limit,i.EndTime,"min")
			} else if i.Period == "hour" {
				value.HourLimit.SetRate(i.Limit,i.EndTime,"min")
			} else if i.Period == "day" {
				value.DayLimit.SetRate(i.Limit,i.EndTime,"init")
			}
		}
		value.IsInit = true
	} else if value.SecLimit.IsNeedReset(){
		flag,r := getStrategyRate(context)
		if flag == false {
			return false,"Forbidden Request"
		}
		for _,i := range r {
			if i.Period == "sec" {
				value.SecLimit.SetRate(i.Limit,i.EndTime,"sec")
			} else if i.Period == "min" {
				value.MinuteLimit.SetRate(i.Limit,i.EndTime,"min")
			} else if i.Period == "hour" {
				value.HourLimit.SetRate(i.Limit,i.EndTime,"hour")
			} else if i.Period == "day" {
				value.DayLimit.SetRate(i.Limit,i.EndTime,"day")
			}
		}
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	if value.Limit == "day" {
		if !value.DayLimit.DayLimit() {
			value.Limit = "day"
			g[c.StrategyID] = value
			return false,"API Rate Limit of Day Exceeded"
		}
		value.Limit = ""
	} else if value.Limit == "hour" {
		if !value.HourLimit.HourLimit() {
			value.Limit = "hour"
			g[c.StrategyID] = value
			return false,"API Rate Limit of Hour Exceeded"
		}else if !value.DayLimit.DayLimit() {
			value.Limit = "day"
			g[c.StrategyID] = value
			return false,"API Rate Limit of Day Exceeded"
		}
		value.Limit = ""
	} else if value.Limit == "minute" {
		if !value.MinuteLimit.MinLimit() {
			value.Limit = "minute"
			g[c.StrategyID] = value
			return false,"API Rate Limit of Minute Exceeded"
		}else if !value.HourLimit.HourLimit() {
			value.Limit = "hour"
			g[c.StrategyID] = value
			return false,"API Rate Limit of Hour Exceeded"
		}else if !value.DayLimit.DayLimit() {
			value.Limit = "day"
			g[c.StrategyID] = value
			return false,"API Rate Limit of Day Exceeded"
		}
		value.Limit = ""
	} else {
		if !value.SecLimit.SecLimit() {
			g[c.StrategyID] = value
			return false,"API Rate Limit of Second Exceeded"
		}else if !value.MinuteLimit.MinLimit() {
			value.Limit = "minute"
			g[c.StrategyID] = value
			return false,"API Rate Limit of Minute Exceeded"
		}else if !value.HourLimit.HourLimit() {
			value.Limit = "hour"
			g[c.StrategyID] = value
			return false,"API Rate Limit of Hour Exceeded"
		}else if !value.DayLimit.DayLimit() {
			value.Limit = "day"
			g[c.StrategyID] = value
			return false,"API Rate Limit of Day Exceeded"
		}
	}
	g[c.StrategyID] = value
	return true,""
}
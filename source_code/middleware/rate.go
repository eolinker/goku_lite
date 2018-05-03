package middleware

import (
	"time"
	"goku-ce/conf"
	"goku-ce/goku"
)

func getStrategyRate(c conf.StrategyInfo) (bool,[]conf.RateLimitInfo) {
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

func RateLimit(g *goku.Goku,c conf.StrategyInfo) (bool,string) {
	value, ok := g.Rate[c.StrategyID]
	if !ok {
		var w goku.Rate
		g.Rate[c.StrategyID] = w
	} 
	if !value.IsInit{
		flag,r := getStrategyRate(c)
		if flag == false {
			return false,"Don't allow visit!"
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
		flag,r := getStrategyRate(c)
		if flag == false {
			return false,"Don't allow visit!"
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
	
	if value.Limit == "day" {
		if !value.DayLimit.DayLimit() {
			value.Limit = "day"
			g.Rate[c.StrategyID] = value
			return false,"Day visit limit exceeded"
		}
		value.Limit = ""
	} else if value.Limit == "hour" {
		if !value.HourLimit.HourLimit() {
			value.Limit = "hour"
			g.Rate[c.StrategyID] = value
			return false,"Hour visit limit exceeded"
		}else if !value.DayLimit.DayLimit() {
			value.Limit = "day"
			g.Rate[c.StrategyID] = value
			return false,"Day visit limit exceeded"
		}
		value.Limit = ""
	} else if value.Limit == "minute" {
		if !value.MinuteLimit.MinLimit() {
			value.Limit = "minute"
			g.Rate[c.StrategyID] = value
			return false,"Minute visit limit exceeded"
		}else if !value.HourLimit.HourLimit() {
			value.Limit = "hour"
			g.Rate[c.StrategyID] = value
			return false,"Hour visit limit exceeded"
		}else if !value.DayLimit.DayLimit() {
			value.Limit = "day"
			g.Rate[c.StrategyID] = value
			return false,"Day visit limit exceeded"
		}
		value.Limit = ""
	} else {
		if !value.SecLimit.SecLimit() {
			g.Rate[c.StrategyID] = value
			return false,"Second visit limit exceeded"
		}else if !value.MinuteLimit.MinLimit() {
			value.Limit = "minute"
			g.Rate[c.StrategyID] = value
			return false,"Minute visit limit exceeded"
		}else if !value.HourLimit.HourLimit() {
			value.Limit = "hour"
			g.Rate[c.StrategyID] = value
			return false,"Hour visit limit exceeded"
		}else if !value.DayLimit.DayLimit() {
			value.Limit = "day"
			g.Rate[c.StrategyID] = value
			return false,"Day visit limit exceeded"
		}
	}
	g.Rate[c.StrategyID] = value
	return true,""
}
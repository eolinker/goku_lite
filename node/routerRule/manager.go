package routerRule

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/utils"
)

//Router router
type Router struct {
	Host       string `json:"host"`
	StrategyID string `json:"strategyID"`
	ID         string `json:"id"`
}

//Match match
func (r *Router) Match(c *config.RouterRule) (bool, string) {
	if r.Host != "" && !matchHost(r.Host, c.Host) {
		return false, ""
	}

	if r.StrategyID != "" && r.StrategyID != c.StrategyID {
		return false, ""
	}
	return true, r.ID
}

var router = make([]*Router, 0)

func newRouter(rs []*config.Router) []*Router {
	newRs := make([]*Router, 0, len(rs))
	if rs == nil {
		return newRs
	}
	for _, r := range rs {
		rls := make([]*config.RouterRule, 0)
		err := json.Unmarshal([]byte(r.Rules), &rls)
		if err != nil {
			continue
		}
		ts := make([]int, 0, 2)
		err = json.Unmarshal([]byte(r.Target), &ts)
		if err != nil {
			continue
		}
		if len(ts) == 1 && ts[0] == 0 {
			// 指标只有策略ID
			newRs = append(newRs, &Router{Host: "", StrategyID: ""})
			continue
		}
		commonRs := make([]*Router, 0, len(rls))
		for _, rl := range rls {
			var host, strategyID string
			if strings.Contains(r.Target, "0") {
				// 指标包括策略ID
				strategyID = rl.StrategyID
			}
			if strings.Contains(r.Target, "1") {
				// 指标包括Host
				host = rl.Host
			}
			commonRs = append(commonRs, &Router{Host: host, StrategyID: strategyID, ID: rl.StrategyID})
		}
		sort.Sort(Routers(commonRs))
		newRs = append(newRs, commonRs...)
	}
	return newRs
}

//Load load
func Load(rs []*config.Router) {
	router = newRouter(rs)
}

//Get get
func Get() []*Router {
	return router
}

func matchHost(org, match string) bool {
	if org == "*" || org == match {
		return true
	}

	_, o := utils.Intercept(org, ".")
	_, m := utils.Intercept(match, ".")
	if o == m && string(org[0]) == "*" {
		return true
	}

	return false
}

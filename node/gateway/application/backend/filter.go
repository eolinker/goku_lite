package backend

import (
	"github.com/eolinker/goku-api-gateway/config"
	"github.com/eolinker/goku-api-gateway/node/gateway/application/action"
)

func genFilter(blackList, whiteList []string, acs []*config.ActionConfig) action.Filter {
	size := len(blackList) + 1 + len(acs)

	filters := make(action.Filters, 0, size)

	for _, b := range blackList {
		filters = append(filters, action.Blacklist(b))
	}
	if len(whiteList) > 0 {
		filters = append(filters, action.GenWhite(whiteList))
	}

	for _, ac := range acs {

		f := action.GenByconfig(ac)
		if f != nil {
			filters = append(filters, f)
		}
	}
	return filters
}

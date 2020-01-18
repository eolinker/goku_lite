package dao_version_config

import (
	"fmt"
	"strings"

	"github.com/eolinker/goku-api-gateway/config"
)

//GetRouterRules GetRouterRules
func (d *VersionConfigDao) GetRouterRules(enable int) ([]*config.Router, error) {
	db := d.db
	sql := "SELECT rules,target FROM goku_gateway_router %s ORDER BY priority DESC;"
	rules := make([]string, 0, 1)
	if enable != -1 {
		rules = append(rules, fmt.Sprintf("enable = %d", enable))
	}
	ruleStr := ""
	if len(rules) > 0 {
		ruleStr += "WHERE " + strings.Join(rules, " AND ")
	}
	rows, err := db.Query(fmt.Sprintf(sql, ruleStr))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rs := make([]*config.Router, 0)
	for rows.Next() {
		var r config.Router
		err = rows.Scan(&r.Rules, &r.Target)
		if err != nil {
			return nil, err
		}
		rs = append(rs, &r)
	}
	return rs, nil
}

package daostrategy

import (
	log "github.com/eolinker/goku-api-gateway/goku-log"

	"github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

//GetAllStrategy 获取所有策略列表
func GetAllStrategy() (map[string]*entity.Strategy, *entity.Strategy, error) {
	//const sql = "SELECT `strategyID`,`strategyName`,`auth` ,`strategyType` FROM `goku_gateway_strategy` WHERE `enableStatus` = 1;"
	const sql = "SELECT `strategyID`,`strategyName`,`auth`,`enableStatus` ,`strategyType` FROM `goku_gateway_strategy`;"

	stmt, e := database.GetConnection().Prepare(sql)
	if e != nil {
		return nil, nil, e
	}
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		return nil, nil, e
	}
	defer rows.Close()
	var def *entity.Strategy
	strategys := make(map[string]*entity.Strategy)
	for rows.Next() {
		s := new(entity.Strategy)
		err := rows.Scan(&s.StrategyID, &s.StrategyName, &s.Auth, &s.EnableStatus, &s.StrategyType)
		if err != nil {
			log.Warn("GetAllStrategy ", err)
			continue
		}

		strategys[s.StrategyID] = s
		if s.StrategyType == 1 {
			def = s
		}
	}

	return strategys, def, nil
}

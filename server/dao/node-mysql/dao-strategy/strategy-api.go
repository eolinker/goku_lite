package daostrategy

import (
	"github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

//GetAllStrategyAPI 获取所有策略接口列表
func GetAllStrategyAPI() ([]*entity.StrategyAPI, error) {

	const sql = "SELECT A.`strategyID`,A.`apiID`,IFNULL(A.`target`,'') FROM `goku_conn_strategy_api` A  JOIN `goku_gateway_strategy` B ON A.`strategyID` = B.`strategyID` ;"

	stmt, e := database.GetConnection().Prepare(sql)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	apis := make([]*entity.StrategyAPI, 0, 1000)

	for rows.Next() {
		api := new(entity.StrategyAPI)
		err := rows.Scan(&api.StrategyID, &api.APIID, &api.Target)
		if err != nil {
			continue
		}
		apis = append(apis, api)
	}
	return apis, nil

}

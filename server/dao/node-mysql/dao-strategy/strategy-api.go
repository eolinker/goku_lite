package dao_strategy

import (
	"github.com/eolinker/goku-api-gateway/common/database"
	entity "github.com/eolinker/goku-api-gateway/server/entity/node-entity"
)

func GetAllStrategyApi() ([]*entity.StrategyApi, error) {

	//const sql = "SELECT A.`strategyID`,A.`apiID`,A.`target` FROM `goku_conn_strategy_api` A  JOIN `goku_gateway_strategy` B ON A.`strategyID` = B.`strategyID` AND B.`enableStatus` = 1;"
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

	apis := make([]*entity.StrategyApi, 0, 1000)

	for rows.Next() {
		api := new(entity.StrategyApi)
		err := rows.Scan(&api.StrategyID, &api.ApiId, &api.Target)
		if err != nil {
			continue
		}
		apis = append(apis, api)
	}
	return apis, nil

}

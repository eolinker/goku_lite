package dao_balance_update

import (
	"github.com/eolinker/goku-api-gateway/common/database"
	dao_service2 "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/dao-service"
	entity "github.com/eolinker/goku-api-gateway/server/entity/balance-entity"
)

//GetAllOldVerSion 获取所有旧负载配置
func GetAllOldVerSion() ([]*entity.BalanceInfoEntity, error) {
	const sql = "SELECT `balanceName`,IFNULL(`balanceDesc`, ''),IFNULL(`balanceConfig`, ''),IFNULL(`defaultConfig`, ''),IFNULL(`clusterConfig`, ''),`updateTime`,`createTime` FROM `goku_balance` WHERE `serviceName` = '';"
	db := database.GetConnection()
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	r := make([]*entity.BalanceInfoEntity, 0, 20)
	for rows.Next() {
		v := new(entity.BalanceInfoEntity)
		err := rows.Scan(&v.Name, &v.Desc, &v.OldVersionConfig, &v.DefaultConfig, &v.ClusterConfig, &v.UpdateTime, &v.CreateTime)
		if err != nil {
			return nil, err
		}
		r = append(r, v)
	}

	return r, nil
}

//GetDefaultServiceStatic 获取默认静态负载
func GetDefaultServiceStatic() string {

	tx := database.GetConnection()
	name := ""
	err := tx.QueryRow("SELECT `name` FROM `goku_service_config` WHERE `driver`='static' ORDER BY  `default` DESC LIMIT 1; ").Scan(&name)
	if err != nil {
		name = "static"
		dao_service2.Add(name, "static", "默认静态服务", "", "", false, false, "", "", 5, 300)
	}

	return name
}

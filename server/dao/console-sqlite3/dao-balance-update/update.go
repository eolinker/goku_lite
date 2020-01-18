package dao_balance_update

import (
	"database/sql"

	"github.com/eolinker/goku-api-gateway/common/pdao"
	"github.com/eolinker/goku-api-gateway/server/dao"
	entity "github.com/eolinker/goku-api-gateway/server/entity/balance-entity"
)

var serviceDao dao.ServiceDao

func init() {
	pdao.Need(&serviceDao)
}

//BalanceUpdateDao BalanceUpdateDao
type BalanceUpdateDao struct {
	db *sql.DB
}

//NewBalanceUpdateDao new BalanceUpdateDao
func NewBalanceUpdateDao() *BalanceUpdateDao {
	return &BalanceUpdateDao{}
}

//Create create
func (d *BalanceUpdateDao) Create(db *sql.DB) (interface{}, error) {
	d.db = db
	i := dao.BalanceUpdateDao(d)
	return &i, nil
}

//GetAllOldVerSion 获取所有旧负载配置
func (d *BalanceUpdateDao) GetAllOldVerSion() ([]*entity.BalanceInfoEntity, error) {
	const sql = "SELECT `balanceName`,IFNULL(`balanceDesc`, ''),IFNULL(`balanceConfig`, ''),IFNULL(`defaultConfig`, ''),IFNULL(`clusterConfig`, ''),`updateTime`,`createTime` FROM `goku_balance` WHERE `serviceName` = '';"
	db := d.db
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
func (d *BalanceUpdateDao) GetDefaultServiceStatic() string {

	tx := d.db
	name := ""
	err := tx.QueryRow("SELECT `name` FROM `goku_service_config` WHERE `driver`='static' ORDER BY  `default` DESC LIMIT 1; ").Scan(&name)
	if err != nil {
		name = "static"
		serviceDao.Add(name, "static", "默认静态服务", "", "", false, false, "", "", 5, 300)
	}

	return name
}

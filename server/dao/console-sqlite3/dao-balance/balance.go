package dao_balance

import (
	"database/sql"

	dao "github.com/eolinker/goku-api-gateway/server/dao"
)

//BalanceDao BalanceDao
type BalanceDao struct {
	db *sql.DB
}

//NewBalanceDao new BalanceDao
func NewBalanceDao() *BalanceDao {
	return &BalanceDao{}
}

//Create create
func (b *BalanceDao) Create(db *sql.DB) (interface{}, error) {
	b.db = db
	i := dao.BalanceDao(b)
	return &i, nil
}

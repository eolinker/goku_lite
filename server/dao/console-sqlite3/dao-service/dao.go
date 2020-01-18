package dao_service

import (
	"database/sql"

	"github.com/eolinker/goku-api-gateway/server/dao"
)

//ServiceDao ServiceDao
type ServiceDao struct {
	db *sql.DB
}

//NewServiceDao new ServiceDao
func NewServiceDao() *ServiceDao {
	return &ServiceDao{}
}

//Create create
func (d *ServiceDao) Create(db *sql.DB) (interface{}, error) {
	d.db = db
	i := dao.ServiceDao(d)
	return &i, nil
}

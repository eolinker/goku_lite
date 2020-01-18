package dao_version_config

import (
	"database/sql"

	"github.com/eolinker/goku-api-gateway/server/dao"
)

//VersionConfigDao VersionConfigDao
type VersionConfigDao struct {
	db *sql.DB
}

//NewVersionConfigDao new VersionConfigDao
func NewVersionConfigDao() *VersionConfigDao {
	return &VersionConfigDao{}
}

//Create create
func (d *VersionConfigDao) Create(db *sql.DB) (interface{}, error) {
	d.db = db

	i := dao.VersionConfigDao(d)
	return &i, nil
}

package updater

import (
	"database/sql"

	"github.com/eolinker/goku-api-gateway/server/dao"
)

//Dao Dao
type Dao struct {
	db *sql.DB
}

//NewUpdaterDaoWidthDB NewUpdaterDaoWidthDB
func NewUpdaterDaoWidthDB(db *sql.DB) *Dao {
	return &Dao{db: db}
}

//NewUpdaterDao NewUpdaterDao
func NewUpdaterDao() *Dao {
	return &Dao{}
}

//Create create
func (d *Dao) Create(db *sql.DB) (interface{}, error) {
	d.db = db

	i := dao.UpdaterDao(d)
	return &i, nil
}

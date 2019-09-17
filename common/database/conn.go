package database

import (
	"database/sql"
	log "github.com/eolinker/goku/goku-log"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbMysql = "mysql"
)

var (
	defaultDB *sql.DB
)

func InitConnection(config Config) error {
	db, e := getConnection(config)
	defaultDB = db
	return e
}
func getConnection(config Config) (*sql.DB, error) {

	//var dsn string

	//dsn = conf.Value("db_user") + ":" + conf.Value("db_password")
	//dsn = dsn + "@tcp(" + conf.Value("db_host") + ":" + conf.Value("db_port") + ")/" + conf.Value("db_name")
	//dsn = dsn + "?charset=utf8"

	db, e := sql.Open(config.GetDriver(), config.GetSource())

	if e == nil {
		if err := db.Ping(); err != nil {
			log.Info(err)
			return nil, err
		}
		db.SetMaxOpenConns(1000)
		db.SetMaxIdleConns(100)
		defaultDB = db
		return db, nil
	} else {
		log.Info(e)
		return nil, e
	}

}

func GetConnection() *sql.DB {

	return defaultDB
}

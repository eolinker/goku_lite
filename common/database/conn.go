package database

import (
	"database/sql"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	//mysql数据库驱动
	_ "github.com/go-sql-driver/mysql"
	//sqlite3数据库驱动
	_ "github.com/mattn/go-sqlite3"
)

//InitConnection 初始化数据库连接
func InitConnection(config Config) (*sql.DB, error) {
	return getConnection(config)
}
func getConnection(config Config) (*sql.DB, error) {

	db, e := sql.Open(config.GetDriver(), config.GetSource())

	if e == nil {
		if err := db.Ping(); err != nil {
			log.Info(err)
			return nil, err
		}
		db.SetMaxOpenConns(1000)
		db.SetMaxIdleConns(100)
		return db, nil
	}
	log.Info(e)
	return nil, e

}

//CheckConnection 检查数据库连接
func CheckConnection(driver string, source string) error {

	db, e := sql.Open(driver, source)
	defer db.Close()
	if e == nil {
		if err := db.Ping(); err != nil {
			return err
		}
		return nil
	}
	return e

}

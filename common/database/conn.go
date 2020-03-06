package database

import (
	"database/sql"
	"io/ioutil"
	"strings"

	log "github.com/eolinker/goku-api-gateway/goku-log"

	//mysql数据库驱动
	_ "github.com/go-sql-driver/mysql"
	//sqlite3数据库驱动
	_ "github.com/mattn/go-sqlite3"
)

var (
	defaultDB *sql.DB
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
		defaultDB = db
		return db, nil
	}
	log.Info(e)
	return nil, e

}

//GetConnection 获取数据库连接
func GetConnection() *sql.DB {
	return defaultDB
}

//InitTable 初始化表
func InitTable() error {

	content, err := ioutil.ReadFile("sql/goku_ce.sql")
	sqls := strings.Split(string(content), ";")
	Tx, _ := GetConnection().Begin()
	for _, sql := range sqls {
		_, err = Tx.Exec(sql)
		if err != nil {
			Tx.Rollback()
			log.Error("InitTable error:",err,"\t sql:",sql)
			return err
		}
	}
	Tx.Commit()
	return nil
}

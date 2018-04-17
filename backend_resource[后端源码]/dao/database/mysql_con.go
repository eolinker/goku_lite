package database

import (
	"goku-ce-1.0/conf"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	var err error
	var dsn string = conf.Configure["mysql_username"] + ":" + conf.Configure["mysql_password"]
	dsn = dsn + "@tcp(" + conf.Configure["mysql_host"] + ":" + conf.Configure["mysql_port"] + ")/" + conf.Configure["mysql_dbname"]
	dsn = dsn + "?charset=utf8"
	db, err = sql.Open("mysql", dsn)
	if err == nil {
		db.SetMaxOpenConns(2000)
		db.SetMaxIdleConns(1000)
	} else {
		panic(err)
	}
}

func GetConnection() *sql.DB {
	return db
}

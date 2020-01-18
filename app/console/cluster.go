package main

import (
	"github.com/eolinker/goku-api-gateway/common/database"
	"strconv"

	"github.com/pkg/errors"

	"github.com/eolinker/goku-api-gateway/common/conf"
	"github.com/eolinker/goku-api-gateway/server/entity"
)

func getDefaultDatabase() (*entity.ClusterDB, error) {
	dbType := conf.MastValue("db_type", "sqlite3")
	switch dbType {
	case database.MysqlDriver:
		{
			dbPort, err := strconv.Atoi(conf.MastValue("db_port", "3306"))
			if err != nil {
				return nil, err
			}
			return &entity.ClusterDB{
				Driver:   dbType,
				Host:     conf.Value("db_host"),
				Port:     dbPort,
				UserName: conf.Value("db_user"),
				Password: conf.Value("db_password"),
				Database: conf.Value("db_name"),
			}, nil
		}
	case database.Sqlite3Driver:
		{
			return &entity.ClusterDB{
				Driver: dbType,
				Path:   conf.MastValue("db_path", "./work/goku.db"),
			}, nil
		}
	default:
		{
			return nil, errors.New("unsupported database type")
		}
	}

}

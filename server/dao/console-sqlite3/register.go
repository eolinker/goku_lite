package console_sqlite3

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"

	"github.com/eolinker/goku-api-gateway/common/pdao"
	config_log "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/config-log"
	dao_balance "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/dao-balance"
	dao_balance_update "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/dao-balance-update"
	dao_service "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/dao-service"
	dao_version_config "github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/dao-version-config"
	"github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/internal/goku311"
	"github.com/eolinker/goku-api-gateway/server/dao/console-sqlite3/updater"
)

//DBDriver db驱动类型
const DBDriver = "sqlite3"

//DoRegister 注册数据库
func DoRegister() {

	pdao.RegisterDBBuilder(DBDriver, new(TableBuilder))
	goku311.RegisterUpdate()

	pdao.RegisterDao(DBDriver, NewAPIDao(), NewAPIGroupDao(), NewAPIPluginDao(), NewAPIStrategyDao())
	pdao.RegisterDao(DBDriver, NewAuthDao())
	pdao.RegisterDao(DBDriver, NewClusterDao())
	pdao.RegisterDao(DBDriver, NewGatewayDao())
	pdao.RegisterDao(DBDriver, NewGuestDao())
	pdao.RegisterDao(DBDriver, NewImportDao())
	pdao.RegisterDao(DBDriver, NewMonitorModulesDao())
	pdao.RegisterDao(DBDriver, NewNodeDao(), NewNodeGroupDao())
	pdao.RegisterDao(DBDriver, NewPluginDao())
	pdao.RegisterDao(DBDriver, NewProjectDao())
	pdao.RegisterDao(DBDriver, NewStrategyDao(), NewStrategyGroupDao(), NewStrategyPluginDao())
	pdao.RegisterDao(DBDriver, NewUserDao())
	pdao.RegisterDao(DBDriver, NewVersionDao())

	pdao.RegisterDao(DBDriver, config_log.NewConfigLogDao())
	pdao.RegisterDao(DBDriver, dao_balance.NewBalanceDao())
	pdao.RegisterDao(DBDriver, dao_balance_update.NewBalanceUpdateDao())
	pdao.RegisterDao(DBDriver, dao_service.NewServiceDao())
	pdao.RegisterDao(DBDriver, dao_version_config.NewVersionConfigDao())
	pdao.RegisterDao(DBDriver, updater.NewUpdaterDao())
}

//TableBuilder tableBuilder
type TableBuilder struct {
}

//Build build
func (t *TableBuilder) Build(db *sql.DB) error {

	userDao := NewUserDao()
	userDao.db = db

	_, err := userDao.CheckSuperAdminCount()
	if err != nil {
		content, err := ioutil.ReadFile("sql/goku_ce.sql")
		sqls := strings.Split(string(content), ";")
		Tx, _ := db.Begin()
		for _, sql := range sqls {
			_, err = Tx.Exec(sql)
			if err != nil {
				Tx.Rollback()
				log.Panic("InitTable error:", err, "\t sql:", sql)
				return err
			}
		}
		Tx.Commit()
	}

	return nil
}

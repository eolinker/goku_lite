package dao_service

import (
	"fmt"

	"github.com/eolinker/goku-api-gateway/common/database"
)

//SetDefault 设置默认服务
func SetDefault(name string) error {
	count := 0
	err := database.GetConnection().QueryRow("SELECT count(1) FROM `goku_service_config` WHERE `name` = ?;", name).Scan(&count)
	if err != nil {

		return err
	}
	if count != 1 {
		return fmt.Errorf("has no name=%s", name)
	}

	tx, e := database.GetConnection().Begin()
	if e != nil {
		return e
	}

	if _, e := tx.Exec("UPDATE `goku_service_config` SET  `default` = 0 ;"); e != nil {
		tx.Rollback()
		return e
	}
	if _, e := tx.Exec("UPDATE `goku_service_config` SET  `default` = 1 WHERE `name`=? ;", name); e != nil {
		tx.Rollback()
		return e
	}

	return tx.Commit()
}

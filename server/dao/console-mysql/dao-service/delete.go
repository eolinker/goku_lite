package dao_service

import (
	"fmt"
	"github.com/eolinker/goku-api-gateway/common/database"
)

const sqlDelete = "DELETE FROM  `goku_service_config` WHERE  `name` = ? AND NOT EXISTS (SELECT * FROM `goku_balance` B WHERE B.`serviceName` =  `goku_service_config`.`name` ) "

type DeleteError string

func (e DeleteError) Error() string {
	return fmt.Sprintf("can not delete :%s", string(e))
}

func Delete(names []string) error {

	tx, err := database.GetConnection().Begin()
	if err != nil {
		return err
	}

	stmt, e := tx.Prepare(sqlDelete)
	if e != nil {

		return e
	}

	defer stmt.Close()

	for _, name := range names {
		r, e := stmt.Exec(name)
		if e != nil {
			tx.Rollback()
			return e
		}
		rowCount, err := r.RowsAffected()
		if err != nil {
			tx.Rollback()
			return e
		}
		if rowCount == 0 {
			tx.Rollback()
			return DeleteError(name)
		}
	}

	return tx.Commit()

}

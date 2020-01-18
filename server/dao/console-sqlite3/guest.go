package console_sqlite3

import (
	SQL "database/sql"

	"github.com/eolinker/goku-api-gateway/server/dao"

	"github.com/eolinker/goku-api-gateway/utils"
)

//GuestDao GuestDao
type GuestDao struct {
	db *SQL.DB
}

//NewGuestDao new GuestDao
func NewGuestDao() *GuestDao {
	return &GuestDao{}
}

//Create create
func (d *GuestDao) Create(db *SQL.DB) (interface{}, error) {
	d.db = db
	var i dao.GuestDao = d
	return &i, nil
}

//Login 登录
func (d *GuestDao) Login(loginCall, loginPassword string) (bool, int) {
	db := d.db
	var userID int
	err := db.QueryRow("SELECT userID FROM goku_admin WHERE loginCall = ? AND loginPassword = ?;", loginCall, loginPassword).Scan(&userID)
	if err != nil {
		return false, 0
	}
	return true, userID
}

//CheckLogin 检查用户是否登录
func (d *GuestDao) CheckLogin(userToken string, userID int) bool {
	db := d.db
	var loginPassword, loginCall string
	err := db.QueryRow("SELECT loginCall,loginPassword FROM goku_admin WHERE userID = ?;", userID).Scan(&loginCall, &loginPassword)
	if err != nil {
		return false
	}
	if utils.Md5(loginCall+loginPassword) == userToken {
		return true
	}
	return false
}

//Register 用户注册
func (d *GuestDao) Register(loginCall, loginPassword string) bool {
	db := d.db
	sql := "SELECT userID,loginPassword FROM goku_admin WHERE loginCall = ?;"
	password := ""
	userID := 0
	err := db.QueryRow(sql, loginCall).Scan(&userID, &password)
	if err != nil {
		if err == SQL.ErrNoRows {
			sql = "INSERT INTO goku_admin (loginPassword,loginCall) VALUES (?,?);"
		} else {
			return false
		}
	} else {
		if password != loginPassword {
			sql = "UPDATE goku_admin SET loginPassword = ? WHERE loginCall = ?;"
		} else {
			return true
		}
	}
	rows, err := db.Exec(sql, loginPassword, loginCall)
	if err != nil {
		return false
	}
	affectRow, _ := rows.RowsAffected()
	if affectRow > 0 {
		return true
	}
	return false
}

package console_mysql

import (
	SQL "database/sql"
	"github.com/eolinker/goku/common/database"
	"github.com/eolinker/goku/utils"
 
)

func Login(loginCall, loginPassword string) (bool, int) {
	db := database.GetConnection()
	var userID int
	err := db.QueryRow("SELECT userID FROM goku_admin WHERE loginCall = ? AND loginPassword = ?;", loginCall, loginPassword).Scan(&userID)
	if err != nil {
		return false, 0
	}
	return true, userID
}

// 检查用户是否登录
func CheckLogin(userToken string, userID int) bool {
	db := database.GetConnection()
	var loginPassword, loginCall string
	err := db.QueryRow("SELECT loginCall,loginPassword FROM goku_admin WHERE userID = ?;", userID).Scan(&loginCall, &loginPassword)
	if err != nil {
		return false
	}
	if utils.Md5(loginCall+loginPassword) == userToken {
		return true
	} else {
		return false
	}
}

// 用户注册
func Register(loginCall, loginPassword string) bool {
	db := database.GetConnection()
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
	} else {
		return false
	}
}

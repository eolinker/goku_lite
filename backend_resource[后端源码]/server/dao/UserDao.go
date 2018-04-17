package dao

import (
	"goku-ce-1.0/dao/database"	
)

// 修改密码
func EditPassword(userID int,oldPassword,newPassword string) bool{
	db := database.GetConnection()
	stmt,err := db.Prepare("UPDATE eo_admin SET loginPassword = ? WHERE userID = ? AND loginPassword = ?;")
	if err != nil{
		return false
	}
	_,err = stmt.Exec(newPassword,userID,oldPassword)
	if err != nil{
		return false
	}
	return true
}
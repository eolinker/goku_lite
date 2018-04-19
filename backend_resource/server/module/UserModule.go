package module

import (
	"goku-ce-1.0/server/dao"
)

func EditPassword(userID int,oldPassword,newPassword string) bool{
	return dao.EditPassword(userID,oldPassword,newPassword)
}
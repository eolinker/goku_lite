package console_mysql

import (
	"encoding/json"
	"errors"
	database2 "github.com/eolinker/goku/common/database"
	"github.com/eolinker/goku/utils"
)
type permissionsJson map[string]bool


// 修改账户信息
func EditPassword(oldPassword, newPassword string, userID int) (bool, string, error) {
	db := database2.GetConnection()
	// 查询旧密码是否存在
	var loginCall, password string
	oldPassword = utils.Md5(oldPassword)
	newPassword = utils.Md5(newPassword)
	sql := "SELECT loginCall,loginPassword FROM goku_admin WHERE loginPassword = ? AND userID = ?;"
	err := db.QueryRow(sql, oldPassword, userID).Scan(&loginCall, &password)
	if err != nil {
		return false, "[ERROR]Old password error!", err
	}

	sql = "UPDATE goku_admin SET loginPassword = ? WHERE loginPassword = ? AND userID = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return false, "[ERROR]Illegal SQL Statement!", err
	}
	defer stmt.Close()
	r, err := stmt.Exec(newPassword, oldPassword, userID)
	if err != nil {
		return false, "[ERROR]Fail to update data!", err
	}
	_, err = r.RowsAffected()
	if err != nil {
		return false, "[ERROR]Fail to update data!", err
	}
	return true, loginCall, nil
}

// 获取账户信息
func GetUserInfo(userID int) (bool, interface{}, error) {
	db := database2.GetConnection()
	sql := `SELECT loginCall,IFNULL(remark,""),IFNULL(permissions,""),userType FROM goku_admin WHERE userID = ?;`
	var loginCall, remark, permissions string
	var userType int
	err := db.QueryRow(sql, userID).Scan(&loginCall, &remark, &permissions, &userType)
	if err != nil {
		return false, "[ERROR]This user does not exist!", err
	}
	var perssionMap map[string]interface{}
	if permissions == "" {
		permissions = "{}"
	}
	err = json.Unmarshal([]byte(permissions), &perssionMap)
	if err != nil {
		return false, "[ERROR]Fail to parse json!", err
	}
	userInfo := map[string]interface{}{
		"userID":     userID,
		"loginCall":  loginCall,
		"remark":     remark,
		"permission": perssionMap,
		"userType":   userType,
	}
	return true, userInfo, nil
}

// 获取用户类型
func GetUserType(userID int) (bool, interface{}, error) {
	db := database2.GetConnection()
	sql := "SELECT userType FROM goku_admin WHERE userID = ?;"
	var userType int
	err := db.QueryRow(sql, userID).Scan(&userType)
	if err != nil {
		return false, "[ERROR]This user does not exist!", err
	}
	return true, userType, nil
}

// 判断是否是管理员
func CheckUserIsAdmin(userID int) (bool, string, error) {
	db := database2.GetConnection()
	sql := "SELECT userType FROM goku_admin WHERE userID = ? AND (userType = 0 OR userType = 1);"
	var userType int
	err := db.QueryRow(sql, userID).Scan(&userType)
	if err != nil {
		return false, "[ERROR]This user is not admin!", errors.New("[ERROR]This user is not admin!")
	}
	return true, "", nil
}

// 判断是否是超级管理员
func CheckUserIsSuperAdmin(userID int) (bool, string, error) {
	db := database2.GetConnection()
	sql := "SELECT userType FROM goku_admin WHERE userID = ? AND userType = 0;"
	var userType int
	err := db.QueryRow(sql, userID).Scan(&userType)
	if err != nil {
		return false, "[ERROR]This user is not super admin!", errors.New("[ERROR]This user is not super admin!")
	}
	return true, "", nil
}

func CheckSuperAdminCount() (int,error){
	db := database2.GetConnection()
	sql := "SELECT count(*) FROM goku_admin WHERE  userType = 0;"
	count := 0
	err := db.QueryRow(sql).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count  , nil
}
// 检查用户权限
func CheckUserPermission(operationType, operation string, userID int) (bool, string, error) {
	db := database2.GetConnection()
	var permissions string
	var userType int
	sql := `SELECT userType,IFNULL(permissions,"") FROM goku_admin WHERE userID = ?;`
	err := db.QueryRow(sql, userID).Scan(&userType, &permissions)
	if err != nil {
		return false, "[ERROR]This user does not exist!", err
	}
	if userType == 0 || userType == 1 {
		return true, "", nil
	}
	if permissions == "" {
		return false, "[ERROR]This user does not assigned permission", nil
	}
	permissionsMap := make(map[string]permissionsJson)
	err = json.Unmarshal([]byte(permissions), &permissionsMap)
	if err != nil {
		return false, "[ERROR]Fail to parse json!!", err
	}
	if value, ok := permissionsMap[operationType]; !ok {
		return false, "[ERROR]Operation type does not exist!", nil
	} else {
		if v, temp := value[operation]; !temp {
			return false, "[ERROR]Operation does not exist!!", nil
		} else {
			if !v {
				return false, "[ERROR]No permissions!", nil
			} else {
				return true, "", nil
			}
		}
	}
}


// 获取具有编辑权限的用户列表
func GetUserListWithPermission(operationType, operation string) (bool, []map[string]interface{}, error) {
	db := database2.GetConnection()
	sql := `SELECT userID,IF(remark IS NULL OR remark = "",loginCall,remark) as userName,userType,IFNULL(permissions,"") FROM goku_admin ORDER BY userType ASC;`
	rows, err := db.Query(sql)
	if err != nil {
		return false, make([]map[string]interface{}, 0), err
	}
	defer rows.Close()
	userList := make([]map[string]interface{}, 0)
	if _, err = rows.Columns(); err != nil {
		return false, make([]map[string]interface{}, 0), err
	} else {
		for rows.Next() {
			var (
				permissions string
				userType    int
				userID      int
				userName    string
			)
			err = rows.Scan(&userID, &userName, &userType, &permissions)
			if err != nil {
				return false, make([]map[string]interface{}, 0), err
			}
			if userType != 0 && userType != 1 {
				if permissions == "" {
					continue
				}
				permissionsMap := make(map[string]permissionsJson)
				err = json.Unmarshal([]byte(permissions), &permissionsMap)
				if err != nil {
					return false, make([]map[string]interface{}, 0), err
				}
				if value, ok := permissionsMap[operationType]; !ok {
					return false, make([]map[string]interface{}, 0), errors.New("[ERROR]Operation type does not exist!")
				} else {
					if v, temp := value[operation]; !temp {
						return false, make([]map[string]interface{}, 0), errors.New("[ERROR]Operation does not exist!")
					} else {
						if !v {
							continue
						}
					}
				}
			}
			userList = append(userList, map[string]interface{}{
				"userID":   userID,
				"userName": userName,
			})
		}
	}
	return true, userList, nil
}

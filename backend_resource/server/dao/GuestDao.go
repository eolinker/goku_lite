package dao

import (
	"goku-ce-1.0/dao/database"	
	"goku-ce-1.0/utils"
	"encoding/json"
	"strconv"
)

func Login(loginCall,loginPassword string) (bool,int){
	db := database.GetConnection()
	var userID int
	err := db.QueryRow("SELECT userID FROM eo_admin WHERE loginCall = ? AND loginPassword = ?;",loginCall,loginPassword).Scan(&userID)
	if err != nil{
		return false,0
	}
	return true,userID
}

func Register(loginCall,loginPassword string) (bool){
	db := database.GetConnection()
	var userID int
	err := db.QueryRow("SELECT userID FROM eo_admin WHERE loginCall = ?",loginCall).Scan(&userID)
	if err == nil{
		return false
	}else{
		stmt,err:= db.Prepare("INSERT INTO eo_admin (loginCall,loginPassword) VALUES (?,?);")
		defer stmt.Close()
		_, err = stmt.Exec(loginCall,loginPassword)
		if err != nil {
			return false
		} else{
			return true
		}
	}
	
}

func CheckUserNameExist(loginCall string) bool{
	db := database.GetConnection()
	var userID int
	err := db.QueryRow("SELECT userID FROM eo_admin WHERE loginCall = ?",loginCall).Scan(&userID)
	if err != nil{
		return false
	}else{
		return true
	}
}

// 写入登录信息 
func ConfirmLoginInfo(userID int,loginCall,userToken string) bool{
	redisConn,err := utils.GetRedisConnection()
	defer redisConn.Close()
	if err != nil{
		return false
	}
	redisInfo := map[string]interface{}{"userID":userID,"loginCall":loginCall,"userToken":userToken}
	redisString,err := json.Marshal(redisInfo)
	if err != nil{
		return false
	}
	_, err = redisConn.Do("hset","userToken",userToken,string(redisString[:])) 
	if err != nil{
		return false
	}
	return true
}

func CheckLogin(userID,userToken string) (bool) {
	redisConn,err := utils.GetRedisConnection()
	if err != nil{
		return false
	}
	defer redisConn.Close()
	var redisJson map[string]interface{}
	redisInfo,err := redisConn.Do("hget","userToken",userToken)
	if redisInfo == nil{
		return false
	}
	val := redisInfo.([]uint8)
	redisStr := []byte(string(val))
	
	err = json.Unmarshal(redisStr,&redisJson)
	if err != nil{
		return true
	}
	if uid,_ := strconv.Atoi(userID); uid != int(redisJson["userID"].(float64)){
		return false
	}else{
		return true
	}
}
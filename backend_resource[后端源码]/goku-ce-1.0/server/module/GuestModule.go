package module

import (
	"goku-ce-1.0/server/dao"
	"github.com/gin-gonic/gin"
	"goku-ce-1.0/utils"
)


func Login(loginCall,loginPassword string) (bool,int,string){
	flag,userID := dao.Login(loginCall,loginPassword)
	if flag{
		userToken := utils.GetHashKey(loginPassword)
		return true,userID,userToken
	}else{
		return false,0,""
	}
}

func CheckLogin(c *gin.Context) bool{
	userID,err := c.Cookie("userID")
	userToken,err := c.Cookie("userToken")
	if err != nil{
		return false
	}else{
		return dao.CheckLogin(userID,userToken)
	}
}


func Register(loginCall,loginPassword string) (bool){
	return dao.Register(loginCall,loginPassword)
}


func ConfirmLoginInfo(userID int,localCall,userToken string) bool{
	return dao.ConfirmLoginInfo(userID,localCall,userToken)
}
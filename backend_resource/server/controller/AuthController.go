package controller
import (
	"goku-ce-1.0/utils"
    "goku-ce-1.0/server/module"
	"github.com/gin-gonic/gin"
	"strconv"
)

func EditAuth(c *gin.Context){
	var userID int
	gatewayHashKey := c.PostForm("gatewayHashKey")
	
	if module.CheckLogin(c) == false{
		c.JSON(200,gin.H{"type":"guest","statusCode":"100000",})
		return 
	}else{
		result,_ := c.Request.Cookie("userID")
		userID,_ = strconv.Atoi(result.Value)
		flag := module.CheckGatewayPermission(gatewayHashKey,userID)
		if flag == false{
			c.JSON(200,gin.H{"statusCode":"100005","type":"gateway",})
			return 
		}
	}
	apiKey := c.PostForm("apiKey")
	userName := c.PostForm("userName")
	userPassword := c.PostForm("userPassword")

	authType,flag := utils.ConvertString(c.PostForm("authType"))
	if !flag{
		c.JSON(200,gin.H{"statusCode":"200001","type":"auth",})
		return
	}
	strategyID,flag := utils.ConvertString(c.PostForm("strategyID"))
	if !flag{
		c.JSON(200,gin.H{"statusCode":"200002","type":"auth",})
		return
	}
	if authType == 0{
		if len(userName) < 1 {
			c.JSON(200,gin.H{"type":"guest","statusCode":"200003"})
			return 
		}else if len(userPassword)<1 || len(userPassword)>255{
			c.JSON(200,gin.H{"type":"guest","statusCode":"200004"})
			return 
		}
	}
	flag = module.EditAuthMethod(authType,strategyID,gatewayHashKey,apiKey,userName,userPassword)
	if flag == false{
		c.JSON(200,gin.H{"statusCode":"200000","type":"auth",})
		return
	}else{
		c.JSON(200,gin.H{"type":"auth","statusCode":"000000"})
		return
	}
}


func GetAuthInfo(c *gin.Context){
	var userID int
	gatewayHashKey := c.PostForm("gatewayHashKey")
	
	if module.CheckLogin(c) == false{
		c.JSON(200,gin.H{"type":"guest","statusCode":"100000",})
		return 
	}else{
		result,_ := c.Request.Cookie("userID")
		userID,_ = strconv.Atoi(result.Value)
		flag := module.CheckGatewayPermission(gatewayHashKey,userID)
		if flag == false{
			c.JSON(200,gin.H{"statusCode":"100005","type":"gateway",})
			return 
		}
	}
	strategyID,flag := utils.ConvertString(c.PostForm("strategyID"))
	if !flag{
		c.JSON(200,gin.H{"statusCode":"200002","type":"auth",})
		return
	}
	flag,authInfo := module.GetAuthInfo(strategyID)
	if flag == false{
		c.JSON(200,gin.H{"statusCode":"200000","type":"auth",})
		return
	}else{
		c.JSON(200,gin.H{"type":"auth","statusCode":"000000","authInfo":authInfo})
		return
	}
}


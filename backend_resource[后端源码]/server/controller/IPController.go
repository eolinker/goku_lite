package controller
import (
	"goku-ce-1.0/utils"
    "goku-ce-1.0/server/module"
	"github.com/gin-gonic/gin"
	"strconv"
)


func EditIPList(c *gin.Context){
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
			c.JSON(200,gin.H{"statusCode":"100005","type":"guest",})
			return 
		}
	}
	var gatewayID int
	_,gatewayID = module.GetIDFromHashKey(gatewayHashKey)

	ipList := c.PostForm("ipList")
	chooseType,flag := utils.ConvertString(c.PostForm("chooseType"))
	if !flag{
		c.JSON(200,gin.H{"type":"ip","statusCode":"180001"})
		return
	}
	strategyID,flag := utils.ConvertString(c.PostForm("strategyID"))
	if !flag{
		c.JSON(200,gin.H{"type":"ip","statusCode":"180002"})
		return
	}
	flag = module.CheckStrategyPermission(gatewayID,strategyID)
	if !flag{
		c.JSON(200,gin.H{"type":"ip","statusCode":"180003"})
		return
	}else{
		flag = module.EditIPList(strategyID,chooseType,gatewayHashKey,ipList)
		if flag == true{
			c.JSON(200,gin.H{"type":"ip","statusCode":"000000"})
			return
		}else{
			c.JSON(200,gin.H{"type":"ip","statusCode":"180000"})
			return
		}
	}
}

// 获取IP名单列表
func GetIPInfo(c *gin.Context){
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
			c.JSON(200,gin.H{"statusCode":"100005","type":"guest",})
			return 
		}
	}
	var gatewayID int
	_,gatewayID = module.GetIDFromHashKey(gatewayHashKey)

	strategyID,flag := utils.ConvertString(c.PostForm("strategyID"))
	if !flag{
		c.JSON(200,gin.H{"type":"ip","statusCode":"180002"})
		return
	}
	flag = module.CheckStrategyPermission(gatewayID,strategyID)
	if !flag{
		c.JSON(200,gin.H{"type":"ip","statusCode":"180003"})
		return
	}else{
		flag,ipList := module.GetIPList(strategyID)
		if flag == true{
			c.JSON(200,gin.H{"type":"ip","statusCode":"000000","ipInfo":ipList})
			return
		}else{
			c.JSON(200,gin.H{"type":"ip","statusCode":"180000"})
		}
	}
}

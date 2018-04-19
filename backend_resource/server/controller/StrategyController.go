package controller
import (
	"github.com/gin-gonic/gin"
	"goku-ce-1.0/server/module"
	"strconv"
	"goku-ce-1.0/utils"
)

// 新增策略
func AddStrategy(c *gin.Context) {
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
	strategyName := c.PostForm("strategyName")
	strategyDesc := c.PostForm("strategyDesc")
	strategyNameLen := len([]rune(strategyName))
	strategyDescLen := len([]rune(strategyDesc))
	if strategyNameLen < 1 && strategyNameLen > 15{
		c.JSON(200,gin.H{"type":"strategy","statusCode":"190001"})
		return
	} else if strategyDescLen <1 && strategyDescLen > 50{
		c.JSON(200,gin.H{"type":"strategy","statusCode":"190002"})
		return
	} else {
		flag,id,strategyKey:= module.AddStrategy(strategyName,strategyDesc,gatewayID)
		if flag == true{
			c.JSON(200,gin.H{"type":"strategy","statusCode":"000000","strategyID":id,"strategyKey":strategyKey})
			return
		}else{
			c.JSON(200,gin.H{"type":"strategy","statusCode":"190000"})
			return
		}
	}
}

// 修改策略
func EditStrategy(c *gin.Context) {
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
	strategyName := c.PostForm("strategyName")
	strategyDesc := c.PostForm("strategyDesc")
	strategyNameLen := len([]rune(strategyName))
	strategyDescLen := len([]rune(strategyDesc))
	strategyID,flag := utils.ConvertString(c.PostForm("strategyID"))
	if flag == false{
		c.JSON(200,gin.H{"type":"strategy","statusCode":"190003"})
		return
	}else if strategyNameLen < 1 && strategyNameLen > 15{
		c.JSON(200,gin.H{"type":"strategy","statusCode":"190001"})
		return
	} else if strategyDescLen <1 && strategyDescLen > 50{
		c.JSON(200,gin.H{"type":"strategy","statusCode":"190002"})
		return
	} else {
		flag = module.EditStrategy(strategyName,strategyDesc,gatewayID,strategyID)
		if flag == true{
			c.JSON(200,gin.H{"type":"strategy","statusCode":"000000"})
			return
		}else{
			c.JSON(200,gin.H{"type":"strategy","statusCode":"190000"})
			return
		}
	}
}

// 删除策略
func DeleteStrategy(c *gin.Context) {
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
	if flag == false{
		c.JSON(200,gin.H{"type":"strategy","statusCode":"190003"})
		return
	} else {
		flag = module.DeleteStrategy(gatewayID,strategyID)
		if flag == true{
			c.JSON(200,gin.H{"type":"strategy","statusCode":"000000"})
			return
		}else{
			c.JSON(200,gin.H{"type":"strategy","statusCode":"190000"})
			return
		}
	}
}

// 获取策略列表
func GetStrategyList(c *gin.Context) {
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
	flag,strategyList := module.GetStrategyList(gatewayID)
	if flag == true{
		c.JSON(200,gin.H{"type":"strategy","statusCode":"000000","strategyList":strategyList,})
		return
	}else{
		c.JSON(200,gin.H{"type":"strategy","statusCode":"190000"})
		return
	}
}

// 获取简易策略组列表
func GetSimpleStrategyList(c *gin.Context) {
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
	flag,strategyList := module.GetSimpleStrategyList(gatewayID)
	if flag == true{
		c.JSON(200,gin.H{"type":"strategy","statusCode":"000000","strategyList":strategyList,})
		return
	}else{
		c.JSON(200,gin.H{"type":"strategy","statusCode":"190000"})
		return
	}
}

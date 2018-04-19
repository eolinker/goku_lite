package controller
import (
	_ "goku-ce-1.0/utils"
    "goku-ce-1.0/server/module"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddBackend(c *gin.Context){
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
	backendName := c.PostForm("backendName")
	backendURI := c.PostForm("backendURI")

	var gatewayID int
	_,gatewayID = module.GetIDFromHashKey(gatewayHashKey)

	flag,backendID := module.AddBackend(gatewayID,backendName,backendURI)
	if flag == false{
		c.JSON(200,gin.H{"statusCode":"140000","type":"gateway",})
		return
	}else{
		c.JSON(200,gin.H{"type":"gateway","statusCode":"000000","backendID":backendID})
		return
	}
}

func EditBackend(c *gin.Context){
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
	_,gatewayID := module.GetIDFromHashKey(gatewayHashKey)
	backendID := c.PostForm("backendID")
	backendName := c.PostForm("backendName")
	backendURI := c.PostForm("backendURI")
	id,_ := strconv.Atoi(backendID)
	flag := module.EditBackend(id,gatewayID,backendName,backendURI,gatewayHashKey)
	if flag == false{
		c.JSON(200,gin.H{"statusCode":"140000","type":"gateway",})
		return
	}else{
		c.JSON(200,gin.H{"type":"gateway","statusCode":"000000"})
		return
	}
}

func DeleteBackend(c *gin.Context){
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
	backendID := c.PostForm("backendID")
	var gatewayID int
	_,gatewayID = module.GetIDFromHashKey(gatewayHashKey)
	id,_ := strconv.Atoi(backendID)
	flag := module.DeleteBackend(gatewayID,id)
	if flag == false{
		c.JSON(200,gin.H{"statusCode":"140000","type":"gateway",})
		return
	}else{
		c.JSON(200,gin.H{"type":"gateway","statusCode":"000000"})
		return
	}
}

func GetBackendList(c *gin.Context){
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

	var gatewayID int
	_,gatewayID = module.GetIDFromHashKey(gatewayHashKey)
	flag,backendList := module.GetBackendList(gatewayID)
	if flag == false{
		c.JSON(200,gin.H{"statusCode":"140000","type":"gateway",})
		return
	}else{
		c.JSON(200,gin.H{"backendList":backendList,"statusCode":"000000","type":"gateway",})
		return
	}
}

func GetBackendInfo(c *gin.Context){
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
	backendID := c.PostForm("backendID")
	id,_ := strconv.Atoi(backendID)
	flag,backendInfo := module.GetBackendInfo(id)
	if flag == false{
		c.JSON(200,gin.H{"statusCode":"140000","type":"gateway",})
		return
	}else{
		c.JSON(200,gin.H{"backendInfo":backendInfo,"statusCode":"000000","type":"gateway",})
		return
	}
}


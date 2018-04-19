package controller
import (
	_ "goku-ce-1.0/utils"
	"goku-ce-1.0/server/module"
	"github.com/gin-gonic/gin"
	"strings"
	"strconv"
)

// 新增网关
func AddGateway(c *gin.Context){
	var userID int
	
	if module.CheckLogin(c) == false{
		c.JSON(200,gin.H{"type":"guest","statusCode":"100000",})
		return 
	}else{
		result,_ := c.Request.Cookie("userID")
		userID,_ = strconv.Atoi(result.Value)
	}
    gatewayName := c.PostForm("gatewayName")
	gatewayDesc := c.PostForm("gatewayDesc")
	gatewayAlias := c.PostForm("gatewayAlias")
	gatewayNameLen := strings.Count(gatewayName,"")-1
	if gatewayNameLen<1 || gatewayNameLen > 32 {
		c.JSON(200,gin.H{"type":"gateway","statusCode":"130001",})
		return 
	}
	flag,_ := module.CheckGatewayAliasIsExist(gatewayAlias)
	if flag{
			c.JSON(200,gin.H{"type":"gateway","statusCode":"130002",})
			return 
	}else{
		flag,gatewayHashkey := module.Addgateway(gatewayName,gatewayDesc,gatewayAlias,userID)
		if flag == false {
			c.JSON(200,gin.H{"type":"gateway","statusCode":"130000",})
			return 
		}else{
			c.JSON(200,gin.H{"type":"gateway","statusCode":"000000","gatewayHashKey":gatewayHashkey,})
			return 
		}
	}
}

// 修改网关
func EditGateway(c *gin.Context){
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
	gatewayName := c.PostForm("gatewayName")
	gatewayDesc := c.PostForm("gatewayDesc")
	gatewayAlias := c.PostForm("gatewayAlias")
	gatewayNameLen := strings.Count(gatewayName,"")-1
	if gatewayNameLen<1 || gatewayNameLen > 32 {
		c.JSON(200,gin.H{"type":"gateway","statusCode":"130001",})
		return 
	}
	flag,result := module.CheckGatewayAliasIsExist(gatewayAlias)
	if flag && result != gatewayHashKey{
			c.JSON(200,gin.H{"type":"gateway","statusCode":"130002",})
			return 
	}else{
		flag := module.EditGateway(gatewayName,gatewayAlias,gatewayDesc,gatewayHashKey,userID)
		if flag == false {
			c.JSON(200,gin.H{"type":"gateway","statusCode":"130000",})
			return 
		}else{
			c.JSON(200,gin.H{"type":"gateway","statusCode":"000000"})
			return 
		}
	}
}

// 删除网关
func DeleteGateway(c *gin.Context){
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
	flag := module.DeleteGateway(gatewayHashKey,userID)
	if flag == false {
		c.JSON(200,gin.H{"type":"gateway","statusCode":"130000",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"gateway","statusCode":"000000"})
		return 
	}
}

// 获取网关信息
func GetGatewayInfo(c *gin.Context){
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
	flag,result := module.GetGatewayInfo(gatewayHashKey,userID)
	if flag == false {
		c.JSON(200,gin.H{"type":"gateway","statusCode":"130000",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"gateway","statusCode":"000000","gatewayInfo":result})
		return 
	}
}


// 获取网关列表
func GetGatewayList(c *gin.Context){
	var userID int
	if module.CheckLogin(c) == false{
		c.JSON(200,gin.H{"type":"guest","statusCode":"100000",})
		return 
	}else{
		result,_ := c.Request.Cookie("userID")
		userID,_ = strconv.Atoi(result.Value)
	}
	flag,gatewayList := module.GetGatewayList(userID)
	if flag == false {
		c.JSON(200,gin.H{"type":"gateway","statusCode":"130000",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"gateway","statusCode":"000000","gatewayList":gatewayList,})
		return 
	}
}

// 查询网关别名是否存在
func CheckGatewayAliasIsExist(c *gin.Context) {
	gatewayAlias := c.PostForm("gatewayAlias")
	gatewayHashKey := c.PostForm("gatewayHashKey")
	flag,result := module.CheckGatewayAliasIsExist(gatewayAlias)
	if flag && result != gatewayHashKey{
		c.JSON(200,gin.H{"type":"gateway","statusCode":"000000",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"gateway","statusCode":"130000",})
		return 
	}
}



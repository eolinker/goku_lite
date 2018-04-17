package controller
import (
	"goku-ce-1.0/server/module"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 添加分组
func AddGroup(c *gin.Context) {
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
	parentGroupID := c.DefaultPostForm("parentGroupID","0")
	groupName := c.PostForm("groupName")
	pID,_ := strconv.Atoi(parentGroupID)
	flag,groupID := module.AddGroup(gatewayID,pID,groupName)
	if flag == false{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"150000",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"000000","groupID":groupID,})
		return 
	}
	
}

// 删除网关api分组
func DeleteGroup(c *gin.Context){
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
	
	groupID :=  c.PostForm("groupID")
	gID,_ := strconv.Atoi(groupID)
	flag := module.DeleteGroup(gID)
	if flag == false{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"150000",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"000000"})
		return 
	}
}

// 获取网关分组列表
func GetGroupList(c *gin.Context) {
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

	flag,groupList := module.GetGroupList(gatewayID)
	if flag == false{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"150000",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"000000","groupList":groupList,})
		return 
	}
}

// 修改分组信息
func EditGroup(c *gin.Context) {
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
	groupID := c.PostForm("groupID")
	parentGroupID := c.DefaultPostForm("parentGroupID","0")
	groupName := c.PostForm("groupName")
	gID,_ := strconv.Atoi(groupID)
	pID,_ := strconv.Atoi(parentGroupID)
	flag := module.EditGroup(gID,pID,groupName)
	if flag == false{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"150000",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"000000",})
		return 
	}
	
}

// 获取分组名称
func GetGroupName(c *gin.Context) {
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
	groupID := c.PostForm("groupID")
	gID,_ := strconv.Atoi(groupID)
	flag,groupName := module.GetGroupName(gID)
	if flag == false{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"150000",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"000000","groupName":groupName})
		return 
	}
}

// 获取网关分组列表
func GetGroupListByKeyword(c *gin.Context) {
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
	keyword := c.PostForm("keyword")
	flag,groupList := module.GetGroupListByKeyword(keyword,gatewayID)
	if flag == false{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"150000",})
		return 
	}else{
		c.JSON(200,gin.H{"type":"apiGroup","statusCode":"000000","groupList":groupList,})
		return 
	}
}


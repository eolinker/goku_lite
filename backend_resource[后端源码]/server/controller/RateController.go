package controller
import (
	"github.com/gin-gonic/gin"
	"goku-ce-1.0/server/module"
	"strconv"
	"goku-ce-1.0/utils"
	"regexp"
)

// 新增流量控制
func AddRateLimit(c *gin.Context) {
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
	viewType,flag := utils.ConvertString(c.PostForm("viewType"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210001"})
		return
	}
	strategyID,flag := utils.ConvertString(c.PostForm("strategyID"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210002"})
		return
	}
	intervalType,flag := utils.ConvertString(c.DefaultPostForm("intervalType","0"))
	if !flag && intervalType != 0{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210003"})
		return
	}
	limitCount,flag := utils.ConvertString(c.DefaultPostForm("limitCount","0"))
	if !flag && limitCount != 0{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210004"})
		return
	}
	priorityLevel,flag := utils.ConvertString(c.PostForm("priorityLevel"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210005"})
		return
	}

	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	if match, _ := regexp.MatchString(`^(20|21|22|23|[0-1]\d):[0-5]\d$`, startTime);match == false{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210006"})
		return 
	} else if match, _ := regexp.MatchString(`^(20|21|22|23|[0-1]\d):[0-5]\d$`, endTime);match == false{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210007"})
		return 
	} else{
		flag,id:= module.AddRateLimit(viewType,strategyID,intervalType,limitCount,priorityLevel,gatewayHashKey,startTime,endTime)
		if flag == true{
			c.JSON(200,gin.H{"type":"rateLimit","statusCode":"000000","limitID":id})
			return
		}else{
			c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210000"})
			return
		}
	}
}

// 编辑流量控制
func EditRateLimit(c *gin.Context) {
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
	limitID,flag := utils.ConvertString(c.PostForm("limitID"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210008"})
		return
	}
	viewType,flag := utils.ConvertString(c.PostForm("viewType"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210001"})
		return
	}
	strategyID,flag := utils.ConvertString(c.PostForm("strategyID"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210002"})
		return
	}
	intervalType,flag := utils.ConvertString(c.DefaultPostForm("intervalType","0"))
	if !flag && intervalType != 0{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210003"})
		return
	}
	limitCount,flag := utils.ConvertString(c.DefaultPostForm("limitCount","0"))
	if !flag && limitCount != 0{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210004"})
		return
	}
	priorityLevel,flag := utils.ConvertString(c.PostForm("priorityLevel"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210005"})
		return
	}

	startTime := c.PostForm("startTime")
	endTime := c.PostForm("endTime")
	if match, _ := regexp.MatchString(`^(20|21|22|23|[0-1]\d):[0-5]\d$`, startTime);match == false{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210006"})
		return 
	} else if match, _ := regexp.MatchString(`^(20|21|22|23|[0-1]\d):[0-5]\d$`, endTime);match == false{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210007"})
		return 
	} else{
		flag = module.EditRateLimit(strategyID,limitID,viewType,intervalType,limitCount,priorityLevel,gatewayHashKey,startTime,endTime)
		if flag == true{
			c.JSON(200,gin.H{"type":"rateLimit","statusCode":"000000"})
			return
		}else{
			c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210000"})
			return
		}
	}
}

// 删除流量控制
func DeleteRateLimit(c *gin.Context) {
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
	limitID,flag := utils.ConvertString(c.PostForm("limitID"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210008"})
		return
	}
	strategyID,flag := utils.ConvertString(c.PostForm("strategyID"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210002"})
		return
	}

	flag = module.DeleteRateLimit(strategyID,limitID,gatewayHashKey)
	if flag == true{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"000000"})
		return
	}else{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210000"})
		return
	}
}

// 获取流量控制信息
func GetRateLimitInfo(c *gin.Context) {
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
	limitID,flag := utils.ConvertString(c.PostForm("limitID"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210008"})
		return
	}

	flag,result := module.GetRateLimitInfo(limitID)
	if flag == true{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"000000","rateLimitInfo":result})
		return
	}else{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210000"})
		return
	}
}

// 获取流量控制列表
func GetRateLimitList(c *gin.Context) {
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
	strategyID,flag := utils.ConvertString(c.PostForm("strategyID"))
	if !flag{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210002"})
		return
	}

	flag,result := module.GetRateLimitList(strategyID)
	if flag == true{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"000000","rateLimitList":result})
		return
	}else{
		c.JSON(200,gin.H{"type":"rateLimit","statusCode":"210000"})
		return
	}
}
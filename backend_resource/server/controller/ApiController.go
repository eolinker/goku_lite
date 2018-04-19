package controller
import (
	"goku-ce-1.0/utils"
    "goku-ce-1.0/server/module"
	"github.com/gin-gonic/gin"
	"strconv"
	"encoding/json"
)

// 新增api
func AddApi(c *gin.Context){
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

	apiName := c.PostForm("apiName")
	gatewayRequestURI := c.PostForm("gatewayRequestURI")
	gatewayRequestPath := c.PostForm("gatewayRequestPath")
	backendRequestURI := c.PostForm("backendURI")
	backendRequestPath := c.PostForm("backendRequestPath")
	gatewayRequestBodyNote := c.PostForm("gatewayRequestBodyNote")

	groupID  := c.PostForm("groupID")
	gatewayProtocol := c.PostForm("gatewayProtocol")
	gatewayRequestType := c.PostForm("gatewayRequestType")
	backendProtocol := c.PostForm("backendProtocol")
	backendRequestType := c.PostForm("backendRequestType")
	backendID := c.PostForm("backendID")
	isRequestBody := c.PostForm("isRequestBody")

	gatewayRequestParam := c.PostForm("gatewayRequestParam")
	constantResultParam := c.PostForm("constantResultParam")
	var requestParam []utils.GatewayParam
	json.Unmarshal([]byte(gatewayRequestParam),&requestParam)
	var resultParam []utils.ConstantMapping
	json.Unmarshal([]byte(constantResultParam),&resultParam)


	gID,_:=strconv.Atoi(groupID)
	gpl,_:=strconv.Atoi(gatewayProtocol)
	grt,_:=strconv.Atoi(gatewayRequestType)
	bpl,_:=strconv.Atoi(backendProtocol)
	brt,_:=strconv.Atoi(backendRequestType)
	bID,_:=strconv.Atoi(backendID)
	isBody,_:=strconv.Atoi(isRequestBody)

	flag,apiID := module.AddApi(gatewayHashKey,apiName,gatewayRequestURI,gatewayRequestPath,backendRequestURI,backendRequestPath,gatewayRequestBodyNote,gatewayID,gID,gpl,grt,bpl,brt,bID,isBody,requestParam,resultParam)
	if flag == true{
		c.JSON(200,gin.H{"statusCode":"000000","type":"api","apiID":apiID})
		return 
	}else{
		c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
		return 
	}
}

// 修改api
func EditApi(c *gin.Context){
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

	apiName := c.PostForm("apiName")
	gatewayRequestURI := c.PostForm("gatewayRequestURI")
	gatewayRequestPath := c.PostForm("gatewayRequestPath")
	backendRequestURI := c.PostForm("backendURI")
	backendRequestPath := c.PostForm("backendRequestPath")
	gatewayRequestBodyNote := c.PostForm("gatewayRequestBodyNote")

	apiID := c.PostForm("apiID")
	groupID  := c.PostForm("groupID")
	gatewayProtocol := c.PostForm("gatewayProtocol")
	gatewayRequestType := c.PostForm("gatewayRequestType")
	backendProtocol := c.PostForm("backendProtocol")
	backendRequestType := c.PostForm("backendRequestType")
	backendID := c.PostForm("backendID")
	isRequestBody := c.PostForm("isRequestBody")

	gatewayRequestParam := c.PostForm("gatewayRequestParam")
	constantResultParam := c.PostForm("constantResultParam")
	var requestParam []utils.GatewayParam
	json.Unmarshal([]byte(gatewayRequestParam),&requestParam)
	var resultParam []utils.ConstantMapping
	json.Unmarshal([]byte(constantResultParam),&resultParam)

	aID,_:=strconv.Atoi(apiID)
	gID,_:=strconv.Atoi(groupID)
	gpl,_:=strconv.Atoi(gatewayProtocol)
	grt,_:=strconv.Atoi(gatewayRequestType)
	bpl,_:=strconv.Atoi(backendProtocol)
	brt,_:=strconv.Atoi(backendRequestType)
	bID,_:=strconv.Atoi(backendID)
	isBody,_:=strconv.Atoi(isRequestBody)

	flag,id := module.EditApi(gatewayHashKey,apiName,gatewayRequestURI,gatewayRequestPath,backendRequestURI,backendRequestPath,gatewayRequestBodyNote,aID,gatewayID,gID,gpl,grt,bpl,brt,bID,isBody,requestParam,resultParam)
	if flag == true{
		c.JSON(200,gin.H{"statusCode":"000000","type":"api","apiID":id})
		return 
	}else{
		c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
		return 
	}
}

// 彻底删除Api
func DeleteApi(c *gin.Context) {
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
	apiID := c.PostForm("apiID")
	aID,_ := strconv.Atoi(apiID)
	flag := module.DeleteApi(aID,gatewayID,gatewayHashKey)
	if flag == true{
		c.JSON(200,gin.H{"statusCode":"000000","type":"api"})
		return 
	}else{
		c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
		return 
	}
}

// 获取api列表并按照名称排序
func GetApiListOrderByName(c *gin.Context) {
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
	flag,apiList := module.GetApiListOrderByName(gID)
	if flag == true{
		flag,gatewayInfo := module.GetSimpleGatewayInfo(gatewayHashKey)
		if flag {
			c.JSON(200,gin.H{"statusCode":"000000","type":"api","apiList":apiList,"gatewayInfo":gatewayInfo})
			return 
		}else{
			c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
			return 
		}
		
	}else{
		c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
		return 
	}
}

func GetApi(c *gin.Context) {
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
	
	apiID := c.PostForm("apiID")
	aID,_ := strconv.Atoi(apiID)
	flag,apiInfo := module.GetApi(aID)
	if flag == true{
		flag,gatewayInfo := module.GetSimpleGatewayInfo(gatewayHashKey)
		if flag {
			c.JSON(200,gin.H{"statusCode":"000000","type":"api","apiInfo":apiInfo,"gatewayInfo":gatewayInfo})
			return 
		}else{
			c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
			return 
		}
		return 
	}else{
		c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
		return 
	}

}

// 获取所有API列表并依据接口名称排序
func GetAllApiListOrderByName(c *gin.Context) {
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
	flag,apiList:= module.GetAllApiListOrderByName(gatewayID)
	if flag == true{
		flag,gatewayInfo := module.GetSimpleGatewayInfo(gatewayHashKey)
		if flag {
			c.JSON(200,gin.H{"statusCode":"000000","type":"api","apiList":apiList,"gatewayInfo":gatewayInfo})
			return 
		}else{
			c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
			return 
		}
	}else{
		c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
		return 
	}
}

//搜索api
func SearchApi(c *gin.Context) {
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
	tips := c.PostForm("tips")
	flag,apiList:= module.SearchApi(tips,gatewayID)
	if flag == true{
		c.JSON(200,gin.H{"statusCode":"000000","type":"api","apiList":apiList})
		return 
	}else{
		c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
		return 
	}
}

// 查重
func CheckGatewayURLIsExist(c *gin.Context) {
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
	gatewayURL := c.PostForm("gatewayRequestPath")
	flag := module.CheckGatewayURLIsExist(gatewayID,gatewayURL)
	if flag == true{
		c.JSON(200,gin.H{"statusCode":"000000","type":"api"})
		return 
	}else{
		c.JSON(200,gin.H{"statusCode":"190000","type":"api"})
		return 
	}
}	


package dao

import (
	"goku-ce-1.0/dao/database"	
	"goku-ce-1.0/utils"
	"strconv"
	"encoding/json"
)

// 新增api
func AddApi(gatewayHashKey,apiName,gatewayRequestURI,gatewayRequestPath,backendRequestURI,backendRequestPath,gatewayRequestBodyNote,apiCacheJson,redisCacheJson string,gatewayID,groupID,gatewayProtocol,gatewayRequestType,backendProtocol,backendRequestType,backendID,isRequestBody int,gatewayRequestParam []utils.GatewayParam,constantResultParam []utils.ConstantMapping) (bool,int){
	db := database.GetConnection()
	Tx,_ := db.Begin()
	res,err := Tx.Exec("INSERT INTO eo_gateway_api (eo_gateway_api.apiName,eo_gateway_api.gatewayID,eo_gateway_api.groupID,eo_gateway_api.gatewayProtocol,eo_gateway_api.gatewayRequestType,eo_gateway_api.gatewayRequestURI,eo_gateway_api.backendProtocol,eo_gateway_api.backendRequestType,eo_gateway_api.backendID,eo_gateway_api.backendRequestURI,eo_gateway_api.isRequestBody,eo_gateway_api.gatewayRequestBodyNote) VALUES (?,?,?,?,?,?,?,?,?,?,?,?);",apiName,gatewayID,groupID,gatewayProtocol,gatewayRequestType,gatewayRequestPath,backendProtocol,backendRequestType,backendID,backendRequestURI,isRequestBody,gatewayRequestBodyNote)
	if err != nil {
		Tx.Rollback()
		return false,0
	} 
	
	if err != nil {
		Tx.Rollback()
		return false,0
	} else{
		apiID,_ := res.LastInsertId()
		for i := 0;i<len(gatewayRequestParam);i++{
			Tx.Exec("INSERT INTO eo_gateway_api_request_param (eo_gateway_api_request_param.apiID,eo_gateway_api_request_param.gatewayParamPostion,eo_gateway_api_request_param.isNotNull,eo_gateway_api_request_param.paramType,eo_gateway_api_request_param.gatewayParamKey,eo_gateway_api_request_param.backendParamPosition,eo_gateway_api_request_param.backendParamKey) VALUES (?,?,?,?,?,?,?);",apiID,gatewayRequestParam[i].ParamPosition,gatewayRequestParam[i].IsNotNull,gatewayRequestParam[i].ParamType,gatewayRequestParam[i].ParamKey,gatewayRequestParam[i].BackendParamPosition,gatewayRequestParam[i].BackendParamKey)
		}

		//插入常量参数
		for i := 0;i<len(constantResultParam);i++{
			Tx.Exec("INSERT INTO eo_gateway_api_constant (eo_gateway_api_constant.apiID,eo_gateway_api_constant.backendParamPosition,eo_gateway_api_constant.paramKey,eo_gateway_api_constant.paramName,eo_gateway_api_constant.paramValue) VALUES (?,?,?,?,?);",apiID,constantResultParam[i].ParamPosition,constantResultParam[i].BackendParamKey,constantResultParam[i].ParamName,constantResultParam[i].ParamValue)
		}
		// 写api信息进缓存表
		_,err = Tx.Exec("INSERT INTO eo_gateway_api_cache (eo_gateway_api_cache.apiID,eo_gateway_api_cache.apiJson,eo_gateway_api_cache.redisJson,eo_gateway_api_cache.gatewayHashKey,eo_gateway_api_cache.gatewayID,eo_gateway_api_cache.groupID,eo_gateway_api_cache.path,eo_gateway_api_cache.backendID) VALUES (?,?,?,?,?,?,?,?);",apiID,apiCacheJson,redisCacheJson,gatewayHashKey,gatewayID,groupID,backendRequestURI,backendID)
		redisConn,err := utils.GetRedisConnection()
		defer redisConn.Close()
		if err != nil{
			Tx.Rollback()
			return false,0
		}
		var queryJson utils.QueryJson
		var redisJson utils.RedisCacheJson
		err = json.Unmarshal([]byte(redisCacheJson),&redisJson)
		if err != nil{
			Tx.Rollback()
			return false,0
		}
		
		queryJson.OperationType = "api"
		queryJson.Operation = "add"
		queryJson.Data = redisJson
		redisString,_ := json.Marshal(queryJson)
		_, err = redisConn.Do("rpush", "gatewayQueue", string(redisString[:]))  
		if err != nil{
			Tx.Rollback()
			return false,0
		}
		
		Tx.Commit()
		return true,int(apiID)
	}
}

// 修改api
func EditApi(gatewayHashKey,apiName,gatewayRequestURI,gatewayRequestPath,backendRequestURI,backendRequestPath,gatewayRequestBodyNote,apiCacheJson,redisCacheJson string,apiID,gatewayID,groupID,gatewayProtocol,gatewayRequestType,backendProtocol,backendRequestType,backendID,isRequestBody int,gatewayRequestParam []utils.GatewayParam,constantResultParam []utils.ConstantMapping) (bool,int){
	db := database.GetConnection()
	Tx,_ := db.Begin()
	_,err := Tx.Exec("UPDATE eo_gateway_api SET eo_gateway_api.apiName = ?,eo_gateway_api.gatewayProtocol = ?,eo_gateway_api.gatewayRequestType = ?,eo_gateway_api.gatewayRequestURI = ?,eo_gateway_api.gatewayRequestBodyNote = ?,eo_gateway_api.isRequestBody = ?,eo_gateway_api.backendID = ?,eo_gateway_api.backendProtocol = ?,eo_gateway_api.backendRequestType = ?,eo_gateway_api.backendRequestURI = ?,eo_gateway_api.groupID = ? WHERE eo_gateway_api.apiID = ? AND eo_gateway_api.gatewayID = ?;",apiName,gatewayProtocol,gatewayRequestType,gatewayRequestPath,gatewayRequestBodyNote,isRequestBody,backendID,backendProtocol,backendRequestType,backendRequestURI,groupID,apiID,gatewayID)
	if err != nil {
		Tx.Rollback()
		return false,0
	} else{
		 Tx.Exec("DELETE FROM eo_gateway_api_request_param WHERE eo_gateway_api_request_param.apiID = ?;",apiID)
		 Tx.Exec("DELETE FROM eo_gateway_api_constant WHERE eo_gateway_api_constant.apiID = ?;",apiID)
 
		for i := 0;i<len(gatewayRequestParam);i++{
			Tx.Exec("INSERT INTO eo_gateway_api_request_param (eo_gateway_api_request_param.apiID,eo_gateway_api_request_param.gatewayParamPostion,eo_gateway_api_request_param.isNotNull,eo_gateway_api_request_param.paramType,eo_gateway_api_request_param.gatewayParamKey,eo_gateway_api_request_param.backendParamPosition,eo_gateway_api_request_param.backendParamKey) VALUES (?,?,?,?,?,?,?);",apiID,gatewayRequestParam[i].ParamPosition,gatewayRequestParam[i].IsNotNull,gatewayRequestParam[i].ParamType,gatewayRequestParam[i].ParamKey,gatewayRequestParam[i].BackendParamPosition,gatewayRequestParam[i].BackendParamKey)
		}

		//插入常量参数
		for i := 0;i<len(constantResultParam);i++{
			Tx.Exec("INSERT INTO eo_gateway_api_constant (eo_gateway_api_constant.apiID,eo_gateway_api_constant.backendParamPosition,eo_gateway_api_constant.paramKey,eo_gateway_api_constant.paramName,eo_gateway_api_constant.paramValue) VALUES (?,?,?,?,?);",apiID,constantResultParam[i].ParamPosition,constantResultParam[i].BackendParamKey,constantResultParam[i].ParamName,constantResultParam[i].ParamValue)
		}
		// 写api信息进缓存表
		Tx.Exec("UPDATE eo_gateway_api_cache SET eo_gateway_api_cache.groupID = ?,eo_gateway_api_cache.path = ?,eo_gateway_api_cache.backendID = ?,eo_gateway_api_cache.apiJson = ?,eo_gateway_api_cache.redisJson = ? WHERE eo_gateway_api_cache.apiID = ? AND eo_gateway_api_cache.gatewayID = ?;",groupID,backendRequestURI,backendID,apiCacheJson,redisCacheJson,apiID,gatewayID)
		redisConn,err := utils.GetRedisConnection()
		defer redisConn.Close()
		if err != nil{
			Tx.Rollback()
			return false,0
		}
		flag,gatewayKey := getRedisApi(apiID)
		if flag == false{
			Tx.Rollback()
			return false,0
		}
		var queryJson utils.QueryJson
		var redisJson utils.RedisCacheJson
		err = json.Unmarshal([]byte(redisCacheJson),&redisJson)
		if err != nil{
			Tx.Rollback()
			return false,0
		}
		queryJson.OperationType = "api"
		queryJson.Operation = "update"
		queryJson.Data = redisJson
		redisString,_ := json.Marshal(queryJson)
		_, err = redisConn.Do("DEL", "apiInfo:" + gatewayHashKey + ":" + gatewayKey)  
		if err != nil{
			Tx.Rollback()
			return false,0
		}
		_, err = redisConn.Do("rpush", "gatewayQueue",string(redisString[:]))  
		if err != nil{
			Tx.Rollback()
			return false,0
		}
		Tx.Commit()
		return true,int(apiID)
	}
}

// 彻底删除Api
func DeleteApi(apiID,gatewayID int,gatewayHashKey string) bool{
	db := database.GetConnection()
	Tx,_ := db.Begin()

	//删除基础信息
	stmt,_ := Tx.Prepare("DELETE FROM eo_gateway_api WHERE eo_gateway_api.apiID = ?")
	stmt.Exec(apiID)

	//删除缓存信息
	stmt,_ = Tx.Prepare("DELETE FROM eo_gateway_api_cache WHERE eo_gateway_api_cache.apiID = ?;")
	stmt.Exec(apiID)

	//删除请求参数
	stmt,_ = Tx.Prepare("DELETE FROM eo_gateway_api_request_param WHERE eo_gateway_api_request_param.apiID = ?;")
	stmt.Exec(apiID)

	//删除常量参数
	stmt,_ = Tx.Prepare("DELETE FROM eo_gateway_api_constant WHERE eo_gateway_api_constant.apiID = ?;")
	stmt.Exec(apiID)

	redisConn,err := utils.GetRedisConnection()
	defer redisConn.Close()
	if err != nil{
		Tx.Rollback()
		return false
	}
	var operationData utils.OperationData
	operationData.ApiID = apiID
	operationData.GatewayID = gatewayID
	operationData.GatewayHashKey = gatewayHashKey
	var queryJson utils.QueryJson
	
	queryJson.OperationType = "api"
	queryJson.Operation = "delete"
	queryJson.Data = operationData
	redisString,_ := json.Marshal(queryJson)
	_, err = redisConn.Do("rpush", "gatewayQueue",string(redisString[:]))  
	if err != nil{
		Tx.Rollback()
		return false
	}
	Tx.Commit()
	return true
}

// 获取api列表并按照名称排序
func GetApiListOrderByName(groupID int) (bool,[]*utils.ApiInfo){
	db := database.GetConnection()
	// 获取多级分组列表
	rows,err := db.Query(`SELECT eo_gateway_api_group.groupID FROM eo_gateway_api_group WHERE eo_gateway_api_group.parentGroupID = ?;`,groupID)
	groupSql := "" + strconv.Itoa(groupID)
	defer rows.Close()
	//如果存在子分组,则拼接搜索的范围
	for rows.Next(){
		var childGroupID int
		rows.Scan(&childGroupID)
		groupSql += "," + strconv.Itoa(childGroupID)
	}

	rows,err = db.Query(`SELECT eo_gateway_api.apiID,eo_gateway_api.apiName,eo_gateway_api.groupID,eo_gateway_api.gatewayProtocol,eo_gateway_api.gatewayRequestType,eo_gateway_api.gatewayRequestURI FROM eo_gateway_api WHERE eo_gateway_api.groupID IN (?) ORDER BY eo_gateway_api.apiID DESC;`,groupSql)
	apiList := make([]*utils.ApiInfo,0)

	flag := true
	num :=0
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
    	return false,apiList
	} else {
		for rows.Next(){
			var apiInfo utils.ApiInfo
			err = rows.Scan(&apiInfo.ApiID,&apiInfo.ApiName,&apiInfo.GroupID,&apiInfo.GatewayProtocol,&apiInfo.GatewayRequestType,&apiInfo.GatewayRequestURI)
			if err != nil{
				flag = false
				break
			}
			apiList = append(apiList,&apiInfo)
			num +=1
		}
	}
	if num == 0{
		flag =false
	}
	
	return flag,apiList
}

func GetApi(apiID int) (bool,utils.ApiInfo){
	db := database.GetConnection()
	var apiJson string
	var apiInfo utils.ApiInfo
	sql := "SELECT eo_gateway_api_cache.apiID,eo_gateway_api_cache.groupID,eo_gateway_api_cache.apiJson,eo_gateway_api_group.parentGroupID FROM eo_gateway_api_cache INNER JOIN eo_gateway_api_group ON eo_gateway_api_cache.groupID = eo_gateway_api_group.groupID WHERE eo_gateway_api_cache.apiID = ?;"
	err := db.QueryRow(sql,apiID).Scan(&apiInfo.ApiID,&apiInfo.GroupID,&apiJson,&apiInfo.ParentGroupID)
	if err != nil{
		return false,apiInfo
	}else{
		json.Unmarshal([]byte(apiJson),&apiInfo.ApiJson)
		return true,apiInfo
	}
}

// 获取所有API列表并依据接口名称排序
func GetAllApiListOrderByName(gatewayID int) (bool,[]*utils.ApiInfo){
	db := database.GetConnection()
	rows,err :=db.Query(`SELECT eo_gateway_api.apiID,eo_gateway_api.apiName,eo_gateway_api.groupID,eo_gateway_api.gatewayProtocol,eo_gateway_api.gatewayRequestType,eo_gateway_api.gatewayRequestURI FROM eo_gateway_api WHERE eo_gateway_api.gatewayID = ?;`,gatewayID)
	apiList := make([]*utils.ApiInfo,0)

	flag := true
	num :=0
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
    	return false,apiList
	} else {
		for rows.Next(){
			var apiInfo utils.ApiInfo
			err = rows.Scan(&apiInfo.ApiID,&apiInfo.ApiName,&apiInfo.GroupID,&apiInfo.GatewayProtocol,&apiInfo.GatewayRequestType,&apiInfo.GatewayRequestURI)
			if err != nil{
				flag = false
				break
			}
			apiList = append(apiList,&apiInfo)
			num +=1
		}
	}
	if num == 0{
		flag =false
	}
	
	return flag,apiList
}

//搜索api
func SearchApi(tips string,gatewayID int) (bool,[]*utils.ApiInfo){
	db := database.GetConnection()
	rows,err :=db.Query(`SELECT eo_gateway_api.apiID,eo_gateway_api.apiName,eo_gateway_api.groupID,eo_gateway_api.gatewayProtocol,eo_gateway_api.gatewayRequestType,eo_gateway_api.gatewayRequestURI FROM eo_gateway_api WHERE eo_gateway_api.gatewayID = ? AND (eo_gateway_api.apiName LIKE ? OR eo_gateway_api.gatewayRequestURI LIKE ?) ORDER BY eo_gateway_api.apiName;`,gatewayID,"%" + tips + "%","%" + tips + "%")
	apiList := make([]*utils.ApiInfo,0)

	flag := true
	num :=0
	for rows.Next(){
		var apiInfo utils.ApiInfo
		err = rows.Scan(&apiInfo.ApiID,&apiInfo.ApiName,&apiInfo.GroupID,&apiInfo.GatewayProtocol,&apiInfo.GatewayRequestType,&apiInfo.GatewayRequestURI)
		if err != nil{
			flag = false
			break
		}
		apiList = append(apiList,&apiInfo)
		num +=1
	}
	if num == 0{
		flag =false
	}
	
	return flag,apiList
}

// 获取网关接口信息，方便更新网关服务器的Redis数据
func GetRedisApiList(gatewayID int) (bool,[]*utils.ApiInfo){
	db := database.GetConnection()
	rows,err :=db.Query(`SELECT eo_gateway_api.gatewayProtocol,eo_gateway_api.gatewayRequestType,eo_gateway_api.gatewayRequestURI FROM eo_gateway_api WHERE eo_gateway_api.gatewayID = ?;`,gatewayID)
	apiList := make([]*utils.ApiInfo,0)

	flag := true
	num :=0
	//延时关闭Rows
	defer rows.Close()
	//获取记录列
	if _, err = rows.Columns(); err != nil {
		return false,apiList
	} else {
		for rows.Next(){
			var apiInfo utils.ApiInfo
			err = rows.Scan(&apiInfo.GatewayProtocol,&apiInfo.GatewayRequestType,&apiInfo.GatewayRequestURI)
			if err != nil{
				flag = false
				break
			}
			apiList = append(apiList,&apiInfo)
			num +=1
		}
	}
	if num == 0{
		flag =false
	}
	
	return flag,apiList
}

// 获取接口的核心信息
func GetRedisApi(apiID int) (bool,utils.ApiInfo){
	db := database.GetConnection()
	var apiInfo utils.ApiInfo

	sql := "SELECT eo_gateway_api.gatewayProtocol,eo_gateway_api.gatewayRequestType,eo_gateway_api.gatewayRequestURI FROM eo_gateway_api WHERE eo_gateway_api.apiID = ?;"
	err := db.QueryRow(sql,apiID).Scan(&apiInfo.GatewayProtocol,&apiInfo.GatewayRequestType,&apiInfo.GatewayRequestURI)
	if err != nil{
		return false,apiInfo
	}else{
		return true,apiInfo
	}
}

// 查重
func CheckGatewayURLIsExist(gatewayID int,gatewayURI string) (bool,int){
	db := database.GetConnection()
	var apiID int
	sql := "SELECT apiID FROM eo_gateway_api WHERE gatewayID = ? AND gatewayRequestURI = ?;"
	err := db.QueryRow(sql,gatewayID,gatewayURI).Scan(&apiID)
	if err != nil{
		return false,apiID
	}else{
		return true,apiID
	}
}

func getRedisApi(apiID int) (bool,string){
	db := database.GetConnection()
	var gatewayProtocol,gatewayRequestType,gatewayRequestURI string
	sql := "SELECT eo_gateway_api.gatewayProtocol,eo_gateway_api.gatewayRequestType,eo_gateway_api.gatewayRequestURI FROM eo_gateway_api WHERE eo_gateway_api.apiID = ?;"
	err := db.QueryRow(sql,apiID).Scan(&gatewayProtocol,&gatewayRequestType,&gatewayRequestURI)
	if err != nil{
		return false,""
	}else{
		return true,gatewayProtocol+":"+gatewayRequestType+":"+gatewayRequestURI
	}
}
// /**
// 		 * 获取接口的核心信息
// 		 */
// 		 public function getRedisApi(&$api_id)
// 		 {
// 			 $db = getDatabase();
// 			 $result = $db -> prepareExecute('SELECT eo_gateway_api.gatewayProtocol,eo_gateway_api.gatewayRequestType,eo_gateway_api.gatewayRequestURI FROM eo_gateway_api WHERE eo_gateway_api.apiID = ?;', array($api_id));
 
// 			 if (empty($result))
// 				 return FALSE;
// 			 else
// 				 return $result['gatewayProtocol'] . ':' . $result['gatewayRequestType'] . ':' . $result['gatewayRequestURI'];
// 		 }

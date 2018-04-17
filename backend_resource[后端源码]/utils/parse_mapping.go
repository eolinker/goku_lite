package utils

import (
	"encoding/json"
	"strconv"
)

var (
	TypeMap = map[string]string{"0": "text", "1": "file", "2": "json", "3": "int", "4": "float", "5": "double",
		"6": "date", "7": "datetime", "8": "boolean", "9": "byte", "10": "short", "11": "long"}

	PositionMap = map[string]string{"0": "header", "1": "body", "2": "query"}

	MethodIndicator = map[string]string{"0": "POST", "1": "GET", "2": "PUT", "3": "DELETE", "4": "HEAD",
		"5": "OPTIONS", "6": "PATCH"}
)

type requestMapping struct {
	GwParamKey           string `json:"gatewayParamKey"`
	ParamType            string `json:"paramType"`
	BackendParamPosition string `json:"backendParamPosition"`
	IsNotNull            string `json:"isNotNull"`
	GwParamPosition      string `json:"gatewayParamPosition"`
	BackendParamKey      string `json:"backendParamKey"`
}

type QueryJson struct {
	OperationType string `json:"type"`
	Operation string `json:"operation"`
	Data interface{} `json:"data"`
}

type ProjectInfo struct {
	ProjectID int `json:"projectID"`
	ProjectName string `json:"projectName"`
}

type constantMapping struct {
	ParamValue      string `json:"paramValue"`
	ParamPosition   string `json:"paramPosition"`
	ParamName       string `json:"paramName"`
	BackendParamKey string `json:"backendParamKey"`
}

type OperationData struct{
	ApiID int `json:"apiID,omitempty"`
	GatewayID int `json:"gatewayID"`
	GatewayHashKey string `json:"gatewayHashKey"`
	Token string `json:"token,omitempty"`
	GatewayAlias string `json:"gatewayAlias,omitempty"`
}

type GatewayInfo struct{
	GatewayID int `json:"gatewayID"`
	GatewayName string `json:"gatewayName"`
	GatewayDesc string `json:"gatewayDesc"`
	GatewayStatus string `json:"gatewayStatus,omitempty"`
	ProductType string `json:"productType,omitempty"`
	UpdateTime	string `json:"updateTime"`
	GatewayHashKey string `json:"gatewayHashKey"`
	Token string `json:"token,omitempty"`
	GatewayPort string `json:"gatewayPort,omitempty"`
	GatewayAlias	string	`json:"gatewayAlias,omitempty"`
}

type BackendInfo struct{
	BackendID int `json:"backendID"`
	BackendName string `json:"backendName"`
	BackendURI string `json:"backendURI"`
}

type GroupInfo struct{
	GroupID int `json:"groupID"`
	GroupName string `json:"groupName"`
	ChildGroupList []*ChildGroupInfo `json:"childGroupList"`
}

type ChildGroupInfo struct{
	GroupID int `json:"groupID"`
	GroupName string `json:"groupName"`
}

type ApiInfo struct{
	ApiID int `json:"apiID"`
	ApiName string `json:"apiName,omitempty"`
	GroupID int `json:"groupID"`
	GatewayProtocol int `json:"gatewayProtocol"`
	GatewayRequestType int `json:"gatewayRequestType"`
	GatewayRequestURI string `json:"gatewayRequestURI"`
	ApiJson ApiCacheJson `json:"apiJson,omitempty"`
	ParentGroupID int `json:"parentGroupID"`
	Port string `json:"gatewayPort,omitempty"`
}

type ApiJson struct {
	RequestParams      []requestMapping  `json:"requestParams"`
	ConstantParams     []constantMapping `json:"constantParams"`
	BackendProtocol    int            `json:"backendProtocol"`
	BackendURI         string            `json:"backendURI"`
	BackendPath        string            `json:"backendPath"`
	BackendRequestType int            `json:"backendRequestType"`
	IsRequestBody      int            `json:"isRequestBody"`
	GatewayHashKey     string            `json:"gatewayHashKey"`
}

type MessageInfo struct{
	MsgID 				int 			`json:"msgID"`
	MsgType 			int 			`json:"msgType"`
	Msg					string			`json:"msg"`
	Summary				string			`json:"summary"`
	MsgSendTime			string			`json:"msgSendTime"`
	IsRead				int				`json:"isRead"`
}
type ConfigureInfo struct{
	MysqlUserName		string			`json:"mysql_username"`
	MysqlPassword		string			`json:"mysql_password"`
	MysqlHost			string			`json:"mysql_host"`
	MysqlPort			string			`json:"mysql_port"`
	MysqlDBName			string			`json:"mysql_dbname"`
	RedisDB				string			`json:"redis_db"`
	RedisHost			string			`json:"redis_host"`
	RedisPort			string			`json:"redis_port"`
	RedisPassword		string			`json:"redis_password"`
	GatewayPort			string			`json:"eotest_port"`
	DayVisitLimit		string			`json:"day_visit_limit"`
	MinuteVisitLimit	string			`json:"minute_visit_limit"`
	DayThroughputLimit	string			`json:"day_throughput_limit"`
	IPMinuteVisitLimit	string			`json:"ip_minute_visit_limit"`
}

type RedisCacheJson struct {
	RequestParams      []GatewayParam  `json:"requestParams"`
	ConstantParams     []ConstantMapping `json:"constantParams"`
	BackendProtocol    int            `json:"backendProtocol"`
	BackendURI         string            `json:"backendURI"`
	BackendPath        string            `json:"backendPath"`
	BackendRequestType int            `json:"backendRequestType"`
	IsRequestBody      int            `json:"isRequestBody"`
	GatewayHashKey     string            `json:"gatewayHashKey"`
	GatewayID		   int					 `json:"gatewayID,omitempty"`
	GatewayRequestPath	string				 `json:"gatewayRequestPath,omitempty"`
}
type GatewayParamMapping struct {
	ParamType            string `json:"paramType"`
	IsNotNull            bool   `json:"isNotNull"`
	ParamPosition        string `json:"paramPosition"`
	BackendParamPosition string `json:"backendParamPosition"`
	ParamKey             string `json:"paramKey"`
	BackendParamKey      string `json:"backendParamKey"`
}

type GatewayParam struct {
	ParamType            string `json:"paramType"`
	IsNotNull            string   `json:"isNotNull"`
	ParamPosition        string `json:"gatewayParamPosition"`
	BackendParamPosition string `json:"backendParamPosition"`
	ParamKey             string `json:"gatewayParamKey"`
	BackendParamKey      string `json:"backendParamKey"`
}

type ConstantMapping struct {
	ParamValue      string `json:"paramValue"`
	ParamPosition   string `json:"paramPosition"`
	ParamName       string `json:"paramName"`
	BackendParamKey string `json:"backendParamKey"`
}

type MappingInfo struct {
	RequestParams      []GatewayParamMapping `json:"requestParams,omitempty"`
	ConstantParams     []ConstantMapping     `json:"constantParams,omitempty"`
	BackendProtocol    string                `json:"backendProtocol"`
	BackendURI         string                `json:"backendURI"`
	BackendPath        string                `json:"backendPath"`
	BackendRequestType string                `json:"backendRequestType"`
	IsRequestBody      bool                  `json:"isRequestBody"`
	StrategyKey		   string				 `json:"strategyKey"`
	GatewayHashKey     string                `json:"gatewayHashKey"`
	ApiID              int                   `json:"apiID"`
}

type ApiCacheJson struct{
	ApiName string `json:"apiName"`
	GatewayHashKey string `json:"gatewayHashKey"`
	GatewayProtocol int `json:"gatewayProtocol"`
	GatewayRequestType int `json:"gatewayRequestType"`
	GatewayRequestPath string `json:"gatewayRequestPath"`
	BackendProtocol int `json:"backendProtocol"`
	BackendRequestType int `json:"backendRequestType"`
	BackendID	int `json:"backendID"`
	BackendURI string `json:"backendURI"`
	BackendPath string `json:"backendPath"`
	IsRequestBody int `json:"isRequestBody"`
	GatewayRequestBodyNote string `json:"gatewayRequestBodyNote"`
	RequestParams []GatewayParam `json:"requestParams"`
	ConstantParams     []ConstantMapping     `json:"constantParams"` 
}

type IPListInfo struct {
	IPList     []string `json:"ipList"`
	ChooseType int      `json:"chooseType"`
}

type IPList struct{
	BlackList string `json:"blackList"`
	WhiteList string `json:"whiteList"`
	ChooseType int `json:"chooseType"`
}
func (info MappingInfo) String() string {
	rawInfo, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	return string(rawInfo)
}

func (info IPListInfo) String() string {
	rawInfo, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	return string(rawInfo)
}

func ParseDBJson(str,path string) *MappingInfo {
	
	var apis ApiJson
	
	err := json.Unmarshal([]byte(str), &apis)
	
	if err != nil {
		return nil
	}
	info := &MappingInfo{}
	for _, mapping := range apis.RequestParams {
		var param GatewayParamMapping
		param.ParamType = TypeMap[mapping.ParamType]
		param.IsNotNull = mapping.IsNotNull != "0"
		param.ParamPosition = PositionMap[mapping.GwParamPosition]
		param.BackendParamPosition = PositionMap[mapping.BackendParamPosition]
		param.ParamKey = mapping.GwParamKey
		param.BackendParamKey = mapping.BackendParamKey
		if param.ParamType == "" || param.ParamPosition == "" || param.BackendParamPosition == "" {
			continue
		}
		info.RequestParams = append(info.RequestParams, param)
	}

	for _, mapping := range apis.ConstantParams {
		var param ConstantMapping
		param.ParamValue = mapping.ParamValue
		param.ParamPosition = PositionMap[mapping.ParamPosition]
		param.ParamName = mapping.ParamName
		param.BackendParamKey = mapping.BackendParamKey
		if param.ParamPosition == "" {
			continue
		}
		info.ConstantParams = append(info.ConstantParams, param)
	}
	info.BackendProtocol = strconv.Itoa(apis.BackendProtocol)
	info.BackendURI = path
	info.BackendPath = apis.BackendPath
	info.BackendRequestType = MethodIndicator[strconv.Itoa(apis.BackendRequestType)]
	info.IsRequestBody = strconv.Itoa(apis.IsRequestBody) != "0"
	info.GatewayHashKey = apis.GatewayHashKey
	return info
}

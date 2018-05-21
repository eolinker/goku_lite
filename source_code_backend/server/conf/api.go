package conf

import (
	"strings"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"sort"
	"time"
)

type Api struct {
	ApiList					[]*ApiInfo					`json:"apis" yaml:"apis"`
}

type ApiInfo struct {
	ApiID					int						`json:"apiID" yaml:"api_id"`
	ApiName					string					`json:"apiName" yaml:"api_name"`
	GroupID					int						`json:"groupID" yaml:"group_id"`
	RequestURL				string					`json:"requestURL" yaml:"request_url"`
	RequestMethod			[]string				`json:"requestMethod" yaml:"request_method"`
	BackendID				int						`json:"backendID" yaml:"backend_id"`
	ProxyURL				string					`json:"proxyURL" yaml:"proxy_url"`
	ProxyMethod				string					`json:"proxyMethod" yaml:"proxy_method"`
	IsRaw					bool					`json:"isRaw" yaml:"is_raw"`
	ProxyParams				[]*Param				`json:"proxyParams" yaml:"proxy_params"`
	ConstantParams			[]*ConstantParam		`json:"constantParams" yaml:"constant_params"`
	Follow					bool					`json:"follow" yaml:"follow"`
	UpdateTime				string					`json:"updateTime" yaml:"update_time"`
	CreateTime				string					`json:"createTime" yaml:"createTime"`						
}

type Param struct {
	Key						string					`json:"key" yaml:"key"`
	KeyPosition				string					`json:"keyPosition" yaml:"key_position"`
	NotEmpty				bool					`json:"notEmpty" yaml:"not_empty"`
	ProxyKey				string					`json:"proxyKey" yaml:"proxy_key"`
	ProxyKeyPosition		string					`json:"proxyKeyPosition" yaml:"proxy_key_position"`
}

type ConstantParam struct {
	Position				string					`json:"position" yaml:"position"`
	Key						string					`json:"key" yaml:"key"`
	Value					string					`json:"value" yaml:"value"`
}

type ApiSlice []map[string]interface{}
 
func (a ApiSlice) Len() int {    // 重写 Len() 方法
    return len(a)
}
func (a ApiSlice) Swap(i, j int){     // 重写 Swap() 方法
    a[i], a[j] = a[j], a[i]
}
func (a ApiSlice) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
	t1, t1Err := time.Parse("2006-01-02 15:04:05", a[j]["updateTime"].(string))
	t2, t2Err := time.Parse("2006-01-02 15:04:05", a[i]["updateTime"].(string))

	if t1Err == nil && t2Err != nil {
		return true
	} else if t1Err == nil && t2Err == nil{
		if t1.Before(t2) {
			return false
		} else {
			return true
		}
	} else if t1Err != nil && t2Err == nil {
		return false
	} else {
		s := []string{a[i]["apiName"].(string),a[j]["apiName"].(string)}
		sort.Strings(s)
		if s[0] == a[i]["apiName"].(string) {
			return false
		}else {
			return true
		}
	}
}

// 读入接口信息
func ParseApiInfo(path string) ([]*ApiInfo,map[int]*ApiInfo,[]map[string]interface{}) {
	apiInfo := make(map[int]*ApiInfo)
	mapApiList := make([]map[string]interface{},0)
	apiList := make([]*ApiInfo,0)
	var api Api
	content,err := ioutil.ReadFile(path)
	if err != nil {
		return apiList,apiInfo,mapApiList
	}

	err = yaml.Unmarshal(content,&api)
	if err != nil {
		panic(err)
	}

	
	if len(api.ApiList) != 0 {
		apiList = api.ApiList
	}

	maxID := 0
	for _,a := range api.ApiList {
		if a.ApiID > maxID {
			maxID = a.ApiID
		}
	}

	for _,a := range api.ApiList {
		apiID := a.ApiID
		if apiID == 0 {
			apiID = maxID + 1 
		}
		value := map[string]interface{}{
			"apiID":apiID,
			"apiName":a.ApiName,
			"groupID":a.GroupID,
			"requestURL":a.RequestURL,
			"requestMethod":strings.Join(a.RequestMethod,","),
			"updateTime":a.UpdateTime,
			"follow": a.Follow,
		}
		mapApiList = append(mapApiList,value)
		apiInfo[apiID] = a
		maxID += 1
	}
	sort.Sort(sort.Reverse(ApiSlice(mapApiList)))
	return apiList,apiInfo,mapApiList
}

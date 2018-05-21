package conf

import (
	"sort"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
	"goku-ce/conf"
	"time"
)

type GatewayInfo struct {
	GatewayName				string					`json:"gatewayName" yaml:"gateway_name"`
	GatewayAlias			string					`json:"gatewayAlias" yaml:"gateway_alias"`
	GatewayStatus			string					`json:"gatewayStatus" yaml:"gateway_status"`
	ApiConfPath				string					`json:"apiConfPath" yaml:"api_conf_path"`
	ApiGroupConfPath		string					`json:"apiGroupConfPath" yaml:"api_group_conf_path"`
	StrategyConfPath		string					`json:"strategyConfPath" yaml:"strategy_conf_path"`
	BackendConfPath			string					`json:"backendConfPath" yaml:"backend_conf_path"`
	IPLimitType				string					`json:"ipLimitType,omitempty" yaml:"ip_limit_type,omitempty"`
	IPWhiteList				[]string				`json:"ipWhiteList,omitempty" yaml:"ip_white_list,omitempty"`
	IPBlackList				[]string				`json:"ipBlackList,omitempty" yaml:"ip_black_list,omitempty"`
	UpdateTime				string					`json:"updateTime,omitempty" yaml:"update_time,omitempty"`	
	CreateTime				string					`json:"createTime,omitempty" yaml:"create_time,omitempty"`
}

// 按照 Person.Age 从大到小排序
type GatewaySlice []map[string]interface{}
 
func (a GatewaySlice) Len() int {    // 重写 Len() 方法
    return len(a)
}
func (a GatewaySlice) Swap(i, j int){     // 重写 Swap() 方法
    a[i], a[j] = a[j], a[i]
}
func (a GatewaySlice) Less(i, j int) bool {    // 重写 Less() 方法， 从大到小排序
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
		s := []string{a[i]["gatewayAlias"].(string),a[j]["gatewayAlias"].(string)}
		sort.Strings(s)
		if s[0] == a[i]["gatewayAlias"].(string) {
			return false
		}else {
			return true
		}
	}
}

// 读入网关信息
func ParseGatewayInfo(path string) (map[string]*GatewayInfo){
	gatewayInfo := make(map[string]*GatewayInfo)
	dirPath,err := conf.GetDir(path)
	if err == nil {
		PthSep := string(os.PathSeparator)
		for _,p := range dirPath {
			gateway := &GatewayInfo{}
			c,err := ioutil.ReadFile(p + PthSep + "gateway.conf")
			if err != nil {
				continue
			}
			err = yaml.Unmarshal(c,&gateway)
			if err != nil {
				continue
			}
			_,ok := gatewayInfo[gateway.GatewayAlias]
			if ok {
				continue
			}
			gatewayInfo[gateway.GatewayAlias] = gateway
		}
	}
	return gatewayInfo
}
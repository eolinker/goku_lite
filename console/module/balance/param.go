package balance

import (
	"encoding/json"

	entity "github.com/eolinker/goku-api-gateway/server/entity/balance-entity-service"
)

//Param 负载参数
type Param struct {
	Name          string `opt:"balanceName,require"`
	ServiceName   string `opt:"serviceName,require"`
	AppName       string `opt:"appName"`
	Static        string `opt:"static"`
	StaticCluster string `opt:"staticCluster"`
	Desc          string `opt:"balanceDesc"`
}

//Info 负载信息
type Info struct {
	Name          string            `json:"balanceName"`
	ServiceName   string            `json:"serviceName"`
	ServiceType   string            `json:"serviceType"`
	ServiceDriver string            `json:"serviceDriver"`
	AppName       string            `json:"appName"`
	Static        string            `json:"static"`
	StaticCluster map[string]string `json:"staticCluster"`
	Desc          string            `json:"balanceDesc"`
	CreateTime    string            `json:"createTime"`
	UpdateTime    string            `json:"updateTime"`
	CanDelete     int               `json:"canDelete"`
}

//ReadInfo 读取负载信息
func ReadInfo(balance *entity.Balance) *Info {
	info := &Info{
		Name:          balance.Name,
		ServiceName:   balance.ServiceName,
		ServiceType:   balance.ServiceType,
		ServiceDriver: balance.ServiceDriver,
		AppName:       balance.AppName,
		Static:        balance.Static,
		StaticCluster: nil,
		Desc:          balance.Desc,
		CreateTime:    balance.CreateTime,
		UpdateTime:    balance.UpdateTime,
		CanDelete:     balance.CanDelete,
	}
	json.Unmarshal([]byte(balance.StaticCluster), &info.StaticCluster)
	return info
}

package entity

import (
	"bytes"
	"strconv"
	"strings"
	"unicode"

	jsoniter "github.com/json-iterator/go"
)

//BalanceInfoEntity 负载信息实体
type BalanceInfoEntity struct {
	Name string
	Desc string

	DefaultConfig    string
	ClusterConfig    string
	OldVersionConfig string

	CreateTime string
	UpdateTime string
}

//BalanceInfo 负载信息
type BalanceInfo struct {
	Name string `json:"balanceName"`
	Desc string `json:"balanceDesc"`

	Default    *BalanceConfig            `json:"defaultConfig"`
	Cluster    map[string]*BalanceConfig `json:"clusterConfig"`
	CreateTime string                    `json:"createTime"`
	UpdateTime string                    `json:"updateTime"`
}

//BalanceConfig 负载配置
type BalanceConfig struct {
	DiscoveryID      int64                  `json:"serviceDiscoveryID"`
	ServiceName      string                 `json:"serviceName"`
	Servers          []*BalanceServerConfig `json:"-"`
	ServersConfig    []string               `json:"static"`
	ServersConfigOrg string                 `json:"staticOrg"`
}

//BalanceServerConfig 负载服务配置
type BalanceServerConfig struct {
	Server string `json:"server"`
	Weight int    `json:"weight"`
	Status string `json:"status"`
}
type _OldVersionBalanceInfo struct {
	BalanceConfig []*BalanceServerConfig `json:"loadBalancingServer"`
}

//
//func (info *BalanceInfo) Write() *BalanceInfoEntity {
//
//	ent := new(BalanceInfoEntity)
//	ent.Name = info.Name
//	ent.Desc = info.Name
//	ent.UpdateTime = info.UpdateTime
//	ent.CreateTime = info.CreateTime
//	ent.DefaultConfig, _ = jsoniter.MarshalToString(info.Default)
//	ent.ClusterConfig, _ = jsoniter.MarshalToString(info.Cluster)
//	return ent
//
//}

//Decode decode
func (ent *BalanceInfoEntity) Decode() (*BalanceInfo, error) {
	info := new(BalanceInfo)
	info.Name = ent.Name
	info.Desc = ent.Desc
	info.CreateTime = ent.CreateTime
	info.UpdateTime = ent.UpdateTime

	if ent.DefaultConfig != "" {
		info.Default = new(BalanceConfig)
		if err := jsoniter.UnmarshalFromString(ent.DefaultConfig, info.Default); err != nil {
			return nil, err
		}
		info.Default.Decode()

	} else {
		config, err := TryOld(ent.OldVersionConfig)
		if err != nil {
			return nil, err
		}

		info.Default = &BalanceConfig{
			DiscoveryID:   0,
			ServiceName:   "",
			Servers:       config,
			ServersConfig: FormatServers(config),
		}
		info.Default.ServersConfigOrg = strings.Join(info.Default.ServersConfig, ";")
	}

	if ent.ClusterConfig != "" {
		info.Cluster = make(map[string]*BalanceConfig)
		if err := jsoniter.UnmarshalFromString(ent.ClusterConfig, &info.Cluster); err != nil {
			return nil, err
		}
		for key := range info.Cluster {
			info.Cluster[key].Decode()
		}
	}
	return info, nil

}

//TryOld 解析旧配置
func TryOld(oldversionConfig string) ([]*BalanceServerConfig, error) {
	if oldversionConfig == "" {
		return nil, nil
	}

	old := new(_OldVersionBalanceInfo)
	err := jsoniter.UnmarshalFromString(oldversionConfig, old)

	if err != nil {
		return nil, err
	}
	return old.BalanceConfig, nil
	////Default := new(BalanceConfig)
	//Default.Servers = old.BalanceConfig
	//Default.Format()
	//Default.ServersConfigOrg = strings.Join(Default.ServersConfig,";")
	//return Default, err
}

//GetConfig 获取配置
func (info *BalanceInfo) GetConfig(clusterName string) *BalanceConfig {
	c, has := info.Cluster[clusterName]
	if !has || len(c.Servers) < 1 {
		return info.Default
	}
	// todo: 服务发现需要另外处理，暂时不接入
	return c
}

//FormatServers 格式化服务
func FormatServers(servers []*BalanceServerConfig) []string {
	if len(servers) == 0 {
		return nil
	}

	serversConfig := make([]string, 0, len(servers))
	buf := bytes.NewBuffer(make([]byte, 0, len(servers)*32))
	for _, node := range servers {
		if strings.TrimSpace(node.Server) == "" {
			continue
		}
		buf.WriteString(node.Server)
		buf.WriteString(" ")
		buf.WriteString(strconv.Itoa(node.Weight))
		if node.Status != "" {
			buf.WriteString(" ")
			buf.WriteString(node.Status)
		}
		serversConfig = append(serversConfig, buf.String())
		buf.Reset()
	}
	return serversConfig
}
func fields(str string) []string {

	words := strings.FieldsFunc(strings.Join(strings.Split(str, ";"), " ; "), func(r rune) bool {
		if unicode.IsSpace(r) {
			return true
		}

		return false
	})
	return words
}

//Decode decode配置
func (c *BalanceConfig) Decode() error {

	words := fields(c.ServersConfigOrg)

	s := make([]*BalanceServerConfig, 0, 5)

	var node *BalanceServerConfig
	index := 0
	for _, word := range words {
		if word == ";" {
			index = 0
			node = nil
			continue
		}
		l := len(word)
		value := word
		if word[l-1] == ';' {
			value = word[:l-1]
		}
		switch index {
		case 0:
			{
				node = new(BalanceServerConfig)
				node.Server = value
				s = append(s, node)
			}
		case 1:
			{
				weight, err := strconv.Atoi(value)
				if err != nil {
					return err
				}
				node.Weight = weight

			}
		case 2:
			{
				node.Status = value
			}
		}
		if node.Weight == 0 {
			node.Weight = 1
		}
		if word[l-1] == ';' {
			index = 0
			node = nil
		} else {
			index++
		}
	}

	if len(s) > 0 {
		c.Servers = s
	}
	return nil

}

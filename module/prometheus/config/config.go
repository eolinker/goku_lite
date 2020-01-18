package config

import (
	"encoding/json"

	"github.com/eolinker/goku-api-gateway/ksitigarbha"
)

//ModuleNameSpace 模块空间名称
const ModuleNameSpace = "diting.prometheus"
const moduleName = "Prometheus"

//Pattern 请求路径
const Pattern = "/prometheus/metrics"
const desc = "对接Prometheus"
const content = `[
        {
            "type": "line",
            "label":"监控数据读取：http://{{节点IP}}:{{节点管理地址的端口}}/prometheus/metrics ",
            "descript":"(节点的 [管理地址端口] 可在节点信息处查看)",
            "items":[]
        }
    ]`

var (
	mode []ksitigarbha.Model
)

func init() {
	err := json.Unmarshal([]byte(content), &mode)
	if err != nil {
		panic("init prometheus config error")
		return
	}
}

//PrometheusModule 配置
type PrometheusModule struct {
}

//GetNameSpace 获取命名空间
func (c *PrometheusModule) GetNameSpace() string {
	return ModuleNameSpace
}

//Encode encode
func (c *PrometheusModule) Encode(v interface{}) (string, error) {
	return "", nil
}

//GetDefaultConfig 获取默认配置
func (c *PrometheusModule) GetDefaultConfig() interface{} {
	return nil
}

//GetModel 获取模板
func (c *PrometheusModule) GetModel() []ksitigarbha.Model {
	return mode
}

//GetDesc 获取描述
func (c *PrometheusModule) GetDesc() string {
	return desc
}

//GetName 获取名称
func (c *PrometheusModule) GetName() string {
	return moduleName
}

//Decode decode
func (c *PrometheusModule) Decode(config string) (interface{}, error) {
	return nil, nil
}

//Register 模板注册
func Register() {
	ksitigarbha.Register(moduleName, new(PrometheusModule))
}

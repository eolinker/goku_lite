package entity

import (
	goku_plugin "github.com/eolinker/goku-plugin"
)

const (
	PluginTypeGateway  = 0
	PluginTypeStrategy = 1
	PluginTypeApi      = 2
)

type PluginInfo struct {
	Name      string
	Priority  int
	Config    string
	IsStop    bool
	Type      int
	UpdateTag string
}
type MapString map[string]string

type PluginFactoryHandler struct {
	Info    *PluginInfo
	Factory goku_plugin.PluginFactory

	//Config    string
	//UpdateTag string
}

type PluginHandlerExce struct {
	PluginObj *goku_plugin.PluginObj
	Name      string
	Priority  int
	IsStop    bool
}

type PluginSlice []*PluginHandlerExce

func (p PluginSlice) Len() int { // 重写 Len() 方法
	return len(p)
}
func (p PluginSlice) Swap(i, j int) { // 重写 Swap() 方法
	p[i], p[j] = p[j], p[i]
}
func (p PluginSlice) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return p[i].Priority < p[j].Priority
}

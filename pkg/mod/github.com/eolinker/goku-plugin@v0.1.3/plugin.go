package goku_plugin

type PluginObj struct {
	BeforeMatch PluginBeforeMatch
	Access      PluginAccess
	Proxy       PluginProxy
}

// 定义插件实现函数
type PluginFactory interface {
	Create(config string, clusterName string, updateTag string, strategyId string, apiId int) (*PluginObj, error)
}

type PluginBeforeMatch interface {
	BeforeMatch(ctx ContextBeforeMatch) (isContinue bool, e error)
}

type PluginAccess interface {
	Access(ctx ContextAccess) (isContinue bool, e error)
}

type PluginProxy interface {
	Proxy(ctx ContextProxy) (isContinue bool, e error)
}

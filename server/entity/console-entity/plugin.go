package entity

//Plugin 插件
type Plugin struct {
	PluginID     int    `json:"pluginID"`
	PluginName   string `json:"pluginName"`
	ChineseName  string `json:"chineseName,omitempty"`
	PluginStatus int    `json:"pluginStatus"`
	PluginIndex  int    `json:"pluginPriority"`
	PluginConfig string `json:"pluginConfig,omitempty"`
	PluginInfo   string `json:"pluginInfo,omitempty"`
	PluginType   int    `json:"pluginType"`
	// Official     bool   `json:"official"`
	Version    string `json:"version"`
	PluginDesc string `json:"pluginDesc"`
	IsStop     int    `json:"isStop"`
	IsCheck    int    `json:"isCheck"`
}

//PluginList 插件列表
type PluginList struct {
	PluginList []*PluginParams `json:"pluginList"`
}

//PluginParams 插件参数
type PluginParams struct {
	PluginName   string            `json:"pluginName"`
	PluginConfig string            `json:"pluginConfig"`
	PluginIndex  int               `json:"pluginPriority"`
	PluginInfo   map[string]string `json:"pluginInfo"`
}

//PluginSlice 插件切片
type PluginSlice []*Plugin

func (p PluginSlice) Len() int { // 重写 Len() 方法
	return len(p)
}
func (p PluginSlice) Swap(i, j int) { // 重写 Swap() 方法
	p[i], p[j] = p[j], p[i]
}
func (p PluginSlice) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return p[i].PluginIndex < p[j].PluginIndex
}

//ProxyCachingConf 代理缓存配置
type ProxyCachingConf struct {
	ResponseCodes  string `json:"responseCodes"`  //缓存条件：返回的HTTP状态码在该状态码列表中
	RequestMethods string `json:"requestMethods"` //缓存条件：请求的Method在该列表中
	ContentTypes   string `json:"contentTypes"`   //缓存条件：返回的Content-Type在该列表中
	CacheTTL       int    `json:"cacheTTL"`
	RedisHost      string `json:"redisHost"`
	RedisTimeout   int    `json:"redisTimeout"`
	RedisPort      string `json:"redisPort"`
	RedisPassword  string `json:"redisPassword"`
	RedisDatabase  int    `json:"redisDatabase"`
}

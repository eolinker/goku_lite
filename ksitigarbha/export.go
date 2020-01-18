package ksitigarbha

import (
	log "github.com/eolinker/goku-api-gateway/goku-log"
)

var mManager = newModelManager()

//GetMonitorModuleNames 获取监控模块名称列表
func GetMonitorModuleNames() []string {
	return mManager.getModuleNames()
}

//GetMonitorModuleModel 获取
func GetMonitorModuleModel(name string) (IModule, bool) {
	return mManager.getModuleModel(name)
}

//GetNameSpaceByName 获取namespace
func GetNameSpaceByName(name string) (namespace string, has bool) {
	return mManager.getNameSpace(name)

}

//Register 注册
func Register(name string, f IModule) {
	if f == nil {
		log.Panic("register ksitigarbha nil")
	}
	mManager.add(name, f)
}

//Close close
func Close(name string) {
	mManager.Close(name)
}

//Open open
func Open(name string, config string) {
	mManager.Open(name, config)
}

//HandlerConfig HandlerConfig
func HandlerConfig(name string, handler ConfigHandler) {
	mManager.Handler(name, handler)
}

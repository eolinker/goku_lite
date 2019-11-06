package ksitigarbha

import (
	"log"


)

// 这里只是在程序启动的时候会执行新增操作，所以不需要加锁
type modelManager struct {
	modules map[string]IModule
	names   []string
	namespaceNames map[string]string
}

func newModelManager() *modelManager {
	return &modelManager{
		modules:        make(map[string]IModule),
		names:          make([]string,0,5),
		namespaceNames: make(map[string]string),
	}
}

var mManager = newModelManager()

//GetMonitorModuleNames 获取监控模块名称列表
func GetMonitorModuleNames() []string {
	return mManager.getModuleNames()
}

//GetMonitorModuleModel 获取
func GetMonitorModuleModel(name string) (IModule ,bool){
	return mManager.getModuleModel(name)
}

//GetNameSpaceByName 获取namespace
func GetNameSpaceByName(name string) string {
	return mManager.getNameSpace(name)
}
//Register 注册
func Register(name string,f IModule) {
	if f==nil {
		log.Panic("register ksitigarbha nil")
	}
	mManager.add(name,f)

}

func (m *modelManager) add(name string,f IModule) {

	_,has:=m.modules[name]
	if has{
		log.Panic("register ksitigarbha duplicate name")
	}
	m.modules[name] =f
	m.names = append(m.names, f.GetName())
	m.namespaceNames[name] = f.GetNameSpace()
}

func (m *modelManager) getModuleNames() []string {
	return m.names
}

func (m *modelManager) getModuleModel(name string) (IModule ,bool){

	v, ok := m.modules[name]
	return v,ok
}

func (m *modelManager) isExisted(name string) bool {

	_, ok := m.modules[name]

	return ok
}
func (m *modelManager) getNameSpace(name string) string {
	return m.namespaceNames[name]
}
func (m *modelManager) getModuleCount() int {
	count := len(m.modules)
	return count
}

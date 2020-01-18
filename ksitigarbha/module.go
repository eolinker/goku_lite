package ksitigarbha

import (
	"log"


)

// 这里只是在程序启动的时候会执行新增操作，所以不需要加锁
type modelManager struct {
	modules map[string]IModule
	names   []string
	namespaceNames map[string]string
	handlers map[string][]ConfigHandler
}

func (m *modelManager) Handler(name string, handler ConfigHandler) {
	m.handlers[name] = append(m.handlers[name],handler)
}

func (m *modelManager) Close(name string) {
	namespace,has:= m.getNameSpace(name)
	if !has{
		return
	}
	hs:=m.handlers[name]
	for _,handler:=range hs{
		handler.OnClose(namespace,name)
	}
}

func (m *modelManager) Open(name string, config string) {
	namespace,has:= m.getNameSpace(name)
	if !has{
		return
	}
	hs:=m.handlers[name]

	for _,handler:=range hs{
		handler.OnOpen(namespace,name,config)
	}
}

func newModelManager() *modelManager {
	return &modelManager{
		modules:        make(map[string]IModule),
		names:          make([]string,0,5),
		namespaceNames: make(map[string]string),
		handlers: 		make(map[string][]ConfigHandler),
	}
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
func (m *modelManager) getNameSpace(name string) (namespace string,has bool) {
	namespace,has = m.namespaceNames[name]
	return
}
func (m *modelManager) getModuleCount() int {
	count := len(m.modules)
	return count
}

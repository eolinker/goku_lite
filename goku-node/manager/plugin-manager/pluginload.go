package plugin_manager

import (
	"fmt"
	"path/filepath"
	"plugin"
	"reflect"
	"sync"

	goku_plugin "github.com/eolinker/goku-plugin"
	node_common "github.com/eolinker/goku/goku-node/node-common"

	entity "github.com/eolinker/goku/server/entity/node-entity"
)

var (
	globalPluginManager = &_GlodPluginManager{
		gloadPlugin:       make(map[string]goku_plugin.PluginFactory),
		gloadPluginLocker: sync.RWMutex{},
		errors:            make(map[string]error),
		errorCodes:        make(map[string]int),
	}
)

type _GlodPluginManager struct {
	gloadPlugin       map[string]goku_plugin.PluginFactory
	gloadPluginLocker sync.RWMutex
	errors            map[string]error
	errorCodes        map[string]int
}

func (m *_GlodPluginManager) check(name string) (int, error) {

	has, code, err := m.checkError(name)
	if has {
		return code, err
	}

	_, e, errorCode := m.loadPlugin(name)
	return errorCode, e

}
func (m *_GlodPluginManager) checkError(name string) (bool, int, error) {
	m.gloadPluginLocker.RLock()
	defer m.gloadPluginLocker.RUnlock()
	code, has := m.errorCodes[name]
	if has {
		return true, code, m.errors[name]
	}

	return false, 0, nil
}

func (m *_GlodPluginManager) getPluginHandle(name string) (goku_plugin.PluginFactory, bool) {
	m.gloadPluginLocker.RLock()
	defer m.gloadPluginLocker.RUnlock()

	p, has := m.gloadPlugin[name]
	return p, has
}

// 获取所有插件列表
func LoadPlugin(pis map[string]*entity.PluginInfo) (allFactory map[string]*entity.PluginFactoryHandler, defaultPlugins []*entity.PluginHandlerExce, beforMatchs []*entity.PluginHandlerExce) {
	plugins := make(map[string]*entity.PluginFactoryHandler)
	def := make([]*entity.PluginHandlerExce, 0, len(pis))
	before := make([]*entity.PluginHandlerExce, 0, len(pis))

	for key, value := range pis {
		handle, err, _ := globalPluginManager.loadPlugin(key)
		if err != nil {
			goku_plugin.Warn("LoadPlugin:",err.Error())
			continue
		}
		factory := &entity.PluginFactoryHandler{
			Factory: handle,
			Info:    value,
		}
		plugins[key] = factory

		pobj, err := factory.Factory.Create(value.Config, node_common.ClusterName(), value.UpdateTag, "", 0)
		if err != nil {
			continue
		}
		if value.Type == entity.PluginTypeGateway {

			def = append(def, &entity.PluginHandlerExce{
				PluginObj: pobj,
				Priority:  value.Priority,
				Name:      value.Name,
				IsStop:    value.IsStop,
			})
		} else {
			if pobj.BeforeMatch == nil || reflect.ValueOf(pobj.BeforeMatch).IsNil() {
				continue
			}
			before = append(before, &entity.PluginHandlerExce{
				PluginObj: pobj,
				Priority:  value.Priority,
				Name:      value.Name,
				IsStop:    value.IsStop,
			})
		}

	}
	return plugins, def, before
}

// 加载动态库
func (m *_GlodPluginManager) loadPlugin(name string) (goku_plugin.PluginFactory, error, int) {
	handle, has := m.getPluginHandle(name)
	if has {
		return handle, nil, 0
	}
	m.gloadPluginLocker.Lock()
	defer m.gloadPluginLocker.Unlock()

	handle, has = m.gloadPlugin[name]
	if has {
		return handle, nil, LoadOk
	}

	path, _ := filepath.Abs(fmt.Sprintf("plugin/%s.so", name))
	pdll, err := plugin.Open(path)
	if err != nil {
		e := fmt.Errorf("The plugin file named '%s.so' can not be found in plugin:%s ", name, err.Error())
		m.errors[name] = e
		m.errorCodes[name] = LoadFileError
		return nil, e, LoadFileError
	}

	//structName := strings.Replace(name, "-", "_", -1)
	//variableName := strings.ToUpper(string(structName[0])) + structName[1:]

	// 在插件中寻找相关的对象，将其方法加载
	v, err := pdll.Lookup("Builder")
	if err != nil {

		e := fmt.Errorf("The Builder  can not be found in plugin/%s.so ", name)
		m.errors[name] = e
		m.errorCodes[name] = LoadLookupError

		return nil, e, LoadLookupError
	}

	vp, ok := v.(func() goku_plugin.PluginFactory)
	if !ok {
		e := fmt.Errorf("The builder func  can not  implemented interface named goku_plugin.PluginFactory:%s ",name)
		m.errors[name] = e
		m.errorCodes[name] = LoadInterFaceError
		return nil, e, LoadInterFaceError
	}
	factory := vp()
	if factory == nil || reflect.ValueOf(factory).IsNil() {
		e := fmt.Errorf("The builder result is nil:%s ",name)
		m.errors[name] = e
		m.errorCodes[name] = LoadInterFaceError
		return nil, e, LoadInterFaceError
	}
	m.gloadPlugin[name] = factory
	m.errorCodes[name] = LoadOk
	m.errors[name] = nil
	return factory, nil, LoadOk

}

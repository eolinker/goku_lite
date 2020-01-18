package plugin_loader

import (
	"fmt"
	"path/filepath"
	"plugin"
	"reflect"
	"sync"

	goku_plugin "github.com/eolinker/goku-plugin"
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

	_, errorCode, e := m.loadPlugin(name)
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

//LoadPlugin 加载插件
func LoadPlugin(name string) (goku_plugin.PluginFactory, error) {
	factory, _, err := globalPluginManager.loadPlugin(name)

	return factory, err

}

// 加载动态库
func (m *_GlodPluginManager) loadPlugin(name string) (goku_plugin.PluginFactory, int, error) {
	handle, has := m.getPluginHandle(name)
	if has {
		return handle, LoadOk, nil
	}
	m.gloadPluginLocker.Lock()
	defer m.gloadPluginLocker.Unlock()

	path, _ := filepath.Abs(fmt.Sprintf("plugin/%s.so", name))

	pdll, err := plugin.Open(path)
	if err != nil {
		e := fmt.Errorf("plugin:%s ", name, err.Error())
		m.errors[name] = e
		m.errorCodes[name] = LoadFileError
		return nil, LoadFileError, e
	}

	//structName := strings.Replace(name, "-", "_", -1)
	//variableName := strings.ToUpper(string(structName[0])) + structName[1:]

	// 在插件中寻找相关的对象，将其方法加载
	v, err := pdll.Lookup("Builder")
	if err != nil {

		e := fmt.Errorf("The Builder can not be found in plugin/%s.so ", name)
		m.errors[name] = e
		m.errorCodes[name] = LoadLookupError
		return nil, LoadLookupError, e
	}

	vp, ok := v.(func() goku_plugin.PluginFactory)
	if !ok {
		e := fmt.Errorf("The builder func  can not  implemented interface named goku_plugin.PluginFactory:%s ", name)
		m.errors[name] = e
		m.errorCodes[name] = LoadInterFaceError
		return nil, LoadInterFaceError, e
	}
	factory := vp()
	if factory == nil || reflect.ValueOf(factory).IsNil() {
		e := fmt.Errorf("The builder result is nil:%s ", name)
		m.errors[name] = e
		m.errorCodes[name] = LoadInterFaceError
		fmt.Println(e)
		return nil, LoadInterFaceError, e
	}
	m.gloadPlugin[name] = factory
	m.errorCodes[name] = LoadOk
	m.errors[name] = nil
	return factory, LoadOk, nil

}

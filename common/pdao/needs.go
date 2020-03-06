package pdao

import (
	"fmt"
	"reflect"
	"sync"
)

//NeedsManager needsManager
type NeedsManager struct {
	daoInterfaces map[string][]*reflect.Value
	lock          sync.Mutex
}

//NewNeedsManager 创建新的needsManager
func NewNeedsManager() *NeedsManager {
	return &NeedsManager{
		daoInterfaces: make(map[string][]*reflect.Value),
		lock:          sync.Mutex{},
	}
}

func (m *NeedsManager) add(key string, v *reflect.Value) {
	m.lock.Lock()
	m.daoInterfaces[key] = append(m.daoInterfaces[key], v)
	m.lock.Unlock()
}

func (m *NeedsManager) set(key string, v reflect.Value) {
	m.lock.Lock()
	for _, e := range m.daoInterfaces[key] {
		e.Set(v)
	}
	delete(m.daoInterfaces, key)
	m.lock.Unlock()
}
func (m *NeedsManager) check() []string {
	m.lock.Lock()
	r := make([]string, 0, len(m.daoInterfaces))
	for pkg := range m.daoInterfaces {

		r = append(r, pkg)
	}
	m.lock.Unlock()
	return r
}

//Need 声明
func (m *NeedsManager) Need(p interface{}) {

	v := reflect.ValueOf(p)

	if v.Kind() != reflect.Ptr {
		panic("must ptr")
	}
	e := v.Elem()
	pkg := key(e.Type())
	if pkg == "" {
		panic("invalid interface")
	}

	if !e.CanSet() {
		panic("invalid interface")
	}
	m.add(pkg, &e)
}

//Set 注入
func (m *NeedsManager) Set(i interface{}) {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		panic("must ptr")
	}
	e := v.Elem()
	pkg := key(e.Type())

	if pkg == "" {
		panic("invalid interface")
	}
	m.set(pkg, e)
}
func key(t reflect.Type) string {
	return fmt.Sprintf("%s.%s", t.PkgPath(), t.String())
}

//Check 检查是否实现相关dao类
func (m *NeedsManager) Check() error {
	rs := m.check()
	if len(rs) > 0 {
		return fmt.Errorf("need:%v", rs)
	}
	return nil
}

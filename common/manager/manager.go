package manager

import "sync"

//Manager manager
type Manager struct {
	locker sync.RWMutex
	objs   map[string]interface{}
}

//NewManager new manager
func NewManager() *Manager {
	return &Manager{
		locker: sync.RWMutex{},
		objs:   make(map[string]interface{}),
	}
}

//Get get
func (m *Manager) Get(key string) (interface{}, bool) {

	m.locker.RLock()
	v, has := m.objs[key]
	m.locker.RUnlock()
	return v, has
}

//Set set
func (m *Manager) Set(key string, value interface{}) {
	m.locker.Lock()
	m.objs[key] = value
	m.locker.Unlock()
}

//Value value
type Value struct {
	locker sync.RWMutex
	isHas  bool
	value  interface{}
}

//NewValue new value
func NewValue() *Value {
	return &Value{
		locker: sync.RWMutex{},
		isHas:  false,
		value:  nil,
	}
}

//Set set
func (v *Value) Set(value interface{}) {
	v.locker.Lock()

	v.value = value
	v.isHas = true
	v.locker.Unlock()
}

//Get get
func (v *Value) Get() (interface{}, bool) {
	v.locker.RLock()
	value, has := v.value, v.isHas

	v.locker.RUnlock()
	return value, has
}

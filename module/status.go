package module

import "sync"

//State state
type State struct {
	name   string
	isOpen bool
}

//Manager manager
type Manager struct {
	locker  sync.RWMutex
	modules map[string]bool
}

var m = Manager{
	modules: make(map[string]bool),
	locker:  sync.RWMutex{},
}

//Register register
func Register(name string, isOpen bool) {
	m.locker.Lock()
	m.modules[name] = isOpen
	m.locker.Unlock()
}

//IsOpen isOpen
func IsOpen(name string) bool {
	m.locker.RLock()
	isOpen := m.modules[name]
	m.locker.RUnlock()
	return isOpen
}

//Close close
func Close(name string) {
	m.locker.Lock()
	m.modules[name] = false
	m.locker.Unlock()

}

//Open open
func Open(name string) {
	m.locker.Lock()
	m.modules[name] = true
	m.locker.Unlock()
}

//Refresh refresh
func Refresh(States []State) {

	m.locker.Lock()

	for _, s := range States {
		m.modules[s.name] = s.isOpen
	}
	m.locker.Unlock()
}


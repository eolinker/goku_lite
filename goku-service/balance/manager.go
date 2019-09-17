package balance

import (
	"sync"
)

var manager = &Manager{
	locker:sync.RWMutex{},
	balances:make(map[string]*Balance),
}



type Manager struct {
	locker sync.RWMutex
	balances map[string]*Balance
}

func (m *Manager)set(balances map[string]*Balance)  {
	m.locker.Lock()
	m.balances = balances
	m.locker.Unlock()
}

func (m *Manager)get(name string)( *Balance,bool) {
	m.locker.RLock()

	b,has:=m.balances[name]
	m.locker.RUnlock()

	return b,has
}



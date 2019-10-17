package balance

import (
	"sync"

	"github.com/eolinker/goku-api-gateway/config"
)

var manager = &Manager{
	locker:   sync.RWMutex{},
	balances: make(map[string]*config.BalanceConfig),
}

//Manager manager
type Manager struct {
	locker   sync.RWMutex
	balances map[string]*config.BalanceConfig
}

func (m *Manager) set(balances map[string]*config.BalanceConfig) {
	m.locker.Lock()
	m.balances = balances
	m.locker.Unlock()
}

func (m *Manager) get(name string) (*config.BalanceConfig, bool) {
	m.locker.RLock()

	b, has := m.balances[name]
	m.locker.RUnlock()

	return b, has
}

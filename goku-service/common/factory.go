package common

import (
	"fmt"
	"sync"
)

//InstanceFactory instanceFactory
type InstanceFactory struct {
	locker    sync.RWMutex
	instances map[string]*Instance
}

//NewInstanceFactory 创建InstanceFactory
func NewInstanceFactory() *InstanceFactory {
	return &InstanceFactory{
		locker:    sync.RWMutex{},
		instances: make(map[string]*Instance),
	}
}

//General general
func (m *InstanceFactory) General(ip string, port int, weight int) *Instance {
	if weight < 1 {
		weight = 1
	}
	key := fmt.Sprintf("%s:%d-%d", ip, port, weight)
	m.locker.RLock()
	i, h := m.instances[key]
	m.locker.RUnlock()
	if h {
		i.Weight = weight
		return i
	}
	m.locker.Lock()
	i, h = m.instances[key]
	if h {
		m.locker.Unlock()
		i.Weight = weight

		return i
	}
	i = &Instance{
		InstanceID: fmt.Sprintf("%s:%d", ip, port),
		IP:         ip,
		Port:       port,
		Weight:     weight,
		Status:     InstanceRun,
		locker:     sync.RWMutex{},
	}
	m.instances[key] = i
	m.locker.Unlock()
	return i

}

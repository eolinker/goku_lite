package diting

import "sync"

//Refreshers refreshers
type Refreshers struct {
	proxies []Proxy
	locker  sync.RWMutex
}

//NewRefreshers new refreshers
func NewRefreshers() *Refreshers {
	return &Refreshers{
		proxies: nil,
		locker:  sync.RWMutex{},
	}
}

//Add add
func (r *Refreshers) Add(proxy Proxy) {
	r.locker.Lock()

	r.proxies = append(r.proxies, proxy)

	r.locker.Unlock()
}

//Refresh refresh
func (r *Refreshers) Refresh(factories Factories) {
	r.locker.RLock()
	proxies := r.proxies
	r.locker.RUnlock()
	for _, p := range proxies {
		p.Refresh(factories)
	}
}

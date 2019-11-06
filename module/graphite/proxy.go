package graphite

import (
	"sync"

	"github.com/marpaia/graphite-golang"
)

//Proxy proxy
type Proxy struct {
	locker   sync.RWMutex
	graphite *graphite.Graphite
}

//NewProxy new proxy
func NewProxy() *Proxy {
	return &Proxy{
		locker:   sync.RWMutex{},
		graphite: nil,
	}
}

//Send send
func (p *Proxy) Send(metrics []graphite.Metric) {
	p.locker.RLock()
	g := p.graphite
	p.locker.RUnlock()
	if g != nil {
		g.SendMetrics(metrics)
	}
}

//Connect connect
func (p *Proxy) Connect(host string, port int) error {

	p.close()

	newG, err := graphite.NewGraphite(host, port)
	if err != nil {
		return err
	}
	p.locker.Lock()
	p.graphite = newG
	p.locker.Unlock()
	return nil
}
func (p *Proxy) close() (err error) {
	p.locker.Lock()
	if p.graphite != nil {
		err = p.graphite.Disconnect()
		p.graphite = nil
	}
	p.locker.Unlock()
	return
}

func (p *Proxy) isClose() bool {
	p.locker.Lock()
	isClose := p.graphite == nil
	p.locker.Unlock()
	return isClose
}

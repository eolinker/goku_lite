package diting

import "sync"

//GaugeProxy gaugeProxy
type GaugeProxy struct {
	ConstLabelsProxy
	opt    *GaugeOpts
	gauges Gauges
	locker sync.RWMutex
}

func newGaugeProxy(opt *GaugeOpts) *GaugeProxy {
	return &GaugeProxy{
		ConstLabelsProxy: ConstLabelsProxy(opt.ConstLabels),
		opt:              opt,
		gauges:           nil,
		locker:           sync.RWMutex{},
	}
}

//Refresh refresh
func (g *GaugeProxy) Refresh(factories Factories) {
	gauges, _ := factories.NewGauge(g.opt)

	g.locker.Lock()
	g.gauges = gauges
	g.locker.Unlock()
}

//Set set
func (g *GaugeProxy) Set(value float64, labels Labels) {
	g.compile(labels)
	g.locker.RLock()
	gauges := g.gauges
	g.locker.RUnlock()

	gauges.Set(value, labels)

}

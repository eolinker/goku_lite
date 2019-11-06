package diting

import "sync"

//CounterProxy counterProxy
type CounterProxy struct {
	ConstLabelsProxy
	opt      *CounterOpts
	counters Counters
	locker   sync.RWMutex
}

func newCounterProxy(opt *CounterOpts) *CounterProxy {
	return &CounterProxy{
		ConstLabelsProxy: ConstLabelsProxy(opt.ConstLabels),
		opt:              opt,
		counters:         nil,
		locker:           sync.RWMutex{},
	}
}

//Refresh refresh
func (c *CounterProxy) Refresh(factories Factories) {

	counters, _ := factories.NewCounter(c.opt)
	c.locker.Lock()
	c.counters = counters
	c.locker.Unlock()
}

//Add add
func (c *CounterProxy) Add(value float64, labels Labels) {
	c.compile(labels)
	c.locker.RLock()
	counters := c.counters
	c.locker.RUnlock()
	counters.Add(value, labels)

}

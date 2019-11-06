package diting

import "sync"

//CacheFactory cacheFactory
type CacheFactory struct {
	factory Factory

	cache  map[uint64]interface{}
	locker sync.RWMutex
}

//NewCacheFactory new cacheFactory
func NewCacheFactory(factory Factory) *CacheFactory {
	return &CacheFactory{
		factory: factory,
		cache:   make(map[uint64]interface{}),
		locker:  sync.RWMutex{},
	}

}

//NewCounter new counter
func (f *CacheFactory) NewCounter(opt *CounterOpts) (Counter, error) {
	v, has := f.get(opt.ID)

	if has {
		return v.(Counter), nil
	}

	c, err := f.factory.NewCounter(opt)
	if err != nil {
		return nil, err
	}

	v = f.set(opt.ID, c)

	return v.(Counter), nil
}

//NewHistogram new histogram
func (f *CacheFactory) NewHistogram(opt *HistogramOpts) (Histogram, error) {
	v, has := f.get(opt.ID)

	if has {
		return v.(Histogram), nil
	}

	c, err := f.factory.NewHistogram(opt)
	if err != nil {
		return nil, err
	}

	v = f.set(opt.ID, c)

	return v.(Histogram), nil
}

//NewGauge new gauge
func (f *CacheFactory) NewGauge(opt *GaugeOpts) (Gauge, error) {
	v, has := f.get(opt.ID)

	if has {
		return v.(Gauge), nil
	}

	c, err := f.factory.NewGauge(opt)
	if err != nil {
		return nil, err
	}

	v = f.set(opt.ID, c)

	return v.(Gauge), nil
}
func (f *CacheFactory) get(id uint64) (interface{}, bool) {
	f.locker.RLock()
	v, has := f.cache[id]
	f.locker.RUnlock()
	return v, has
}

func (f *CacheFactory) set(id uint64, v interface{}) interface{} {
	f.locker.Lock()
	vo, has := f.cache[id]
	if has {
		f.locker.Unlock()
		return vo
	}
	f.cache[id] = v
	f.locker.Unlock()
	return v

}

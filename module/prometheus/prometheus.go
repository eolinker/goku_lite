package prometheus

import (
	"sync"

	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/eolinker/goku-api-gateway/module"
	"github.com/eolinker/goku-api-gateway/module/prometheus/config"
	"github.com/eolinker/goku-api-gateway/node/admin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//Register register
func Register() {

	config.Register()
	p := NewPrometheus()

	handler := promhttp.InstrumentMetricHandler(
		p.registry, promhttp.HandlerFor(p.registry, promhttp.HandlerOpts{}),
	)

	module.Register(config.ModuleNameSpace, true)
	admin.Add(config.ModuleNameSpace, config.Pattern, handler)

	diting.Register(config.ModuleNameSpace, p)

}

//Prometheus prometheus
type Prometheus struct {
	registry     *prometheus.Registry
	cacheFactory *diting.CacheFactory

	collectors      []prometheus.Collector
	collectorLocker sync.Mutex
}

func (p *Prometheus) add(c prometheus.Collector) {
	p.collectorLocker.Lock()
	p.collectors = append(p.collectors, c)
	p.collectorLocker.Unlock()
}

//Close close
func (p *Prometheus) Close() {
	module.Close(config.ModuleNameSpace)
	p.collectorLocker.Lock()
	collectors := p.collectors
	l := len(collectors)
	if l == 0 {
		l = 2
	}
	p.collectors = make([]prometheus.Collector, 0, l)
	p.cacheFactory = diting.NewCacheFactory(p)
	p.collectorLocker.Unlock()

	for _, c := range collectors {
		p.registry.Unregister(c)
	}
	return
}

//NewPrometheus new Prometheus
func NewPrometheus() *Prometheus {
	p := &Prometheus{
		registry:     prometheus.NewPedanticRegistry(),
		cacheFactory: nil,
	}
	p.cacheFactory = diting.NewCacheFactory(p)
	return p
}

//NewSummary new Summary
func (p *Prometheus) NewSummary(opts *diting.SummaryOpts) (diting.Summary, error) {
	s := prometheus.NewSummaryVec(ReadSummaryOpts(opts), opts.LabelNames)
	err := p.registry.Register(s)
	if err != nil {
		return nil, err
	}
	return newSummary(s), nil
}

//NewCounter new Counter
func (p *Prometheus) NewCounter(opts *diting.CounterOpts) (diting.Counter, error) {
	c := prometheus.NewCounterVec(ReadCounterOpts(opts), opts.LabelNames)
	err := p.registry.Register(c)
	if err != nil {
		return nil, err
	}
	return newCounter(c), nil
}

//NewHistogram new HistogramObserve
func (p *Prometheus) NewHistogram(opts *diting.HistogramOpts) (diting.Histogram, error) {
	h := prometheus.NewHistogramVec(ReadHistogramOpts(opts), opts.LabelNames)
	err := p.registry.Register(h)
	if err != nil {
		return nil, err
	}
	return newHistogram(h), nil
}

//NewGauge new gauge
func (p *Prometheus) NewGauge(opts *diting.GaugeOpts) (diting.Gauge, error) {
	g := prometheus.NewGaugeVec(ReadGaugeOpts(opts), opts.LabelNames)
	err := p.registry.Register(g)
	if err != nil {
		return nil, err
	}
	return newGauge(g), nil
}

//Namespace namespace
func (p *Prometheus) Namespace() string {
	return config.ModuleNameSpace
}

//Create create
func (p *Prometheus) Create(conf string) (diting.Factory, error) {
	module.Open(config.ModuleNameSpace)
	return p.cacheFactory, nil
}

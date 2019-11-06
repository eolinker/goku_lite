package graphite

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/eolinker/goku-api-gateway/module"
	"github.com/eolinker/goku-api-gateway/module/graphite/config"
)

//Constructor constructor
type Constructor struct {
	address       string
	graphiteProxy *Proxy
	cacheFactory  *diting.CacheFactory
	locker        sync.RWMutex
	metrics       []Metrics

	cancelFunc context.CancelFunc
}

//Register register
func Register() {
	config.Register()
	g := &Constructor{
		graphiteProxy: NewProxy(),
	}
	g.cacheFactory = diting.NewCacheFactory(g)

	module.Register(config.ModuleNameSpace, true)
	diting.Register(config.ModuleNameSpace, g)
}

//Close close
func (g *Constructor) Close() {
	module.Close(config.ModuleNameSpace)
	g.stop()
}
func (g *Constructor) addMetrics(metrics Metrics) {
	g.locker.Lock()
	g.metrics = append(g.metrics, metrics)
	g.locker.Unlock()
}

//Send send
func (g *Constructor) Send() {
	g.locker.RLock()
	ms := g.metrics
	g.locker.RUnlock()

	for _, m := range ms {
		metrics := m.Metrics()
		if len(metrics) > 0 {
			g.graphiteProxy.Send(metrics)
		}

	}
}

//Namespace nameSpace
func (g *Constructor) Namespace() string {
	return config.ModuleNameSpace
}
func (g *Constructor) stop() {
	g.locker.Lock()
	if g.address != "" {
		g.cancelFunc()
		g.address = ""
		g.cancelFunc = nil
	}
	g.locker.Unlock()
}

func (g *Constructor) start(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)
	g.locker.Lock()
	defer g.locker.Unlock()
	if g.address != addr {

		ctx, cancel := context.WithCancel(context.Background())
		go g.doLoop(ctx, host, port)
		g.address = addr
		g.cancelFunc = cancel
	}
}
func (g *Constructor) doLoop(ctx context.Context, host string, port int) {

	e := g.graphiteProxy.Connect(host, port)
	if e != nil {
		return
	}
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			g.Send()
		case <-ctx.Done():
			g.Send()
			g.graphiteProxy.close()
			return
		}
	}

}

//Create create
func (g *Constructor) Create(conf string) (diting.Factory, error) {

	confV, err := config.Decode(conf)
	if err != nil {
		go g.stop()
		return nil, err
	}

	par := strings.Split(confV.AccessAddress, ":")
	if len(par) != 2 {
		go g.stop()
		return nil, errors.New("invalid AccessAddress")
	}

	host := par[0]
	port, err := strconv.Atoi(par[1])
	if err != nil {
		go g.stop()
		return nil, err
	}

	go g.start(host, port)
	module.Open(config.ModuleNameSpace)
	return g.cacheFactory, nil
}

//NewCounter new counter
func (g *Constructor) NewCounter(opt *diting.CounterOpts) (diting.Counter, error) {

	metricKey := NewMetricKey(toLabelName(opt.Namespace, opt.Subsystem, opt.Name), opt.LabelNames)
	c := NewCounter(metricKey)
	g.addMetrics(c)
	return c, nil
}

//NewHistogram new histogram
func (g *Constructor) NewHistogram(opt *diting.HistogramOpts) (diting.Histogram, error) {
	metricKey := NewMetricKey(toLabelName(opt.Namespace, opt.Subsystem, opt.Name), opt.LabelNames)
	h := NewHistogram(metricKey, opt.Buckets)
	g.addMetrics(h)
	return h, nil
}

//NewGauge new gauge
func (g *Constructor) NewGauge(opt *diting.GaugeOpts) (diting.Gauge, error) {
	metricKey := NewMetricKey(toLabelName(opt.Namespace, opt.Subsystem, opt.Name), opt.LabelNames)
	gauge := NewGauge(metricKey)
	g.addMetrics(gauge)
	return gauge, nil
}

func toLabelName(Namespace string, Subsystem string, Name string) string {
	tmp := make([]string, 0, 3)

	if Namespace != "" {
		tmp = append(tmp, Namespace)
	}
	if Subsystem != "" {
		tmp = append(tmp, Subsystem)
	}
	if Name != "" {
		tmp = append(tmp, Name)
	}
	return strings.Join(tmp, "_")
}

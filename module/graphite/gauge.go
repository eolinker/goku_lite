package graphite

import (
	"strconv"
	"sync"
	"time"

	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/marpaia/graphite-golang"
)

//Gauge gauge
type Gauge struct {
	metricKey          MetricKey
	metricsValuesGauge *MetricsValuesGauge
}

//NewGauge new gauge
func NewGauge(metricKey MetricKey) *Gauge {
	return &Gauge{metricKey: metricKey, metricsValuesGauge: NewMetricsValuesGauge()}
}

//Set set
func (g *Gauge) Set(value float64, labels diting.Labels) {

	key := g.metricKey.Key(labels, "value")
	g.metricsValuesGauge.Set(key, value)
}

//Metrics metrics
func (g *Gauge) Metrics() []graphite.Metric {
	values := g.metricsValuesGauge.Collapse()
	if len(values) == 0 {
		return nil
	}
	ms := make([]graphite.Metric, 0, len(values))
	t := time.Now().Unix()

	for k, v := range values {
		ms = append(ms, graphite.NewMetric(k, strconv.FormatFloat(v, 'f', 2, 64), t))
	}
	return ms
}

//MetricsValuesGauge metricsValuesGauge
type MetricsValuesGauge struct {
	values map[string]float64
	locker sync.Mutex
}

//NewMetricsValuesGauge new metricsValuesGauge
func NewMetricsValuesGauge() *MetricsValuesGauge {
	return &MetricsValuesGauge{
		values: make(map[string]float64),
		locker: sync.Mutex{},
	}
}

//Set set
func (m *MetricsValuesGauge) Set(key string, value float64) {
	m.locker.Lock()
	m.values[key] = value
	m.locker.Unlock()
}

//Collapse collapse
func (m *MetricsValuesGauge) Collapse() map[string]float64 {
	n := make(map[string]float64)
	m.locker.Lock()
	v := m.values
	m.values = n
	m.locker.Unlock()

	return v

}

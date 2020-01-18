package graphite

import (
	"strconv"
	"sync"
	"time"

	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/marpaia/graphite-golang"
)

//Count count
type Count struct {
	metricKey          MetricKey
	metricsValuesCount *MetricsValuesCount
}

//RegisterDao add
func (c *Count) Add(value float64, labels diting.Labels) {

	key := c.metricKey.Key(labels, "count")
	c.metricsValuesCount.Add(key, value)
}

//Metrics metrics
func (c *Count) Metrics() []graphite.Metric {
	values := c.metricsValuesCount.Collapse()
	if len(values) == 0 {
		return nil
	}
	ms := make([]graphite.Metric, 0, len(values))
	t := time.Now().Unix()

	for k, v := range values {
		ms = append(ms, graphite.NewMetric(k, strconv.FormatInt(v, 10), t))
	}

	return ms

}

//NewCounter new counter
func NewCounter(metricKey MetricKey) *Count {
	return &Count{
		metricKey:          metricKey,
		metricsValuesCount: NewMetricsValuesCount(),
	}
}

//MetricsValuesCount metricsValuesCount
type MetricsValuesCount struct {
	values map[string]int64
	locker sync.Mutex
}

//NewMetricsValuesCount new metricsValuesCount
func NewMetricsValuesCount() *MetricsValuesCount {
	return &MetricsValuesCount{
		values: make(map[string]int64),
		locker: sync.Mutex{},
	}
}

//RegisterDao add
func (m *MetricsValuesCount) Add(key string, value float64) {
	m.locker.Lock()
	v := int64(value)
	m.values[key] += v
	m.locker.Unlock()
}

//Collapse collapse
func (m *MetricsValuesCount) Collapse() map[string]int64 {
	n := make(map[string]int64)

	m.locker.Lock()
	v := m.values
	m.values = n
	m.locker.Unlock()

	return v
}

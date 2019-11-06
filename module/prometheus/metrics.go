package prometheus

import (
	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/prometheus/client_golang/prometheus"
)

//Counter counter
type Counter struct {
	counterVec *prometheus.CounterVec
}

func newCounter(counterVec *prometheus.CounterVec) *Counter {
	return &Counter{counterVec: counterVec}
}

//Add add
func (c *Counter) Add(value float64, labels diting.Labels) {
	c.counterVec.With(ReadLabels(labels)).Add(value)
}

//Gauge gauge
type Gauge struct {
	gaugeVec *prometheus.GaugeVec
}

func newGauge(gaugeVec *prometheus.GaugeVec) *Gauge {
	return &Gauge{gaugeVec: gaugeVec}

}

//Set set
func (g *Gauge) Set(value float64, labels diting.Labels) {
	g.gaugeVec.With(ReadLabels(labels)).Set(value)
}

//Histogram histogram
type Histogram struct {
	histogramVec *prometheus.HistogramVec
}

func newHistogram(histogramVec *prometheus.HistogramVec) *Histogram {
	return &Histogram{histogramVec: histogramVec}

}

//Observe observe
func (h *Histogram) Observe(value float64, labels diting.Labels) {
	h.histogramVec.With(ReadLabels(labels)).Observe(value)
}

//Summary summary
type Summary struct {
	summaryVec *prometheus.SummaryVec
}

func newSummary(summaryVec *prometheus.SummaryVec) *Summary {
	return &Summary{summaryVec: summaryVec}
}

//Observe observe
func (s *Summary) Observe(value float64, labels diting.Labels) {
	s.summaryVec.With(ReadLabels(labels)).Observe(value)
}

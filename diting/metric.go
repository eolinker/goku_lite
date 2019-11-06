package diting

//Labels labels
type Labels map[string]string

//Counter counter
type Counter interface {
	Add(value float64, labels Labels)
}

//Gauge gauge
type Gauge interface {
	Set(value float64, labels Labels)
}

//Observer observer
type Observer interface {
	Observe(value float64, labels Labels)
}

//Histogram histogram
type Histogram interface {
	Observer
}

//Summary summary
type Summary interface {
	Observer
}

//Counters counters
type Counters []Counter

//Add add
func (cs Counters) Add(value float64, labels Labels) {
	for _, c := range cs {
		c.Add(value, labels)
	}
}

//Gauges gauges
type Gauges []Gauge

//Set set
func (gs Gauges) Set(value float64, labels Labels) {
	for _, g := range gs {
		g.Set(value, labels)
	}
}

//Histograms histograms
type Histograms []Observer

//Observe observe
func (hs Histograms) Observe(value float64, labels Labels) {
	for _, h := range hs {
		h.Observe(value, labels)
	}
}

//Summaries summaries
type Summaries []Summary

//Observe observe
func (ss Summaries) Observe(value float64, labels Labels) {
	for _, s := range ss {
		s.Observe(value, labels)
	}
}

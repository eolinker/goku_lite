package diting

//NewCounter new counter
func NewCounter(opt *CounterOpts) Counter {
	c := newCounterProxy(opt)
	refresher.Add(c)
	return c
}

//NewGauge new gauge
func NewGauge(opt *GaugeOpts) Gauge {
	g := newGaugeProxy(opt)
	refresher.Add(g)
	return g
}

//NewHistogram new HistogramObserve
func NewHistogram(opt *HistogramOpts) Histogram {
	h := newHistogramProxy(opt)
	refresher.Add(h)
	return h
}

//func  NewSummary(opt *SummaryOpts) Summary {
//	s:= newSummariesProxy(opt)
//	refresher.RegisterDao(s)
//	return s
//}

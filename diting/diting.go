package diting

//ConfigHandleFunc configHandleFunc
type ConfigHandleFunc func() interface{}

//Constructor constructor
type Constructor interface {
	Namespace() string
	Create(conf string) (Factory, error)
	Close()
}

//Factory factory
type Factory interface {
	NewCounter(opt *CounterOpts) (Counter, error)
	//NewSummary(opt *SummaryOpts) (Summary,error)
	NewHistogram(opt *HistogramOpts) (Histogram, error)
	NewGauge(opt *GaugeOpts) (Gauge, error)
}

//Factories factories
type Factories []Factory

//NewCounter new counter
func (fs Factories) NewCounter(opt *CounterOpts) (Counters, error) {
	cs := make(Counters, 0, len(fs))
	for _, f := range fs {

		s, err := f.NewCounter(opt)
		if err != nil {
			continue
		}
		cs = append(cs, s)
	}
	return cs, nil
}

//NewHistogram new histogram
func (fs Factories) NewHistogram(opt *HistogramOpts) (Histograms, error) {
	hs := make(Histograms, 0, len(fs))
	for _, f := range fs {

		h, err := f.NewHistogram(opt)
		if err != nil {
			continue
		}
		hs = append(hs, h)
	}
	return hs, nil
}

//NewGauge new gauge
func (fs Factories) NewGauge(opt *GaugeOpts) (Gauges, error) {
	gs := make(Gauges, 0, len(fs))
	for _, f := range fs {

		g, err := f.NewGauge(opt)
		if err != nil {
			continue
		}
		gs = append(gs, g)
	}
	return gs, nil
}

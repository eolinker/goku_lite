package goku_observe

//HistogramObserve HistogramObserve
type HistogramObserve interface {
	Observe(buckets []float64, value float64)
	Collapse() (values []uint64, sum, max, min float64, count uint64)
}

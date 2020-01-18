package goku_observe

import (
	"math"
	"sync"
)

//type Item struct {
//	Reference float64
//	values int64
//}

//Histogram histogram
type Histogram struct {
	size   int
	Values []uint64
	Max    float64
	Min    float64
	Sum    float64
	Count  uint64
}

//HistogramObserver histogramObserver
type HistogramObserver struct {
	Histogram
	locker sync.Mutex
}

//Collapse Collapse
func (h *HistogramObserver) Collapse() (values []uint64, sum, max, min float64, count uint64) {
	return h.Values, h.Sum, h.Max, h.Min, h.Count
}

//Observe Observe
func (h *HistogramObserver) Observe(buckets []float64, value float64) {
	h.locker.Lock()
	if value < 0 {
		value = 0
	}

	h.Count++
	h.Sum += value
	if value > h.Max {
		h.Max = value
	}
	if value < h.Min {
		h.Min = value
	}
	l := h.size - 1
	for i := l; i >= 0; i-- {
		if value >= buckets[i] {
			break
		}
		h.Values[i]++
	}
	h.locker.Unlock()
}

//NewHistogram new Histogram
func NewHistogram(size int) *Histogram {
	return &Histogram{
		size:   size,
		Values: make([]uint64, size, size),
		Max:    0,
		Min:    math.MaxFloat64,
		Sum:    0,
		Count:  0,
	}
}

//NewHistogramObserve new HistogramObserve
func NewHistogramObserve(size int) HistogramObserve {

	h := &HistogramObserver{
		Histogram: Histogram{
			size:   size,
			Values: make([]uint64, size, size),
			Max:    0,
			Min:    math.MaxFloat64,
			Sum:    0,
			Count:  0,
		},
		locker: sync.Mutex{},
	}

	return h
}

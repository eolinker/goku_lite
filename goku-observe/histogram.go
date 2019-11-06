package goku_observe

import (
	"math"
	"sync"
)

//type Item struct {
//	Reference float64
//	values int64
//}

type _Histogram struct {
	size int
	buckets []float64
	values  []int64
	max float64
	min float64
	sum float64
	count int
	locker  sync.Mutex
}

func (h *_Histogram) Observe(value float64) {
	if value < 0 {
		value = 0
	}
	h.locker.Lock()

	h.count++
	h.sum += value
	if value> h.max{
		h.max = value
	}
	if value < h.min{
		h.min= value
	}
	l:= h.size -1
	for i:=l;i>=0;i--{
		if value >= h.buckets[i]{
			break
		}
		h.values[i]++
	}
	h.locker.Unlock()
}
//func (h*_Histogram) Reset()  {
//	h.locker.Lock()
//	h.values = make([]int64,len(h.values),len(h.values))
//	h.sum ,h.max ,h.min,h.count = 0,0,math.MaxFloat64,0
//	h.locker.Unlock()
//}
func (h *_Histogram) Collapse() (values []int64, sum ,max,min float64,count int){
	h.locker.Lock()
	values,sum,max,min,count = h.values,h.sum,h.max,h.min,h.count
	h.locker.Unlock()
	return
}
func NewHistogram(buckets []float64) *_Histogram {
	max:= len(buckets) +1
	h:=&_Histogram{
		size:    max,
		buckets: make([]float64, 0, max),
		values:  make([]int64, max, max),
		max:     0,
		min:     math.MaxFloat64,
		sum:     0,
		count:   0,
		locker:  sync.Mutex{},
	}
	h.buckets = append(h.buckets,buckets...)
	h.buckets = append(h.buckets,math.MaxFloat64)

	//sort.Float64s(h.buckets)
	return h
}



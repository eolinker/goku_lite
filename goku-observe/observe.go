package goku_observe

type Observe interface {
	Observe(value float64)
}

type Histogram interface {
	Observe
	Collapse() (values []int64, sum ,max,min float64,count int)
}

package diting

import "sync"

//HistogramProxy histogramProxy
type HistogramProxy struct {
	ConstLabelsProxy
	opt    *HistogramOpts
	locker sync.RWMutex

	histograms Histograms
}

func newHistogramProxy(opt *HistogramOpts) *HistogramProxy {
	return &HistogramProxy{
		ConstLabelsProxy: ConstLabelsProxy(opt.ConstLabels),
		opt:              opt,
		locker:           sync.RWMutex{},
		histograms:       nil,
	}
}

//Refresh refresh
func (h *HistogramProxy) Refresh(factories Factories) {
	histograms, _ := factories.NewHistogram(h.opt)

	h.locker.Lock()
	h.histograms = histograms
	h.locker.Unlock()
}

//Observe observe
func (h *HistogramProxy) Observe(value float64, labels Labels) {
	h.compile(labels)
	h.locker.RLock()
	histograms := h.histograms
	h.locker.RUnlock()
	histograms.Observe(value, labels)
}

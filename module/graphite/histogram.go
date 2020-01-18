package graphite

import (
	"math"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/eolinker/goku-api-gateway/diting"
	observe "github.com/eolinker/goku-api-gateway/goku-observe"
	"github.com/marpaia/graphite-golang"
)

//HistogramObserve histogram
type Histogram struct {
	metricKey    MetricKey
	buckets    []float64
	histograms map[string]observe.HistogramObserve
	locker     sync.Mutex
}

//NewHistogram new histogram
func NewHistogram(metricKey MetricKey, buckets []float64) *Histogram {
	return &Histogram{
		metricKey:    metricKey,
		buckets:    buckets,
		histograms: make(map[string]observe.HistogramObserve),
		locker:     sync.Mutex{},
	}
}

//Observe observe
func (h *Histogram) Observe(value float64, labels diting.Labels) {
	key := h.metricKey.Key(labels, "")
	h.locker.Lock()

	hv, has := h.histograms[key]
	if !has {
		hv = observe.NewHistogramObserve(len(h.buckets))
		h.histograms[key] = hv
	}

	h.locker.Unlock()
	hv.Observe(h.buckets,value)
}

//Metrics metrics
func (h *Histogram) Metrics() []graphite.Metric {
	all := h.Collapse()

	if len(all) == 0 {
		return nil
	}
	keySize := len(all)
	size := keySize * (len(h.buckets) + 5)
	ms := make([]graphite.Metric, 0, size)
	t := time.Now().Unix()
	tmpName := make([]string, 2)
	tmpBucketName := make([]string, 3)
	tmpBucketName[0] = "bucket_le"
	//overIndex := len(h.keyHistogram.buckets)

	for k, v := range all {
		tmpName[0] = k
		values, sum, max, min, count := v.Collapse()
		tmpName[1] = "count"
		ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatUint(count, 10), t))
		tmpName[1] = "max"
		ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatFloat(max, 'f', 2, 64), t))
		tmpName[1] = "min"
		ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatFloat(min, 'f', 2, 64), t))

		tmpName[1] = "sum"
		ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatFloat(sum, 'f', 2, 64), t))

		tmpName[1] = "bucket_le_inf"
		ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatUint(count, 10), t))

		for i, b := range h.buckets {
			floor := math.Floor(b)
			tmpBucketName[1] = strconv.FormatInt(int64(floor), 10)
			if floor-b < 0 {
				tmpBucketName[2] = strconv.FormatInt(int64(b*100-floor*100), 10)
				tmpName[1] = strings.Join(tmpBucketName, "_")
			} else {
				tmpName[1] = strings.Join(tmpBucketName[:2], "_")
			}

			ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatUint(values[i], 10), t))
		}

	}
	return ms
}

//Collapse collapse
func (h *Histogram) Collapse() map[string]observe.HistogramObserve {

	n := make(map[string]observe.HistogramObserve)
	h.locker.Lock()
	histograms := h.histograms
	h.histograms = n
	h.locker.Unlock()

	return histograms
}


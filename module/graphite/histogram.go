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

//Histogram histogram
type Histogram struct {
	metricKey    MetricKey
	keyHistogram *_KeyHistogram
}

//NewHistogram new histogram
func NewHistogram(metricKey MetricKey, buckets []float64) *Histogram {
	return &Histogram{
		metricKey:    metricKey,
		keyHistogram: newKeyHistogram(buckets),
	}
}

//Observe observe
func (h *Histogram) Observe(value float64, labels diting.Labels) {
	key := h.metricKey.Key(labels, "")
	h.keyHistogram.Observe(key, value)
}

//Metrics metrics
func (h *Histogram) Metrics() []graphite.Metric {
	all := h.keyHistogram.Collapse()

	if len(all) == 0 {
		return nil
	}
	keySize := len(all)
	size := keySize * (len(h.keyHistogram.buckets) + 5)
	ms := make([]graphite.Metric, 0, size)
	t := time.Now().Unix()
	tmpName := make([]string, 2)
	tmpBucketName := make([]string, 3)
	tmpBucketName[0] = "bucket_le"
	overIndex := len(h.keyHistogram.buckets)

	for k, v := range all {
		tmpName[0] = k

		tmpName[1] = "count"
		ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.Itoa(v.count), t))
		tmpName[1] = "max"
		ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatFloat(v.max, 'f', 2, 64), t))
		tmpName[1] = "min"
		ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatFloat(v.min, 'f', 2, 64), t))

		tmpName[1] = "sum"
		ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatFloat(v.sum, 'f', 2, 64), t))

		tmpName[1] = "bucket_le_inf"
		ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatInt(v.list[overIndex], 10), t))

		for i, b := range h.keyHistogram.buckets {
			floor := math.Floor(b)
			tmpBucketName[1] = strconv.FormatInt(int64(floor), 10)
			if floor-b < 0 {
				tmpBucketName[2] = strconv.FormatInt(int64(b*100-floor*100), 10)
				tmpName[1] = strings.Join(tmpBucketName, "_")
			} else {
				tmpName[1] = strings.Join(tmpBucketName[:2], "_")
			}

			ms = append(ms, graphite.NewMetric(strings.Join(tmpName, "."), strconv.FormatInt(v.list[i], 10), t))
		}

	}
	return ms
}

type _KeyHistogram struct {
	buckets    []float64
	histograms map[string]observe.Histogram
	locker     sync.Mutex
}

func newKeyHistogram(buckets []float64) *_KeyHistogram {
	return &_KeyHistogram{
		buckets:    buckets,
		histograms: make(map[string]observe.Histogram),
		locker:     sync.Mutex{},
	}
}

//HistogramValue histogramValue
type HistogramValue struct {
	list  []int64
	max   float64
	min   float64
	count int
	sum   float64
}

//Collapse collapse
func (k *_KeyHistogram) Collapse() map[string]HistogramValue {

	new := make(map[string]observe.Histogram)
	k.locker.Lock()
	histograms := k.histograms
	k.histograms = new
	k.locker.Unlock()

	values := make(map[string]HistogramValue)
	for k, hm := range histograms {
		col, sum, max, min, count := hm.Collapse()
		values[k] = HistogramValue{
			list:  col,
			max:   max,
			min:   min,
			count: count,
			sum:   sum,
		}
	}

	return values
}
func (k *_KeyHistogram) get(key string) observe.Histogram {
	k.locker.Lock()

	h, has := k.histograms[key]
	if !has {
		h = observe.NewHistogram(k.buckets)
		k.histograms[key] = h
	}

	k.locker.Unlock()
	return h
}

//Observe observe
func (k *_KeyHistogram) Observe(key string, value float64) {
	k.get(key).Observe(value)
}

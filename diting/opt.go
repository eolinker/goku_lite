package diting

import (
	"time"
)

//SeparatorByte separatorByte
const SeparatorByte byte = 255

var separatorByteSlice = []byte{SeparatorByte}

//Opts opts
type Opts struct {
	ID uint64

	LabelNames []string
	// Namespace, Subsystem, and Name are components of the fully-qualified
	// name of the Metric (created by joining these components with
	// "_"). Only Name is mandatory, the others merely help structuring the
	// name. Note that the fully-qualified name of the metric must be a
	// valid Prometheus metric name.
	Namespace string
	Subsystem string
	Name      string

	// Help provides information about this metric.
	//
	// Metrics with the same fully-qualified name must have the same Help
	// string.
	Help string

	// ConstLabels are used to attach fixed labels to this metric. Metrics
	// with the same fully-qualified name must have the same label names in
	// their ConstLabels.
	//
	// ConstLabels are only used rarely. In particular, do not use them to
	// attach the same labels to all your metrics. Those use cases are
	// better covered by target labels set by the scraping Prometheus
	// server, or by one specific metric (e.g. a build_info or a
	// machine_role metric). See also
	// https://prometheus.io/docs/instrumenting/writing_exporters/#target-labels,-not-static-scraped-labels
	ConstLabels Labels
}

//CounterOpts couterOpts
type CounterOpts Opts

//GaugeOpts gaugeOpts
type GaugeOpts Opts

//HistogramOpts histogramOpts
type HistogramOpts struct {
	ID uint64

	LabelNames []string
	// Namespace, Subsystem, and Name are components of the fully-qualified
	// name of the HistogramObserve (created by joining these components with
	// "_"). Only Name is mandatory, the others merely help structuring the
	// name. Note that the fully-qualified name of the HistogramObserve must be a
	// valid Prometheus metric name.
	Namespace string
	Subsystem string
	Name      string

	// Help provides information about this HistogramObserve.
	//
	// Metrics with the same fully-qualified name must have the same Help
	// string.
	Help string

	// ConstLabels are used to attach fixed labels to this metric. Metrics
	// with the same fully-qualified name must have the same label names in
	// their ConstLabels.
	//
	// ConstLabels are only used rarely. In particular, do not use them to
	// attach the same labels to all your metrics. Those use cases are
	// better covered by target labels set by the scraping Prometheus
	// server, or by one specific metric (e.g. a build_info or a
	// machine_role metric). See also
	// https://prometheus.io/docs/instrumenting/writing_exporters/#target-labels,-not-static-scraped-labels
	ConstLabels Labels

	// buckets defines the buckets into which observations are counted. Each
	// element in the slice is the upper inclusive bound of a bucket. The
	// values must be sorted in strictly increasing order. There is no need
	// to add a highest bucket with +Inf bound, it will be added
	// implicitly. The default value is DefBuckets.
	Buckets []float64
}

//SummaryOpts summaryOpts
type SummaryOpts struct {
	ID uint64

	LabelNames []string
	// Namespace, Subsystem, and Name are components of the fully-qualified
	// name of the Summary (created by joining these components with
	// "_"). Only Name is mandatory, the others merely help structuring the
	// name. Note that the fully-qualified name of the Summary must be a
	// valid Prometheus metric name.
	Namespace string
	Subsystem string
	Name      string

	// Help provides information about this Summary.
	//
	// Metrics with the same fully-qualified name must have the same Help
	// string.
	Help string

	// ConstLabels are used to attach fixed labels to this metric. Metrics
	// with the same fully-qualified name must have the same label names in
	// their ConstLabels.
	//
	// Due to the way a Summary is represented in the Prometheus text format
	// and how it is handled by the Prometheus server internally, “quantile”
	// is an illegal label name. Construction of a Summary or SummaryVec
	// will panic if this label name is used in ConstLabels.
	//
	// ConstLabels are only used rarely. In particular, do not use them to
	// attach the same labels to all your metrics. Those use cases are
	// better covered by target labels set by the scraping Prometheus
	// server, or by one specific metric (e.g. a build_info or a
	// machine_role metric). See also
	// https://prometheus.io/docs/instrumenting/writing_exporters/#target-labels,-not-static-scraped-labels
	ConstLabels Labels

	// Objectives defines the quantile rank estimates with their respective
	// absolute error. If Objectives[q] = e, then the value reported for q
	// will be the φ-quantile value for some φ between q-e and q+e.  The
	// default value is an empty map, resulting in a summary without
	// quantiles.
	Objectives map[float64]float64

	// MaxAge defines the duration for which an observation stays relevant
	// for the summary. Must be positive. The default value is DefMaxAge.
	MaxAge time.Duration

	// AgeBuckets is the number of buckets used to exclude observations that
	// are older than MaxAge from the summary. A higher number has a
	// resource penalty, so only increase it if the higher resolution is
	// really required. For very high observation rates, you might want to
	// reduce the number of age buckets. With only one age bucket, you will
	// effectively see a complete reset of the summary each time MaxAge has
	// passed. The default value is DefAgeBuckets.
	AgeBuckets uint32

	// BufCap defines the default sample stream buffer size.  The default
	// value of DefBufCap should suffice for most uses. If there is a need
	// to increase the value, a multiple of 500 is recommended (because that
	// is the internal buffer size of the underlying package
	// "github.com/bmizerany/perks/quantile").
	BufCap uint32
}

//NewCounterOpts new counterOpts
func NewCounterOpts(namespace string, subsystem string, name string, help string, constLabels Labels, labelNames []string) *CounterOpts {
	return &CounterOpts{
		ID:          GetID(),
		LabelNames:  labelNames,
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}
}

//NewGaugeOpts new gaugeOpts
func NewGaugeOpts(namespace string, subsystem string, name string, help string, constLabels Labels, labelNames []string) *GaugeOpts {
	return &GaugeOpts{
		ID:          GetID(),
		LabelNames:  labelNames,
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}
}

//NewHistogramOpts new histogramOpts
func NewHistogramOpts(namespace string, subsystem string, name string, help string, constLabels Labels, labelNames []string, buckets []float64) *HistogramOpts {
	return &HistogramOpts{
		ID:          GetID(),
		LabelNames:  labelNames,
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
		Buckets:     buckets,
	}
}

//NewSummaryOpts new summaryOpts
func NewSummaryOpts(namespace string, subsystem string, name string, help string, constLabels Labels, labelNames []string, objectives map[float64]float64, maxAge time.Duration, ageBuckets uint32, bufCap uint32) *SummaryOpts {
	return &SummaryOpts{
		ID:          GetID(),
		LabelNames:  labelNames,
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
		Objectives:  objectives,
		MaxAge:      maxAge,
		AgeBuckets:  ageBuckets,
		BufCap:      bufCap,
	}
}

package prometheus

import (
	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/prometheus/client_golang/prometheus"
)

//ReadCounterOpts read CounterOpts
func ReadCounterOpts(opts *diting.CounterOpts) prometheus.CounterOpts {

	return prometheus.CounterOpts{
		Namespace: opts.Namespace,
		Subsystem: opts.Subsystem,
		Name:      opts.Name,
		Help:      opts.Help,
		//ConstLabels: prometheus.Labels( opts.ConstLabels),
	}
}

//ReadGaugeOpts read GaugeOpts
func ReadGaugeOpts(opts *diting.GaugeOpts) prometheus.GaugeOpts {

	return prometheus.GaugeOpts{
		Namespace: opts.Namespace,
		Subsystem: opts.Subsystem,
		Name:      opts.Name,
		Help:      opts.Help,
		//ConstLabels: prometheus.Labels( opts.ConstLabels),
	}
}

//ReadHistogramOpts read HistogramOpts
func ReadHistogramOpts(opts *diting.HistogramOpts) prometheus.HistogramOpts {

	return prometheus.HistogramOpts{
		Namespace: opts.Namespace,
		Subsystem: opts.Subsystem,
		Name:      opts.Name,
		Help:      opts.Help,
		//ConstLabels: prometheus.Labels(opts.ConstLabels),
		Buckets: opts.Buckets,
	}
}

//ReadSummaryOpts read SummaryOpts
func ReadSummaryOpts(opts *diting.SummaryOpts) prometheus.SummaryOpts {

	return prometheus.SummaryOpts{
		Namespace: opts.Namespace,
		Subsystem: opts.Subsystem,
		Name:      opts.Name,
		Help:      opts.Help,
		//ConstLabels: prometheus.Labels(opts.ConstLabels),
		Objectives: opts.Objectives,
		MaxAge:     opts.MaxAge,
		AgeBuckets: opts.AgeBuckets,
		BufCap:     opts.BufCap,
	}
}

package monitor

import (
	"github.com/eolinker/goku-api-gateway/diting"
	goku_labels "github.com/eolinker/goku-api-gateway/goku-labels"
)

var (
	//APICount diting.Counter
	APIMonitor diting.Histogram
	ProxyMonitor diting.Histogram
)

func initCollector(constLabels diting.Labels)  {
	//apiLabelNames  := []string{
	//	goku_labels.Cluster,
	//	goku_labels.Instance,
	//	goku_labels.API,
	//	goku_labels.Strategy,
	//	goku_labels.Status,
	//}
	//apiCounterOpt:= diting.NewCounterOpts("goku","api","count","api 请求计数",constLabels,apiLabelNames)
	//APICount = diting.NewCounter(apiCounterOpt)

	buckets:=[]float64{5, 25, 50, 100, 200,400, 600,800,1000, 2500,5000}
	//buckets:=[]float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}
	apiDelayLabelNames := []string{
		goku_labels.Cluster,
		goku_labels.Instance,
		goku_labels.API,
		goku_labels.Strategy,
		goku_labels.Status,
	}
	apiHistogramOpt := diting.NewHistogramOpts("goku","","api","api整体请求统计",constLabels,apiDelayLabelNames,buckets)
	APIMonitor = diting.NewHistogram(apiHistogramOpt)


	proxyDelayLabelNames := []string{
		goku_labels.Cluster,
		goku_labels.Instance,
		goku_labels.Proto,
		goku_labels.Host,
		goku_labels.Path,
		goku_labels.Method,
	}

	proxyMonitorOpt := diting.NewHistogramOpts("goku","","proxy","转发统计",constLabels,proxyDelayLabelNames,buckets)
	ProxyMonitor = diting.NewHistogram(proxyMonitorOpt)

}
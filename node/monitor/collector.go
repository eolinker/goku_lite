package monitor

import (
	"github.com/eolinker/goku-api-gateway/diting"
	goku_labels "github.com/eolinker/goku-api-gateway/goku-labels"
)

var (
	//APIMonitor diting.APIMonitor
	APIMonitor diting.Histogram
	//ProxyMonitor diting.Histogram
	ProxyMonitor diting.Histogram
)

func initCollector(constLabels diting.Labels) {
	//apiLabelNames  := []string{
	//	goku_labels.Cluster,
	//	goku_labels.Instance,
	//	goku_labels.API,
	//	goku_labels.Strategy,
	//	goku_labels.Status,
	//}
	//apiCounterOpt:= diting.NewCounterOpts("goku","api","count","api 请求计数",constLabels,apiLabelNames)
	//APICount = diting.NewCounter(apiCounterOpt)

	apiHistogramOpt := diting.NewHistogramOpts(goku_labels.Namespace, goku_labels.Subsystem, goku_labels.APIName, "api整体请求统计", constLabels, goku_labels.APIDelayLabelNames, goku_labels.APIBuckets)
	APIMonitor = diting.NewHistogram(apiHistogramOpt)

	proxyMonitorOpt := diting.NewHistogramOpts(goku_labels.Namespace, goku_labels.Subsystem, goku_labels.ProxyName, "转发统计", constLabels, goku_labels.ProxyDelayLabelNames, goku_labels.ProxyBuckets)
	ProxyMonitor = diting.NewHistogram(proxyMonitorOpt)

}

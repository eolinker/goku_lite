package goku_labels // constlabel
// api labels
// proxy
const (
	Cluster  = "cluster"
	Instance = "instance"
	//Namespace 命名空间
	Namespace = "goku"
	//Subsystem subSystem
	Subsystem = ""
	//APIName apiName
	APIName = "api"
	//ProxyName proxyName
	ProxyName = "proxy"

	API      = "api"
	Strategy = "strategy"
	Status   = "status"
	Proto    = "proto"
	Method   = "method"
	Host     = "host"
	Path     = "path"
)

var (
	//APIBuckets apiBuckets
	APIBuckets = []float64{5, 25, 50, 100, 200, 400, 600, 800, 1000, 2500, 5000}
	//ProxyBuckets proxyBuckets
	ProxyBuckets = []float64{5, 25, 50, 100, 200, 400, 600, 800, 1000, 2500, 5000}

	//APIDelayLabelNames apiDelayLabelNames
	APIDelayLabelNames = []string{
		Cluster,
		Instance,
		API,
		Strategy,
		Status,
	}
	//ProxyDelayLabelNames proxyDelayLabelNames
	ProxyDelayLabelNames = []string{
		Cluster,
		Instance,
		API,
		Strategy,
		Proto,
		Host,
		Path,
		Method,
		Status,
	}
)

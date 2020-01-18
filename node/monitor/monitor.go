package monitor

import (
	"strings"
	"sync"

	"github.com/eolinker/goku-api-gateway/diting"
	goku_labels "github.com/eolinker/goku-api-gateway/goku-labels"
)

var (
	once = sync.Once{}
)

//Init init
func Init(cluster string, instance string) {

	once.Do(func() {
		constLabels := make(diting.Labels)
		constLabels[goku_labels.Cluster] = cluster
		constLabels[goku_labels.Instance] = strings.ReplaceAll(instance, ".", "_")

		initCollector(constLabels)
	})
}

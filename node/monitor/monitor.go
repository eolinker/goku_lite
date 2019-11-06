package monitor

import (
	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/eolinker/goku-api-gateway/goku-labels"
	"strings"
	"sync"
)


var (
	once =sync.Once{}

)

func Init(cluster string,instance string)  {

	once.Do(func() {
		constLabels := make(diting.Labels)
		constLabels[goku_labels.Cluster] = cluster
		constLabels[goku_labels.Instance] = strings.ReplaceAll(instance,".","_")

		initCollector(constLabels)
	})
}


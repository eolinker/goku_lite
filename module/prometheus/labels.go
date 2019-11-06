package prometheus

import (
	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/prometheus/client_golang/prometheus"
)

//ReadLabels readLabels
func ReadLabels(labels diting.Labels) prometheus.Labels {
	return prometheus.Labels(labels)
}

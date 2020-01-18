package graphite

import (
	"strings"
	"unicode"

	"github.com/eolinker/goku-api-gateway/diting"
	"github.com/marpaia/graphite-golang"
)

//Metrics Metrics
type Metrics interface {
	Metrics() []graphite.Metric
}

//MetricKey MetricKey
type MetricKey interface {
	Key(labels diting.Labels, valueType string) string
}

type _MetricKey struct {
	name       string
	labelNames []string
}

//NewMetricKey new MetricKey
func NewMetricKey(name string, labelNames []string) MetricKey {
	return &_MetricKey{name: name, labelNames: labelNames}
}

//Key key
func (m *_MetricKey) Key(labels diting.Labels, valueType string) string {
	tmp := make([]string, 0, len(m.labelNames)+2)
	tmp = append(tmp, m.name)

	for _, name := range m.labelNames {
		labelValue := labels[name]
		tmp = append(tmp, formatLabelValue(labelValue))
	}

	if valueType != "" {
		tmp = append(tmp, formatLabelValue(valueType))
	}
	return strings.Join(tmp, ".")
}

const rep = '_'

//formatLabelValue 将label value的所有字母、数字转换成 _
func formatLabelValue(value string) string {
	s := []rune(value)
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			s[i] = rep
		}
	}
	return string(s)
}

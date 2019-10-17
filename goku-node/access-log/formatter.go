package access_log

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	access_field "github.com/eolinker/goku-api-gateway/server/access-field"
	"github.com/sirupsen/logrus"
)

//AccessLogFormatter access日志格式器
type AccessLogFormatter struct {
	fields          []access_field.AccessFieldKey
	locker          sync.RWMutex
	TimestampFormat string
}

//SetFields 设置域
func (f *AccessLogFormatter) SetFields(fields []access_field.AccessFieldKey) {
	f.locker.Lock()
	f.fields = fields
	f.locker.Unlock()
}

//Format 格式化
func (f *AccessLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = DefaultTimeStampFormatter
	}

	data := entry.Data
	data[access_field.TimeLocal] = entry.Time.Format(timestampFormat)
	data[access_field.TimeIso8601] = entry.Time.Format(TimeIso8601Formatter)

	msec := entry.Time.UnixNano() / int64(time.Millisecond)
	data[access_field.Msec] = fmt.Sprintf("%d.%d", msec/1000, msec%1000)

	requestTIme := data[access_field.RequestTime].(time.Duration)
	data[access_field.RequestTime] = fmt.Sprintf("%dms", requestTIme/time.Millisecond)

	for _, key := range f.fields {
		b.WriteByte('\t')
		if v, has := data[key.Key()]; has {
			f.appendValue(b, v)
		} else {
			f.appendValue(b, "-")
		}
	}
	b.WriteByte('\n')
	p := b.Bytes()
	return p[1:], nil
}

func (f *AccessLogFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	//if !f.needsQuoting(stringVal) {
	b.WriteString(stringVal)
	//} else {
	//	b.WriteString(fmt.Sprintf("%q", stringVal))
	//}
}

//NewAccessLogFormatter 创建AccessLogFormatter
func NewAccessLogFormatter(fields []access_field.AccessFieldKey) *AccessLogFormatter {
	return &AccessLogFormatter{fields: fields}
}

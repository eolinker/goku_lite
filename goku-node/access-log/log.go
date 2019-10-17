package access_log

import (
	"time"

	log "github.com/eolinker/goku-api-gateway/goku-log"
	access_field "github.com/eolinker/goku-api-gateway/server/access-field"
	"github.com/sirupsen/logrus"
)

var (
	logger    *logrus.Logger
	formatter *AccessLogFormatter
	writer    *log.FileWriterByPeriod
)

//Fields 域
type Fields = logrus.Fields

//Log log
func Log(fields Fields) {
	if logger == nil {
		return
	}
	logger.WithFields(fields).Info()
}

//SetFields 设置access域
func SetFields(fields []access_field.AccessFieldKey) {
	if formatter == nil {
		formatter = NewAccessLogFormatter(fields)
	} else {
		formatter.SetFields(fields)
	}
}

//SetOutput 设置输出
func SetOutput(enable bool, dir, file string, period log.LogPeriod, expire int) {

	if enable {
		if writer == nil {
			writer = log.NewFileWriteBytePeriod()
		}

		writer.Set(dir, file, period, time.Duration(expire)*time.Hour*24)
		writer.Open()
		if logger == nil {
			logger = logrus.New()
			logger.SetFormatter(formatter)
			logger.SetOutput(writer)
			logger.SetLevel(logrus.InfoLevel)
		}

	} else {
		if writer != nil {
			writer.Close()
		}

	}

}

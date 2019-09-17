package access_log

import (
	log "github.com/eolinker/goku/goku-log"
	access_field "github.com/eolinker/goku/server/access-field"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	logger    *logrus.Logger
	formatter *AccessLogFormatter
	writer    *log.FileWriterByPeriod
)

//func InitLogger(enable bool,fields []string,dir,file string,period log.LogPeriod) {
//
//	SetFields(fields)
//	SetOutput(enable,dir,file,period)
//
//}
type Fields = logrus.Fields

func Log(fields Fields) {
	if logger == nil {
		return
	}
	logger.WithFields(fields).Info()
}

func SetFields(fields []access_field.AccessFieldKey) {
	if formatter == nil {
		formatter = NewAccessLogFormatter(fields)
	} else {
		formatter.SetFields(fields)
	}
}
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

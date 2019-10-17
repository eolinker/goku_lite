package goku_log

import (
	"os"

	"github.com/sirupsen/logrus"
)

//StartDebug 开启debug模式
func StartDebug() {
	logger.AddHook(new(debugHook))
}

type debugHook struct {
}

func (h *debugHook) Levels() []logrus.Level {
	return []Level{
		TraceLevel,
		DebugLevel,
		InfoLevel,
		WarnLevel,
		ErrorLevel,
		FatalLevel,
		PanicLevel,
	}
}

func (h *debugHook) Fire(entry *logrus.Entry) error {
	s, e := logger.Formatter.Format(entry)
	if e != nil {
		return nil
	}
	os.Stdout.Write(s)
	return nil
}

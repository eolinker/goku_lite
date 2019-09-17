package access_log

import "time"

const (
	DefaultTimeStampFormatter ="[2006-01-02 15:04:05]"
	TimeIso8601Formatter = "["+time.RFC3339+"]"
)

package access_log

import "time"

const (
	//DefaultTimeStampFormatter 时间戳默认格式化字符串
	DefaultTimeStampFormatter = "[2006-01-02 15:04:05]"
	//TimeIso8601Formatter iso8601格式化
	TimeIso8601Formatter = "[" + time.RFC3339 + "]"
)

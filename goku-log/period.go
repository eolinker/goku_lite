package goku_log

import (
	"fmt"
	"strings"
)

//LogPeriod 日志周期
type LogPeriod interface {
	String() string
	FormatLayout() string
}

//LogPeriodType 日志周期类型
type LogPeriodType int

//ParsePeriod 解析周期
func ParsePeriod(v string) (LogPeriod, error) {
	switch strings.ToLower(v) {
	case "month":
		return PeriodMonth, nil
	case "day":
		return PeriodDay, nil
	case "hour":
		return PeriodHour, nil
	}

	return nil, fmt.Errorf("not a valid period: %q", v)
}
func (period LogPeriodType) String() string {
	switch period {
	case PeriodMonth:
		return "month"
	case PeriodDay:
		return "day"
	case PeriodHour:
		return "hour"
	default:
		return "unknown"
	}
}

const (
	//PeriodMonth 月
	PeriodMonth LogPeriodType = iota
	//PeriodDay 日
	PeriodDay
	//PeriodHour 时
	PeriodHour
)

//FormatLayout 格式化
func (period LogPeriodType) FormatLayout() string {
	switch period {
	case PeriodHour:
		{
			return "2006-01-02-15"
		}
	case PeriodDay:
		{
			return "2006-01-02"
		}
	case PeriodMonth:
		{
			return "2006-01"
		}
	default:
		return "2006-01-02-15"
	}
}

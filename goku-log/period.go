package goku_log

import (
	"fmt"
	"strings"
)

type LogPeriod interface {
	String() string
	FormatLayout() string
}
type LogPeriodType int

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
	PeriodMonth LogPeriodType = iota
	PeriodDay
	PeriodHour
)

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

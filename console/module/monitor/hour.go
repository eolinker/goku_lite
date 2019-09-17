package monitor

import (
	"strconv"
	"time"
)

func genHour(beginTime, endTime string, period int) (int, int) {
	startHour := 0
	endHour, _ := strconv.Atoi(time.Now().Add(time.Hour).Format("2006010215"))

	switch period {
	case 3:
		{

			bt, e := time.Parse("2006-01-02", beginTime)
			if e == nil {
				startHour, _ = strconv.Atoi(bt.Format("2006010215"))
			}
			et, e := time.Parse("2006-01-02", endTime)
			if e == nil {
				et.Add(time.Hour*24 - time.Minute)
				endHour, _ = strconv.Atoi(et.Format("2006010215"))
			}
		}
	case 2:
		startHour, _ = strconv.Atoi(time.Now().Add(- time.Hour * 24 * 7).Format("2006010215"))
	case 1:
		startHour, _ = strconv.Atoi(time.Now().Add(- time.Hour * 24 * 3).Format("2006010215"))
	default:
		startHour, _ = strconv.Atoi(time.Now().Format("2006010200"))
	}

	return startHour, endHour
}

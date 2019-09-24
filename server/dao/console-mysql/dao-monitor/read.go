package dao_monitor

import monitor_key "github.com/eolinker/goku-api-gateway/server/monitor/monitor-key"

func read(s SCAN, args ...interface{}) (monitor_key.MonitorValues, error) {
	v := monitor_key.MakeValue()
	vp := make([]interface{}, 0, monitor_key.MonitorKeyTypeSize+len(args))

	vp = append(vp, args...)

	for i := range v {
		vp = append(vp, &v[i])
	}
	err := s.Scan(vp...)
	if err != nil {
		return v, err
	}
	return v, nil
}

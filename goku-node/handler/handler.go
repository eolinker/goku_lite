package handler

import (
	"net/http"
)

//Entry entry
type Entry struct {
	Pattern     string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

func init() {

}

//Handler handler
func Handler() []Entry {

	return []Entry{
		{
			Pattern: "/goku-update", HandlerFunc: gokuUpdate,
		},
		{
			Pattern: "/goku-check_update", HandlerFunc: gokuCheckUpdate},
		{
			Pattern: "/goku-check_plugin", HandlerFunc: gokuCheckPlugin},
		{
			Pattern: "/goku-monitor", HandlerFunc: gokuMonitor},
	}
}

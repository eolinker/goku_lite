package handler

import "net/http"

func gokuMonitor(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	return
}

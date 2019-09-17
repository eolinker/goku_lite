package handler

import (
	"encoding/json"
	"github.com/eolinker/goku/goku-node/manager/updater"
	"net/http"
)

func gokuUpdate(w http.ResponseWriter, r *http.Request) {
	updater.Update()
}
func gokuCheckUpdate(w http.ResponseWriter, r *http.Request) {
	resultInfo := map[string]interface{}{
		"type":       "update",
		"statusCode": "000000",
		"version":    Version,
	}
	resultStr, _ := json.Marshal(resultInfo)

	w.WriteHeader(200)
	_, _ = w.Write(resultStr)
	return
}

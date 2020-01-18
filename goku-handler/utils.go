package goku_handler

import (
	"encoding/json"
	log "github.com/eolinker/goku-api-gateway/goku-log"
	"net/http"
)



//WriteError 返回错误
func WriteError(w http.ResponseWriter, statusCode, resultType, resultDesc string, resuleErr error) {

	ret := map[string]interface{}{
		"type":       resultType,
		"statusCode": statusCode,
		"resultDesc": resultDesc,
	}

	if resuleErr != nil {
		log.Info(resuleErr)
	}

	data, err := json.Marshal(ret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.WithFields(ret).Debug("write error:", err)
	}

	w.Write(data)
	log.WithFields(ret).Debug("write error:", err)

}


package admin

import "net/http"

//noinspection GoTypesCompatibility
func router() http.Handler {
	serverHandler := http.NewServeMux()
	serverHandler.HandleFunc("/register", Register)
	serverHandler.HandleFunc("/node/heartbeat", heartbead)
	serverHandler.HandleFunc("/node/stop", stopNode)
	serverHandler.HandleFunc("/alert/msg/add", AddAlertMsg)
	return serverHandler
}

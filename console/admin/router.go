package admin

import "net/http"

//noinspection GoTypesCompatibility
func router() http.Handler {
	serverHandler := http.NewServeMux()

	serverHandler.HandleFunc("/version/config/get", GetVersionConfig)

	return serverHandler
}

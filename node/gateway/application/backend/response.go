package backend

import "net/http"

type BackendResponse struct {
	Method string
	Protocol string
	TargetUrl string
	FinalTargetServer string
	RetryTargetServers []string
	BodyOrg []byte
	Header http.Header
	Body interface{}
	StatusCode int
	Status string
}


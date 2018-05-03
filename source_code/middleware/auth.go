package middleware

import (
	"goku-ce/conf"
	"net/http"
	"strings"
	"encoding/base64"
)

func Auth(c conf.StrategyInfo,res http.ResponseWriter, req *http.Request) (bool,string) {
	if c.Auth == "basic" {
		authStr := []byte(c.BasicUserName + ":" + c.BasicUserPassword)
		authorization := "Basic " + base64.StdEncoding.EncodeToString(authStr)
		auth := strings.Join(req.Header["Authorization"],", ")
		if authorization != auth {
			return false, "Error username or userpassword"
		}
	} else if c.Auth == "apikey" {
		apiKey := strings.Join(req.Header["Apikey"],", ")
		if c.ApiKey != apiKey {
			return false,"Error apiKey"
		}
	}
	return true,""
}
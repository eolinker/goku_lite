package middleware

import (
	"goku-ce/goku"
	"net/http"
	"strings"
	"encoding/base64"
)

func Auth(context *goku.Context,res http.ResponseWriter, req *http.Request) (bool,string) {
	c := context.StrategyInfo
	if strings.ToLower(c.Auth) == "basic" {
		authStr := []byte(c.BasicUserName + ":" + c.BasicUserPassword)
		authorization := "Basic " + base64.StdEncoding.EncodeToString(authStr)
		auth := strings.Join(req.Header["Authorization"],", ")
		if authorization != auth {
			return false, "Username or UserPassword Error"
		}
	} else if strings.ToLower(c.Auth) == "apikey" {
		apiKey := strings.Join(req.Header["Apikey"],", ")
		if c.ApiKey != apiKey {
			return false,"Invalid ApiKey"
		}
	}
	return true,""
}
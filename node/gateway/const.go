package gateway

var (
	authNames = map[string]string{
		"Oauth2": "goku-oauth2_auth",
		"Apikey": "goku-apikey_auth",
		"Basic":  "goku-basic_auth",
		"Jwt":    "goku-jwt_auth",
	}
)

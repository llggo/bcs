package web

import (
	"net/http"
	"strings"
)

const bearerHeader = "Bearer "
const accessToken = "access_token"

func GetToken(r *http.Request) string {
	var authHeader = r.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, bearerHeader) {
		return strings.TrimPrefix(authHeader, bearerHeader)
	}
	return r.URL.Query().Get(accessToken)
}

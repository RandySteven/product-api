package middleware

import "net/http"

// func MiddlewareHandler(h http.Handler) http.Handler {
// 	// return http.HandlerFunc()
// }

func MiddlewareAuth(res http.ResponseWriter, req *http.Request) bool {
	// username, password, ok := req.BasicAuth()
	return false
}

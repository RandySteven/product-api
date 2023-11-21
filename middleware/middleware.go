package middleware

import "net/http"

func MiddlewareAuth(next http.Handler) http.Handler {
	// username, password, ok := req.BasicAuth()
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if Authenticated(req) {
			http.Redirect(res, req, "/users/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(res, req)
	})
}

func Authenticated(req *http.Request) bool {
	return true
}

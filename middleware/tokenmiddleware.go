package middleware

import (
	"net/http"
)

// TokenMiddleware : return http.Handler
func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("auth-token")
		if len(tokenString) == 0 {
			ErrorResponse(w, http.StatusUnauthorized, "Authentication failure")
			return
		}
		// call db to validate the token Todo: use gorilla context to set param in context

		r.Header.Set("userId", "33995")
		next.ServeHTTP(w, r)
	})
}

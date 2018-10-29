package middleware

import (
	"fmt"
	"net/http"
)

// PanicRecoveryMiddleware : internal server error on error
func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// log the error
				fmt.Println(rec)
				// write the error response
				ErrorResponse(w, http.StatusInternalServerError, "Internal Error")
				return
			}
		}()

		next.ServeHTTP(w, r)
	})
}

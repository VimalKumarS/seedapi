package middleware

import (
	"net/http"
	"seedapi/model"
	util "seedapi/util"

)

// ErrorHandlerMiddleware : internal server error on error
func ErrorHandlerMiddleware(next model.HttpHandlerWithError) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Execute the final handler, and deal with errors
		var status int
		err := next(w, r)
		if err != nil {
			errorType := util.GetType(err)
			switch errorType {
			case util.BadRequest:
				status = http.StatusBadRequest
			case util.NotFound:
				status = http.StatusNotFound
			default:
				status = http.StatusInternalServerError
			}
			util.Logger.Error("Error", (util.GetCustomError(err))...)
			ErrorResponse(w, status, util.GetErrorContext(err))
		}
	})
}

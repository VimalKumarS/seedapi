package route

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"seedapi/middleware"
	"seedapi/model"
	
)

// Routes list of route
type Routes []model.Route

// Middleware handler func
type Middleware func(http.HandlerFunc) http.HandlerFunc

type heartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}


func chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

//LoadRouter : load all routes defined in routes.go
func LoadRouter(config *model.AppSetting) *mux.Router {
	
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range bundleRoute(config) {
		var handler http.Handler

		//handler = route.HandlerFunc
		
		if route.Protected {
			//Todo: chain the middle ware
			handler = middleware.PanicRecoveryMiddleware(
										(middleware.TokenMiddleware(
											middleware.ErrorHandlerMiddleware(route.HandlerFunc))))

			//Chain(route.HandlerFunc,middleware.TokenMiddleware,middleware.PanicRecoveryMiddleware)
		}else{
			handler = middleware.PanicRecoveryMiddleware(middleware.ErrorHandlerMiddleware(route.HandlerFunc))
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}


func bundleRoute(config *model.AppSetting) Routes{

	route := []model.Route{
		{
			"heartbeat",
			"GET",
			"/heartbeat",
			heartbeat,
			false,
		},
	}
	route = append(route, 
		(SearchRoutes(config))...)  // append all route
	
	return route
}


// Heart beat response
func heartbeat(w http.ResponseWriter, r *http.Request) error {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
	return nil
}
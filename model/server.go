package model

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

// Server : struct
type Server struct {
	config *AppSetting
}

// CreateServer :  create server config
func CreateServer(config *AppSetting) *Server {
	return &Server{
		config: config,
	}
}

// Run : Start the server
func (s *Server) Run(router *mux.Router) {

	//enabling cors if requried
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":"+s.config.Port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}

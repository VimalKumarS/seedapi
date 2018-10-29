package model

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	if s.config.EnableCors {
		allowedOrigins := handlers.AllowedOrigins([]string{"*"})
		allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
		fmt.Println("Server running on port:" + s.config.Port)
		log.Fatal(http.ListenAndServe(":"+s.config.Port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
	}else{
		fmt.Println("Server running on port:" + s.config.Port)
		log.Fatal(http.ListenAndServe(":"+s.config.Port, router))
	}
}

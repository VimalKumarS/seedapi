package model

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Server : struct
type Server struct {
	config *AppSetting
	http.Server
	Logger *zap.Logger
}

// CreateServer :  create server config
func CreateServer(config *AppSetting,Logger *zap.Logger) *Server {
	return &Server{
		config: config,
		Server : http.Server{
            Addr:         config.Port,
            ReadTimeout:  10 * time.Second,
            WriteTimeout: 10 * time.Second,
		},
		Logger:Logger,
	}
}

// Run : Start the server
func (s *Server) Run(router *mux.Router) {

	//enabling cors if requried
	if s.config.EnableCors {
		allowedOrigins := handlers.AllowedOrigins([]string{"*"})
		allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
		fmt.Println("Server running on port:" + s.config.Port)
		s.Server.Handler = handlers.CORS(allowedOrigins, allowedMethods)(router)
		
	}else{
		fmt.Println("Server running on port:" + s.config.Port)
		//log.Fatal(http.ListenAndServe(":"+s.config.Port, router))
		s.Server.Handler = router
	}
	//s.Logger.Fatal(s.Server.ListenAndServe())
	go func() {
		if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Logger.Error(err.Error())
		}
	}()
}

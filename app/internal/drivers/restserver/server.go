package restserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ServerConfig struct {
	Address string `yaml:"address"`
}

type Server struct {
	Router  chi.Router
	Address string
}

func NewServerFromConfig(config *ServerConfig, router chi.Router) *Server {
	return &Server{
		Router:  router,
		Address: config.Address,
	}
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.Address, s.Router)
}

package restserver

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ServerConfig struct {
	Address string `yaml:"address"`
}

type Server struct {
	underlying http.Server
}

func NewServerFromConfig(config *ServerConfig, router chi.Router) *Server {
	return &Server{
		underlying: http.Server{
			Addr:    config.Address,
			Handler: router,
		},
	}
}

func (s *Server) Run() error {
	return s.underlying.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.underlying.Shutdown(ctx)
}

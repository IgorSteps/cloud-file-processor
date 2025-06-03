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

// Run runs the http server. It always returns a non-nil error.
// In case of Server.Shutdown it returns a ErrServerClosed error.
func (s *Server) Run() error {
	return s.underlying.ListenAndServe()
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.underlying.Shutdown(ctx)
}

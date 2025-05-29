package restserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router  chi.Router
	Address string
}

func NewServerFromConfig(r chi.Router) *Server {
	return &Server{
		Router:  r,
		Address: "localhost:3000",
	}
}

func (s *Server) Run() error {
	return http.ListenAndServe(s.Address, s.Router)
}

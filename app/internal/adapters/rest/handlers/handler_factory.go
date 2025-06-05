package handlers

import (
	"log/slog"
	"net/http"
)

type HandlerFactory struct {
	logger *slog.Logger
}

func NewHandlerFactory(logger *slog.Logger) *HandlerFactory {
	return &HandlerFactory{
		logger: logger,
	}
}

func (s *HandlerFactory) UploadFile() http.Handler {
	return NewUploadFile(s.logger)
}

package handlers

import (
	"log/slog"
	"net/http"
)

type HandlerFactory struct {
	backupLogger *slog.Logger
}

func NewHandlerFactory(backupLogger *slog.Logger) *HandlerFactory {
	return &HandlerFactory{
		backupLogger: backupLogger,
	}
}

func (s *HandlerFactory) UploadFile() http.Handler {
	return NewUploadFile(s.backupLogger)
}

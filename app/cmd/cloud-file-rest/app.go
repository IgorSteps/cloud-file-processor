package main

import (
	"app/internal/drivers/restserver"
	"log/slog"
)

type App struct {
	server *restserver.Server
	logger *slog.Logger
}

func NewApp(server *restserver.Server, logger *slog.Logger) *App {
	return &App{
		server: server,
		logger: logger,
	}
}

package main

import (
	"app/internal/drivers/otel"
	"app/internal/drivers/restserver"
	"log/slog"
)

type App struct {
	server                 *restserver.Server
	logger                 *slog.Logger
	tracerProviderShutdown otel.TracerProviderShutdown
}

func NewApp(server *restserver.Server, logger *slog.Logger, tracerProviderShutdown otel.TracerProviderShutdown) *App {
	return &App{
		server:                 server,
		logger:                 logger,
		tracerProviderShutdown: tracerProviderShutdown,
	}
}

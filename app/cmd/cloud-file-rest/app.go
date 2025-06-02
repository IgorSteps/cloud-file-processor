package main

import (
	"app/internal/drivers/otel"
	"app/internal/drivers/restserver"
	"log/slog"
	"time"
)

type App struct {
	config                 *Config
	server                 *restserver.Server
	logger                 *slog.Logger
	tracerProviderShutdown otel.TracerProviderShutdown
	shutdownTimeout        time.Duration
}

func NewApp(
	config *Config,
	server *restserver.Server,
	logger *slog.Logger,
	tracerProviderShutdown otel.TracerProviderShutdown,
) *App {
	return &App{
		config:                 config,
		server:                 server,
		logger:                 logger,
		tracerProviderShutdown: tracerProviderShutdown,
		shutdownTimeout:        time.Duration(config.App.ShutdownTimeoutSecs) * time.Second,
	}
}

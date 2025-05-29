//go:build wireinject
// +build wireinject

package main

import (
	"log/slog"

	"github.com/google/wire"

	"app/internal/adapters/rest/handlers"
	"app/internal/adapters/rest/routes"
	"app/internal/drivers/restserver"
)

func SetupApp() (*App, error) {
	wire.Build(
		slog.Default,

		handlers.NewHandlerFactory,
		wire.Bind(new(routes.HandlerFactory), new(*handlers.HandlerFactory)),

		routes.NewRouter,
		restserver.NewServerFromConfig,

		NewApp,
	)

	return &App{}, nil
}

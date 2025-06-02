//go:build wireinject
// +build wireinject

package main

import (
	"log/slog"

	"github.com/google/wire"

	"app/internal/adapters/rest/handlers"
	"app/internal/adapters/rest/routes"
	"app/internal/drivers/otel"
	"app/internal/drivers/restserver"
	"app/internal/drivers/wireproviders"
)

func SetupApp() (*App, error) {
	wire.Build(
		LoadConfig,
		wire.FieldsOf(new(*Config), "Server", "Tracing"),

		slog.Default,
		otel.SetupOtel,

		handlers.NewHandlerFactory,
		wire.Bind(new(routes.HandlerFactory), new(*handlers.HandlerFactory)),

		wireproviders.ProvideMiddlewares,
		routes.NewRouter,
		restserver.NewServerFromConfig,

		NewApp,
	)

	return &App{}, nil
}

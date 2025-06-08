//go:build wireinject
// +build wireinject

package main

import (
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

		wireproviders.ProvideOtelInstrumentedSlogLogger,
		otel.SetupOtel,

		handlers.NewHandlerFactory,
		wire.Bind(new(routes.HandlerFactory), new(*handlers.HandlerFactory)),

		wireproviders.ProvideMiddlewares,
		routes.NewRouter,
		restserver.NewServerFromConfig,

		NewAppFromConfig,
	)

	return &App{}, nil
}

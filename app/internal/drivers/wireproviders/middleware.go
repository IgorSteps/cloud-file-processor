package wireproviders

import (
	"app/internal/adapters/rest/middleware"
	"net/http"
)

func ProvideMiddlewares(config middleware.TracingMiddlewareConfig) []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		middleware.NewTracingMiddlewareFromConfig(config),
	}
}

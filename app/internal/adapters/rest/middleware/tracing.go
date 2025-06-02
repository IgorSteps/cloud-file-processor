package middleware

import (
	"net/http"

	"github.com/riandyrn/otelchi"
)

type TracingMiddlewareConfig struct {
	serviceName string `yaml:"serviceName"`
}

func TracingMiddleware(serviceName string) func(http.Handler) http.Handler {
	return otelchi.Middleware(serviceName)
}

func NewTracingMiddlewareFromConfig(config TracingMiddlewareConfig) func(http.Handler) http.Handler {
	return TracingMiddleware(config.serviceName)
}

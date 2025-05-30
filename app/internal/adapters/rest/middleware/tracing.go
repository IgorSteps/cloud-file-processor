package middleware

import (
	"net/http"

	"github.com/riandyrn/otelchi"
)

func TracingMiddleware() func(http.Handler) http.Handler {
	return otelchi.Middleware("hello")
}

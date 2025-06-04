package wireproviders

import (
	"log/slog"
	"os"

	"github.com/go-slog/otelslog"
)

func ProvideOtelInstrumentedSlogLogger() *slog.Logger {
	baseHandler := slog.NewJSONHandler(os.Stdout, nil)
	otelHandler := otelslog.NewHandler(baseHandler)
	return slog.New(otelHandler)
}

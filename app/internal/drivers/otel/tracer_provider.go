package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Type alias for tracer provider shutdown function to help readability.
type TracerProviderShutdown func(ctx context.Context) error

func SetupOtel() (TracerProviderShutdown, error) {
	// For now export traces to stdout, in the future switch to Jaeger.
	traceExporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter),
	)

	// Register trace provider.
	otel.SetTracerProvider(tracerProvider)

	// Register W3C Trace Context propogator.
	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil
}

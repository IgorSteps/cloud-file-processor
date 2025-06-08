package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	EXIT_FAILURE = 1
	EXIT_SUCCESS = 0
)

func main() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	app, err := SetupApp()
	if err != nil {
		log.Printf("failed to initialize the app: %v", err)
		return EXIT_FAILURE
	}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app.logger.InfoContext(ctx, "starting cloud-file-rest")

	// Run the REST server.
	go func() {
		if err := app.server.Run(); err != http.ErrServerClosed {
			app.logger.ErrorContext(ctx, "failed to start REST server", "error", err)
			stop()
		}
	}()

	<-ctx.Done()

	// Start a fresh context derived from the parent context (ignore cancellation) as the old one will have already
	// been canceled to avoid immediate exit.
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), app.shutdownTimeout)
	defer cancel()

	app.logger.InfoContext(ctx, "started graceful shutdown of cloud-file-rest")

	err = app.server.Shutdown(ctx)
	if err != nil {
		app.logger.ErrorContext(ctx, "failed to shutdown REST server", "error", err)
	}

	err = app.tracerProviderShutdown(ctx)
	if err != nil {
		app.logger.ErrorContext(ctx, "failed to shutdown tracer provider", "error", err)
	}

	app.logger.InfoContext(ctx, "completed graceful shut down of cloud-file-rest")

	return EXIT_SUCCESS
}

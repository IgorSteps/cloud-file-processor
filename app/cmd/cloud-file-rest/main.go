package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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
		log.Printf("failed to initlise the app: %v", err)
		return EXIT_FAILURE
	}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app.logger.InfoContext(ctx, "starting cloud-file-rest")

	// Start the REST server
	go func() {
		if err := app.server.Run(); err != nil {
			app.logger.ErrorContext(ctx, "failed to start REST server", "error", err)
			// return EXIT_FAILURE ?
		}
	}()

	<-ctx.Done()

	// Start a fresh context derived from the parent context (ignore cancellation) as the old one will have already
	// been canceled to avoid immediate exit.
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 30*time.Second)
	defer cancel()

	app.logger.InfoContext(ctx, "started graceful shutdown of cloud-file-rest")

	// TODO: graceful shutdown of the REST server here

	err = app.tracerProviderShutdown(ctx)
	if err != nil {
		app.logger.ErrorContext(ctx, "failed to shutdown tracer provider", "error", err)
	}

	app.logger.InfoContext(ctx, "completed graceful shut down of cloud-file-rest")

	return EXIT_SUCCESS
}

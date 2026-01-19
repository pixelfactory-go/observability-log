// Package main demonstrates the usage of the observability log package with various features.
package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/observability/log"
	"go.pixelfactory.io/pkg/observability/log/fields"
)

func main() {
	// Basic Usage (defaults to Info level)
	logger := log.New()
	logger.Debug("Debug Msg")
	logger.Info("Info Msg")
	logger.Warn("Warn Msg")
	logger.Error("Error Msg", fields.Error(errors.New("an error happened")))

	// Read DSN from the environment.
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		logger.Error("Failed to get a Sentry client", fields.Error(errors.New("SENTRY_DSN is not set")))
	}

	// Instantiate a client.
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		AttachStacktrace: true,
		Environment:      "production",
	})
	if err != nil {
		logger.Error("Failed to get a Sentry client", fields.Error(err))
	}

	// Using Sentry Core
	logger = log.New(
		log.WithLevel("debug"),
		log.WithSentry(sentry.CurrentHub().Client()),
		log.WithZapOption(zap.Development()),
	)

	// Add Fields
	// Fields must be added after logger creation
	logger = logger.With(fields.Service("myapp", "v1.0"))

	client := http.Client{Timeout: 1 * time.Second}
	request, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"https://httpbin.org/delay/2",
		http.NoBody,
	)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		logger.Error("An error occurred while creating request", fields.Error(err))
	}

	response, err := client.Do(request)
	if err != nil {
		logger.Fatal("An error occurred while sending request", fields.Error(err))
	}

	defer func() { _ = response.Body.Close() }()

	logger.Info("Sent Http Request", fields.HTTPRequest(request))
	logger.Info("Got Http Response", fields.HTTPResponse(response))
}

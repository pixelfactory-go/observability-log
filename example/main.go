package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/observability/log"
)

func main() {

	// Basic Usage (defaults to Info level)
	logger := log.New()
	logger.Debug("Debug Msg")
	logger.Info("Info Msg")
	logger.Warn("Warn Msg")
	logger.Error("Error Msg", log.Error(fmt.Errorf("An error happened")))

	// Read DSN from the environment.
	dsn := os.Getenv("SENTRY_DSN")
	if dsn == "" {
		logger.Error("Failed to get a Sentry client", log.Error(errors.New("SENTRY_DSN is not set")))
	}

	// Instantiate a client.
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		AttachStacktrace: true,
		Environment:      "production",
	})
	if err != nil {
		logger.Error("Failed to get a Sentry client", log.Error(err))
	}

	// Using Sentry Core
	logger = log.New(
		log.WithLevel("debug"),
		log.WithSentry(sentry.CurrentHub().Client()),
		log.WithZapOption(zap.Development()),
	)

	// Add Fields
	// Fields must be added after logger creation
	logger = logger.With(log.Service("myapp", "v1.0"))
	defer logger.Sync()

	client := http.Client{Timeout: 1 * time.Second}

	request, err := http.NewRequest("GET", "https://httpbin.org/delay/2", nil)
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		logger.Error("An error happened while creating request", log.Error(err))
	}

	response, err := client.Do(request)
	if err != nil {
		logger.Fatal("An error happened while sending request", log.Error(err))
	}

	logger.Info("Sent Http Request", log.HTTPRequest(request))
	logger.Info("Got Http Response", log.HTTPResponse(response))
}

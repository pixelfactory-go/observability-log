# Log

## Installation

```bash
go get go.pixelfactory.io/pkg/observability/log
```

## Usage

### Basic usage :

```go
package main

import (
	"errors"
	"fmt"
	"os"

	"go.pixelfactory.io/pkg/observability/log"
)

func main() {

	logger := log.New()
	logger.Debug("Debug Msg")
	logger.Info("Info Msg")
	logger.Warn("Warn Msg")
	logger.Error("Error Msg", log.Error(fmt.Errorf("An error happened")))
}
```

### Capture error in Sentry

Sentry is supported by implementing the `zapcore.Core` interface.

```go
package main

import (
	"errors"
	"fmt"
	"os"

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

	err = returnErr()
	if err != nil {
		logger.Error("An error happened", log.Error(err))
	}
}

func returnErr() error {
	return errors.New("Invalid ID '1234' not found")
}
```

## Elastic Common Schema

The Logger is using ECS version 1.5.0 <https://www.elastic.co/guide/en/ecs/current/index.html>.

Usage :

```go
package main

import (
	"go.pixelfactory.io/pkg/observability/log"
)

func main() {

	// New Logger with ECS Service field
	logger := log.New().With(log.Service("myapp", "v1.0.0"))

	uaString := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11"

	// Log UserAgent
	logger.Info("Info Msg", log.UserAgent(uaString))
}
```

The produced log entry follows ECS specification and is understood by Elasticsearch :

```json
{
  "log.level": "info",
  "@timestamp": "2020-07-31T12:51:38.313+0200",
  "log.origin": {
    "file.name": "observability/titi.go",
    "file.line": 15
  },
  "message": "Info Msg",
  "service": {
    "name": "myapp",
    "version": "v1.0.0"
  },
  "user_agent": {
    "original": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11",
    "name": "Chrome",
    "version": "23.0.1271.97"
  },
  "ecs.version": "1.5.0"
}
```

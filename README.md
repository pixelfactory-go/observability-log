# Observability Log

[![Go Reference](https://pkg.go.dev/badge/go.pixelfactory.io/pkg/observability/log.svg)](https://pkg.go.dev/go.pixelfactory.io/pkg/observability/log)
[![Go Report Card](https://goreportcard.com/badge/go.pixelfactory.io/pkg/observability/log)](https://goreportcard.com/report/go.pixelfactory.io/pkg/observability/log)
[![codecov](https://codecov.io/gh/pixelfactory-go/observability-log/branch/main/graph/badge.svg)](https://codecov.io/gh/pixelfactory-go/observability-log)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A structured logging library for Go applications that combines the power of [Zap](https://github.com/uber-go/zap) with [Elastic Common Schema (ECS)](https://www.elastic.co/guide/en/ecs/current/index.html) formatting and optional [Sentry](https://sentry.io) integration.

## Features

- **High Performance**: Built on top of uber-go/zap for blazing-fast structured logging
- **ECS Compliant**: Outputs logs in Elastic Common Schema v1.5.0 format for seamless Elasticsearch integration
- **Sentry Integration**: Optional error tracking with Sentry for error-level logs
- **Structured Fields**: Rich set of pre-built field helpers for HTTP requests/responses, user agents, services, and more
- **Context-Aware**: Easily add contextual information to logs using child loggers
- **Production Ready**: Battle-tested in production environments

## Installation

```bash
go get go.pixelfactory.io/pkg/observability/log
```

## Usage

### Basic Usage

```go
package main

import (
	"fmt"

	"go.pixelfactory.io/pkg/observability/log"
	"go.pixelfactory.io/pkg/observability/log/fields"
)

func main() {
	logger := log.New()
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message", fields.Error(fmt.Errorf("something went wrong")))
}
```

### Customized Logger

```go
package main

import (
	"go.uber.org/zap"
	"go.pixelfactory.io/pkg/observability/log"
	"go.pixelfactory.io/pkg/observability/log/fields"
)

func main() {
	// Create logger with custom options
	logger := log.New(
		log.WithLevel("debug"),
		log.WithZapOption(zap.Development()),
	)

	// Add service context
	logger = logger.With(fields.Service("myapp", "v1.0.0"))
	defer logger.Sync()

	logger.Info("Application started")
}
```

### Sentry Integration

```go
package main

import (
	"os"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.pixelfactory.io/pkg/observability/log"
	"go.pixelfactory.io/pkg/observability/log/fields"
)

func main() {
	// Initialize Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		AttachStacktrace: true,
		Environment:      "production",
	})
	if err != nil {
		panic(err)
	}

	// Create logger with Sentry integration
	logger := log.New(
		log.WithLevel("info"),
		log.WithSentry(sentry.CurrentHub().Client()),
	)

	// Add service context
	logger = logger.With(fields.Service("myapp", "v1.0.0"))
	defer logger.Sync()

	// Error-level logs are automatically sent to Sentry
	logger.Error("Database connection failed", fields.Error(err))
}
```

### Advanced Field Usage

#### HTTP Request Logging

```go
import (
	"net/http"

	"go.pixelfactory.io/pkg/observability/log"
	"go.pixelfactory.io/pkg/observability/log/fields"
)

func handler(w http.ResponseWriter, r *http.Request) {
	logger := log.New()

	// Log HTTP request with ECS fields
	logger.Info("Incoming request",
		fields.HTTPRequest(r),
		fields.UserAgent(r.Header.Get("User-Agent")),
		fields.URL(r.URL),
	)

	// Process request...

	// Log HTTP response
	logger.Info("Request completed",
		fields.HTTPResponse(200, len(responseBody)),
	)
}
```

#### User Agent Parsing

```go
import (
	"go.pixelfactory.io/pkg/observability/log"
	"go.pixelfactory.io/pkg/observability/log/fields"
)

func main() {
	logger := log.New().With(fields.Service("myapp", "v1.0.0"))

	uaString := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11"

	// Automatically parses user agent into structured fields
	logger.Info("User activity", fields.UserAgent(uaString))
}
```

Output (ECS format):
```json
{
  "log.level": "info",
  "@timestamp": "2020-07-31T12:51:38.313+0200",
  "log.origin": {
    "file.name": "main.go",
    "file.line": 15
  },
  "message": "User activity",
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

## Configuration Options

| Option | Description | Default |
|--------|-------------|---------|
| `WithLevel(level string)` | Set log level (debug, info, warn, error, fatal, panic) | `info` |
| `WithSentry(client *sentry.Client)` | Enable Sentry integration for error-level logs | Disabled |
| `WithZapOption(opts ...zap.Option)` | Add custom Zap options | None |

## Available Field Helpers

The `fields` package provides ECS-compliant field helpers:

- `fields.Error(err error)` - Error information
- `fields.Service(name, version string)` - Service identification
- `fields.HTTPRequest(r *http.Request)` - HTTP request details
- `fields.HTTPResponse(statusCode int, bodyBytes int)` - HTTP response details
- `fields.UserAgent(ua string)` - Parsed user agent information
- `fields.URL(url *url.URL)` - URL components
- `fields.Source(ip, port string)` - Source IP and port

## Elastic Common Schema

This logger outputs logs following [Elastic Common Schema (ECS) v1.5.0](https://www.elastic.co/guide/en/ecs/current/index.html), ensuring compatibility with Elasticsearch and Kibana for log aggregation and analysis.

## Development

### Prerequisites

- Go 1.24 or higher
- golangci-lint v2.8.0 or higher

### Available Commands

```bash
# Format code
make fmt

# Run linter
make lint

# Run tests with coverage
make test

# Build the project
make build

# Show all available commands
make help
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

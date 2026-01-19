// Package log provides a structured logging interface with ECS formatting and optional Sentry integration.
package log

import (
	"os"

	"github.com/getsentry/sentry-go"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	zapsentry "go.pixelfactory.io/pkg/observability/log/sentry"
)

// Logger is a simplified abstraction of the zap.Logger.
type Logger interface {
	Debug(msg string, fields ...zapcore.Field)
	Info(msg string, fields ...zapcore.Field)
	Warn(msg string, fields ...zapcore.Field)
	Error(msg string, fields ...zapcore.Field)
	Fatal(msg string, fields ...zapcore.Field)
	Panic(msg string, fields ...zapcore.Field)
}

// DefaultLogger delegates all calls to the underlying zap.Logger.
type DefaultLogger struct {
	level  *zap.AtomicLevel
	logger *zap.Logger
}

// Option type.
type Option func(*DefaultLogger)

// WithLevel logger level option.
func WithLevel(level string) Option {
	return func(l *DefaultLogger) {
		l.level.SetLevel(GetZapLogLevel(level))
	}
}

// WithSentry enables sentry.
func WithSentry(client *sentry.Client) Option {
	return func(l *DefaultLogger) {
		// Get Sentry zap Core that handle only Error level
		sentryCore := zapsentry.NewCore(zapcore.ErrorLevel, client)
		// NewTee creates a Core that duplicates log entries into two or more underlying Cores.
		core := zapcore.NewTee(l.logger.Core(), sentryCore)
		// Create new zap Logger
		l.logger = newZapLogger(core)
	}
}

// WithZapOption add fields.Service.
func WithZapOption(opts ...zap.Option) Option {
	return func(l *DefaultLogger) {
		l.logger = l.logger.WithOptions(opts...)
	}
}

// New returns a new logger with default values.
func New(opts ...Option) *DefaultLogger {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	atomicLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	core := ecszap.NewCore(encoderConfig, os.Stdout, atomicLevel)

	l := &DefaultLogger{
		logger: newZapLogger(core),
		level:  &atomicLevel,
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

// newZapLogger returns zap.Logger from zap Core.
func newZapLogger(core zapcore.Core) *zap.Logger {
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// Debug logs a debug msg with fields.
func (l *DefaultLogger) Debug(msg string, fields ...zapcore.Field) {
	l.logger.Debug(msg, fields...)
}

// Info logs an info msg with fields.
func (l *DefaultLogger) Info(msg string, fields ...zapcore.Field) {
	l.logger.Info(msg, fields...)
}

// Warn logs an warning msg with fields.
func (l *DefaultLogger) Warn(msg string, fields ...zapcore.Field) {
	l.logger.Warn(msg, fields...)
}

// Error logs an error msg with fields.
func (l *DefaultLogger) Error(msg string, fields ...zapcore.Field) {
	l.logger.Error(msg, fields...)
}

// Fatal logs a fatal error msg with fields and panics. Apps will have to recover if ever needed.
func (l *DefaultLogger) Fatal(msg string, fields ...zapcore.Field) {
	// Calls panic, as zap.Fatal calls os.Exit and isn't recoverable.
	l.Panic(msg, fields...)
}

// Panic logs a fatal error msg and fields and panics. Apps will have to recover if ever needed.
func (l *DefaultLogger) Panic(msg string, fields ...zapcore.Field) {
	l.logger.Panic(msg, fields...)
}

// With creates a child logger, and optionally adds some context fields to that logger.
func (l *DefaultLogger) With(fields ...zapcore.Field) *DefaultLogger {
	clone := l.clone()
	clone.logger = l.logger.With(fields...)
	return clone
}

// Sync call zap.Logger Sync() method.
func (l *DefaultLogger) Sync() error {
	return l.logger.Sync()
}

func (l *DefaultLogger) clone() *DefaultLogger {
	clone := *l
	return &clone
}

// GetZapLogLevel returns zap.AtomicLevel from string.
func GetZapLogLevel(logLevel string) zapcore.Level {
	level := zapcore.InfoLevel
	switch logLevel {
	case "info":
		level = zapcore.InfoLevel
	case "debug":
		level = zapcore.DebugLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	case "panic":
		level = zapcore.PanicLevel
	}

	return level
}

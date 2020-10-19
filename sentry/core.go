package zapsentry

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"

	ecsfields "go.pixelfactory.io/pkg/observability/log/fields"
)

// zapcore.Level to sentry.Level map
var zapLevelToSentrySeverity = map[zapcore.Level]sentry.Level{
	zapcore.DebugLevel:  sentry.LevelDebug,
	zapcore.InfoLevel:   sentry.LevelInfo,
	zapcore.WarnLevel:   sentry.LevelWarning,
	zapcore.ErrorLevel:  sentry.LevelError,
	zapcore.DPanicLevel: sentry.LevelFatal,
	zapcore.PanicLevel:  sentry.LevelFatal,
	zapcore.FatalLevel:  sentry.LevelFatal,
}

// errorKey is zap.Field key for ecs.Error
// https://github.com/elastic/ecs-logging-go-zap/blob/master/internal/error.go
const errorKey = "error"

// serviceKey is zap.Field key for fields.Service
// https://github.com/elastic/ecs-logging-go-zap/blob/master/internal/error.go
const serviceKey = "service"

// Option type
type Option func(*Core)

// DefaultSentryFlushTimeout is sentry flush timeout used in
// sentry.Flush() when calling Core.Sync()
const DefaultSentryFlushTimeout = 5 * time.Second

// SetFlushTimeout set sentry flush timeout
func SetFlushTimeout(timeout time.Duration) Option {
	return func(core *Core) {
		core.sentryFlushTimeout = timeout
	}
}

// Core struct
type Core struct {
	zapcore.LevelEnabler
	client             *sentry.Client
	sentryFlushTimeout time.Duration
	fields             []zapcore.Field
}

// NewCore creates a zapcore.Core.
func NewCore(enab zapcore.LevelEnabler, client *sentry.Client, options ...Option) *Core {
	core := &Core{
		LevelEnabler:       enab,
		client:             client,
		sentryFlushTimeout: DefaultSentryFlushTimeout,
	}

	for _, opt := range options {
		opt(core)
	}

	return core
}

// With adds structured context to the Core.
func (c *Core) With(fields []zapcore.Field) zapcore.Core {
	// Clone core.
	clone := *c

	// Clone and append fields.
	clone.fields = make([]zapcore.Field, len(c.fields)+len(fields))
	copy(clone.fields, c.fields)
	copy(clone.fields[len(c.fields):], fields)

	// Done.
	return &clone
}

// Check verifies whether or not the provided entry should be logged,
// by comparing the log level with the configured log level in the core.
// If it should be logged the core is added to the returned entry.
func (c *Core) Check(entry zapcore.Entry, checked *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(entry.Level) {
		return checked.AddCore(entry, c)
	}
	return checked
}

// Write converts entry to Sentry event and send it
func (c *Core) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	// Create a Sentry Event.
	event := sentry.NewEvent()
	event.Message = entry.Message

	// Process entry.
	event.Level = zapLevelToSentrySeverity[entry.Level]
	event.Timestamp = entry.Time
	event.Logger = entry.LoggerName

	// Process fields.
	encoder := zapcore.NewMapObjectEncoder()

	// When set, relevant Sentry interfaces are added.
	var err error
	var svc *ecsfields.ServiceField

	// processField processes the given field.
	// When false is returned, the whole entry is to be skipped.
	processField := func(field zapcore.Field) bool {
		// Look for "service" key.
		switch field.Key {
		case serviceKey:
			if s, ok := field.Interface.(*ecsfields.ServiceField); ok {
				svc = s
			} else {
				field.AddTo(encoder)
			}

		// Look for "error" key.
		case errorKey:
			if ex, ok := field.Interface.(error); ok {
				err = ex
			} else {
				field.AddTo(encoder)
			}

		default:
			// Add to the encoder in case this is not a significant key.
			field.AddTo(encoder)
		}

		return true
	}

	// Process core fields first.
	for _, field := range c.fields {
		if !processField(field) {
			return nil
		}
	}

	// Process the fields passed directly.
	// These can be then used to overwrite the core fields.
	for _, field := range fields {
		if !processField(field) {
			return nil
		}
	}

	// Process error
	if err != nil {
		// In case an error object is present, create an exception.
		// Capture the stack trace in any case.
		stacktrace := sentry.ExtractStacktrace(err)
		if stacktrace == nil {
			stacktrace = sentry.NewStacktrace()
			// stacktrace.Frames = filterFrames(stacktrace.Frames)
		}
		// Handle wrapped errors for github.com/pingcap/errors and github.com/pkg/errors
		cause := errors.Cause(err)
		event.Exception = []sentry.Exception{{
			Value:      cause.Error(),
			Type:       reflect.TypeOf(cause).String(),
			Stacktrace: stacktrace,
		}}
	} else {
		stacktrace := sentry.NewStacktrace()
		stacktrace.Frames = filterFrames(stacktrace.Frames)
		event.Exception = []sentry.Exception{{
			Value:      entry.Message,
			Stacktrace: stacktrace,
		}}
	}

	// fields into tags.
	tags := make(map[string]string)

	// Process service
	if svc != nil {
		tags["service.name"] = svc.Name
		tags["service.version"] = svc.Version
	}

	for key, value := range encoder.Fields {
		if v, ok := value.(string); ok {
			tags[key] = v
		} else {
			tags[key] = fmt.Sprintf("%v", value)
		}
	}

	// Add tags and extra into the packet.
	if len(tags) != 0 {
		event.Tags = tags
	}

	hub := sentry.CurrentHub()
	// Capture the packet.
	_ = c.client.CaptureEvent(event, nil, hub.Scope())
	return nil
}

func filterFrames(frames []sentry.Frame) []sentry.Frame {
	if len(frames) == 0 {
		return nil
	}

	filteredFrames := make([]sentry.Frame, 0, len(frames))

	for _, frame := range frames {
		// Skip Zap internal code in the frames.
		if strings.HasPrefix(frame.Function, "go.uber.org/zap") {
			continue
		}
		// Skip zapsentry code in the frames.
		if strings.HasPrefix(frame.Module, "go.pixelfactory.io/pkg/observability/log") &&
			!strings.HasSuffix(frame.Module, "_test") {
			continue
		}

		filteredFrames = append(filteredFrames, frame)
	}

	return filteredFrames
}

// Sync flushes buffered logs (if any).
func (c *Core) Sync() error {
	c.client.Flush(c.sentryFlushTimeout)
	return nil
}

package zapsentry_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	ecsfields "go.pixelfactory.io/pkg/observability/log/fields"
	zapsentry "go.pixelfactory.io/pkg/observability/log/sentry"
)

// FuzzCoreWrite fuzzes the Core.Write method with arbitrary log entries.
func FuzzCoreWrite(f *testing.F) {
	// Add seed corpus
	f.Add("test message", int8(zapcore.ErrorLevel), time.Now().Unix())
	f.Add("", int8(zapcore.InfoLevel), int64(0))
	f.Add("error occurred", int8(zapcore.FatalLevel), time.Now().Unix())
	f.Add(strings.Repeat("A", 1000), int8(zapcore.WarnLevel), time.Now().Unix())
	f.Add("unicode: 你好世界 🌍", int8(zapcore.ErrorLevel), time.Now().Unix())

	f.Fuzz(func(t *testing.T, message string, level int8, timestamp int64) {
		// Create a mock Sentry client
		client, err := sentry.NewClient(sentry.ClientOptions{
			Dsn: "", // Empty DSN for testing
		})
		if err != nil {
			t.Skip("Failed to create Sentry client")
		}

		// Map int8 to valid zapcore.Level
		zapLevel := zapcore.Level(level % 7)

		// Create core
		core := zapsentry.NewCore(zapLevel, client)

		// Create entry
		entry := zapcore.Entry{
			Level:   zapLevel,
			Time:    time.Unix(timestamp, 0),
			Message: message,
		}

		// Should not panic when writing
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Core.Write panicked: %v", r)
			}
		}()

		err = core.Write(entry, nil)
		if err != nil {
			t.Errorf("Core.Write returned error: %v", err)
		}
	})
}

// FuzzCoreWriteWithFields fuzzes the Core.Write method with various fields.
func FuzzCoreWriteWithFields(f *testing.F) {
	// Add seed corpus
	f.Add("message", "key1", "value1")
	f.Add("", "", "")
	f.Add("error message", "service", "my-service")
	f.Add("test", "特殊键", "特殊值")

	f.Fuzz(func(t *testing.T, message, key, value string) {
		// Create a mock Sentry client
		client, err := sentry.NewClient(sentry.ClientOptions{
			Dsn: "", // Empty DSN for testing
		})
		if err != nil {
			t.Skip("Failed to create Sentry client")
		}

		// Create core
		core := zapsentry.NewCore(zapcore.ErrorLevel, client)

		// Create entry
		entry := zapcore.Entry{
			Level:   zapcore.ErrorLevel,
			Time:    time.Now(),
			Message: message,
		}

		// Create field
		fields := []zapcore.Field{
			zap.String(key, value),
		}

		// Should not panic when writing with fields
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Core.Write with fields panicked: %v", r)
			}
		}()

		err = core.Write(entry, fields)
		if err != nil {
			t.Errorf("Core.Write returned error: %v", err)
		}
	})
}

// FuzzCoreWriteWithError fuzzes the Core.Write method with error fields.
func FuzzCoreWriteWithError(f *testing.F) {
	// Add seed corpus
	f.Add("error occurred", "something went wrong")
	f.Add("", "")
	f.Add("panic", strings.Repeat("X", 500))
	f.Add("failure", "unicode error: 错误")

	f.Fuzz(func(t *testing.T, message, errorMsg string) {
		// Create a mock Sentry client
		client, err := sentry.NewClient(sentry.ClientOptions{
			Dsn: "", // Empty DSN for testing
		})
		if err != nil {
			t.Skip("Failed to create Sentry client")
		}

		// Create core
		core := zapsentry.NewCore(zapcore.ErrorLevel, client)

		// Create entry
		entry := zapcore.Entry{
			Level:   zapcore.ErrorLevel,
			Time:    time.Now(),
			Message: message,
		}

		// Create error field
		fields := []zapcore.Field{
			zap.Error(errors.New(errorMsg)),
		}

		// Should not panic when writing with error
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Core.Write with error panicked: %v", r)
			}
		}()

		err = core.Write(entry, fields)
		if err != nil {
			t.Errorf("Core.Write returned error: %v", err)
		}
	})
}

// FuzzCoreWriteWithService fuzzes the Core.Write method with service fields.
func FuzzCoreWriteWithService(f *testing.F) {
	// Add seed corpus
	f.Add("message", "service-name", "1.0.0")
	f.Add("", "", "")
	f.Add("test", "my-service", "v2.3.4-beta")
	f.Add("log", "特殊服务", "版本1.0")

	f.Fuzz(func(t *testing.T, message, serviceName, serviceVersion string) {
		// Create a mock Sentry client
		client, err := sentry.NewClient(sentry.ClientOptions{
			Dsn: "", // Empty DSN for testing
		})
		if err != nil {
			t.Skip("Failed to create Sentry client")
		}

		// Create core
		core := zapsentry.NewCore(zapcore.ErrorLevel, client)

		// Create entry
		entry := zapcore.Entry{
			Level:   zapcore.ErrorLevel,
			Time:    time.Now(),
			Message: message,
		}

		// Create service field
		fields := []zapcore.Field{
			ecsfields.Service(serviceName, serviceVersion),
		}

		// Should not panic when writing with service field
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Core.Write with service field panicked: %v", r)
			}
		}()

		err = core.Write(entry, fields)
		if err != nil {
			t.Errorf("Core.Write returned error: %v", err)
		}
	})
}

// FuzzCoreWith fuzzes the Core.With method.
func FuzzCoreWith(f *testing.F) {
	// Add seed corpus
	f.Add("key", "value")
	f.Add("", "")
	f.Add("特殊键", "特殊值")

	f.Fuzz(func(t *testing.T, key, value string) {
		// Create a mock Sentry client
		client, err := sentry.NewClient(sentry.ClientOptions{
			Dsn: "", // Empty DSN for testing
		})
		if err != nil {
			t.Skip("Failed to create Sentry client")
		}

		// Create core
		core := zapsentry.NewCore(zapcore.ErrorLevel, client)

		// Should not panic when creating child core with fields
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Core.With panicked: %v", r)
			}
		}()

		fields := []zapcore.Field{
			zap.String(key, value),
		}

		childCore := core.With(fields)
		if childCore == nil {
			t.Error("Core.With returned nil")
		}
	})
}

// FuzzCoreCheck fuzzes the Core.Check method.
func FuzzCoreCheck(f *testing.F) {
	// Add seed corpus
	f.Add("message", int8(zapcore.ErrorLevel), int8(zapcore.ErrorLevel))
	f.Add("", int8(zapcore.InfoLevel), int8(zapcore.ErrorLevel))
	f.Add("test", int8(zapcore.DebugLevel), int8(zapcore.WarnLevel))

	f.Fuzz(func(t *testing.T, message string, entryLevel, coreLevel int8) {
		// Create a mock Sentry client
		client, err := sentry.NewClient(sentry.ClientOptions{
			Dsn: "", // Empty DSN for testing
		})
		if err != nil {
			t.Skip("Failed to create Sentry client")
		}

		// Map int8 to valid zapcore.Level
		zapEntryLevel := zapcore.Level(entryLevel % 7)
		zapCoreLevel := zapcore.Level(coreLevel % 7)

		// Create core
		core := zapsentry.NewCore(zapCoreLevel, client)

		// Create entry
		entry := zapcore.Entry{
			Level:   zapEntryLevel,
			Message: message,
		}

		// Should not panic when checking
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Core.Check panicked: %v", r)
			}
		}()

		checked := &zapcore.CheckedEntry{}
		result := core.Check(entry, checked)

		// Result should never be nil
		if result == nil {
			t.Error("Core.Check returned nil")
		}
	})
}

// FuzzSetFlushTimeout fuzzes the SetFlushTimeout option.
func FuzzSetFlushTimeout(f *testing.F) {
	// Add seed corpus
	f.Add(int64(1000000000))  // 1 second
	f.Add(int64(0))           // 0 nanoseconds
	f.Add(int64(-1000000000)) // negative duration
	f.Add(int64(10000000000)) // 10 seconds

	f.Fuzz(func(t *testing.T, durationNanos int64) {
		// Create a mock Sentry client
		client, err := sentry.NewClient(sentry.ClientOptions{
			Dsn: "", // Empty DSN for testing
		})
		if err != nil {
			t.Skip("Failed to create Sentry client")
		}

		// Should not panic when creating core with any flush timeout
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("NewCore with SetFlushTimeout panicked: %v", r)
			}
		}()

		timeout := time.Duration(durationNanos)
		core := zapsentry.NewCore(zapcore.ErrorLevel, client, zapsentry.SetFlushTimeout(timeout))

		if core == nil {
			t.Error("NewCore returned nil")
		}

		// Test sync
		err = core.Sync()
		if err != nil {
			t.Errorf("Core.Sync returned error: %v", err)
		}
	})
}

// FuzzNewCore fuzzes the NewCore function.
func FuzzNewCore(f *testing.F) {
	// Add seed corpus
	f.Add(int8(zapcore.ErrorLevel))
	f.Add(int8(zapcore.InfoLevel))
	f.Add(int8(zapcore.DebugLevel))

	f.Fuzz(func(t *testing.T, level int8) {
		// Create a mock Sentry client
		client, err := sentry.NewClient(sentry.ClientOptions{
			Dsn: "", // Empty DSN for testing
		})
		if err != nil {
			t.Skip("Failed to create Sentry client")
		}

		// Map int8 to valid zapcore.Level
		zapLevel := zapcore.Level(level % 7)

		// Should not panic when creating core
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("NewCore panicked: %v", r)
			}
		}()

		core := zapsentry.NewCore(zapLevel, client)

		if core == nil {
			t.Error("NewCore returned nil")
		}

		// Verify core can be used
		entry := zapcore.Entry{
			Level:   zapLevel,
			Time:    time.Now(),
			Message: "test",
		}

		err = core.Write(entry, nil)
		if err != nil {
			t.Errorf("Core.Write returned error: %v", err)
		}
	})
}

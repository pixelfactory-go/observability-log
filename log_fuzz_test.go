package log_test

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"go.pixelfactory.io/pkg/observability/log"
)

// FuzzLogMessages fuzzes the logger with arbitrary message strings.
func FuzzLogMessages(f *testing.F) {
	// Add seed corpus
	f.Add("test message", uint8(0))
	f.Add("", uint8(1))
	f.Add("message with special chars: !@#$%^&*()", uint8(2))
	f.Add("unicode: 你好世界 🌍", uint8(3))
	f.Add("very long message "+string(make([]byte, 1000)), uint8(4))
	f.Add("\n\r\t", uint8(5))

	f.Fuzz(func(_ *testing.T, msg string, level uint8) {
		// Setup logger
		atomicLevel := zap.NewAtomicLevelAt(zapcore.DebugLevel)
		core, _ := observer.New(atomicLevel)
		obsOpts := zap.WrapCore(func(_ zapcore.Core) zapcore.Core {
			return core
		})
		logger := log.New(log.WithLevel("debug"), log.WithZapOption(obsOpts))

		// Test different log levels based on fuzzer input
		// Use modulo to map the uint8 to valid log levels
		switch level % 5 {
		case 0:
			logger.Debug(msg)
		case 1:
			logger.Info(msg)
		case 2:
			logger.Warn(msg)
		case 3:
			logger.Error(msg)
		case 4:
			// Panic needs to be recovered
			defer func() {
				_ = recover()
			}()
			logger.Panic(msg)
		}

		// Ensure sync doesn't crash
		_ = logger.Sync()
	})
}

// FuzzLogLevel fuzzes the GetZapLogLevel function with arbitrary level strings.
func FuzzLogLevel(f *testing.F) {
	// Add seed corpus with valid and invalid log levels
	f.Add("debug")
	f.Add("info")
	f.Add("warn")
	f.Add("error")
	f.Add("fatal")
	f.Add("panic")
	f.Add("")
	f.Add("invalid")
	f.Add("DEBUG")
	f.Add("InFo")
	f.Add("warning")
	f.Add("trace")

	f.Fuzz(func(t *testing.T, level string) {
		// Should not panic with any input
		result := log.GetZapLogLevel(level)

		// Result should always be a valid zapcore.Level
		if result < zapcore.DebugLevel || result > zapcore.FatalLevel {
			t.Errorf("GetZapLogLevel returned invalid level: %v", result)
		}
	})
}

// FuzzLoggerWithFields fuzzes the logger with arbitrary field keys and values.
func FuzzLoggerWithFields(f *testing.F) {
	// Add seed corpus
	f.Add("key1", "value1", "msg")
	f.Add("", "", "")
	f.Add("special!@#$", "value", "message")
	f.Add("unicode_key_你好", "unicode_value_世界", "test")

	f.Fuzz(func(_ *testing.T, key, value, msg string) {
		// Setup logger
		atomicLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
		core, _ := observer.New(atomicLevel)
		obsOpts := zap.WrapCore(func(_ zapcore.Core) zapcore.Core {
			return core
		})
		logger := log.New(log.WithLevel("info"), log.WithZapOption(obsOpts))

		// Create child logger with field
		childLogger := logger.With(zap.String(key, value))

		// Should not crash with any input
		childLogger.Info(msg)

		_ = childLogger.Sync()
	})
}

// FuzzNewLogger fuzzes the New function with different options.
func FuzzNewLogger(f *testing.F) {
	// Add seed corpus
	f.Add("debug")
	f.Add("info")
	f.Add("warn")
	f.Add("error")
	f.Add("fatal")
	f.Add("panic")
	f.Add("invalid")
	f.Add("")

	f.Fuzz(func(t *testing.T, level string) {
		// Should not panic when creating logger with any level string
		logger := log.New(log.WithLevel(level))
		if logger == nil {
			t.Error("Expected non-nil logger")
		}

		// Verify logger can be used
		logger.Info("test")
		_ = logger.Sync()
	})
}

// FuzzLoggerMultipleFields fuzzes logging with multiple fields.
func FuzzLoggerMultipleFields(f *testing.F) {
	// Add seed corpus
	f.Add("msg", "k1", "v1", "k2", "v2", int64(42), true)
	f.Add("", "", "", "", "", int64(0), false)
	f.Add("unicode", "键", "值", "key", "value", int64(-1), true)

	f.Fuzz(func(_ *testing.T, msg, k1, v1, k2, v2 string, num int64, flag bool) {
		// Setup logger
		atomicLevel := zap.NewAtomicLevelAt(zapcore.InfoLevel)
		core, _ := observer.New(atomicLevel)
		obsOpts := zap.WrapCore(func(_ zapcore.Core) zapcore.Core {
			return core
		})
		logger := log.New(log.WithLevel("info"), log.WithZapOption(obsOpts))

		// Should not crash with any combination of fields
		logger.Info(msg,
			zap.String(k1, v1),
			zap.String(k2, v2),
			zap.Int64("number", num),
			zap.Bool("flag", flag),
		)

		_ = logger.Sync()
	})
}

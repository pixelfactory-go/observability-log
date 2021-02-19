package log_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"go.pixelfactory.io/pkg/observability/log"
)

var (
	message        = "test log message"
	zapStringField = zap.String("testKey", "testValue")
)

func setupObserver(level zap.AtomicLevel) (zap.Option, *observer.ObservedLogs) {
	core, logs := observer.New(level)
	opts := zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return core
	})
	return opts, logs
}

func setupLogger() (*log.DefaultLogger, *observer.ObservedLogs) {
	level := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	obsOpts, logs := setupObserver(level)
	return log.New(log.WithLevel("debug"), log.WithZapOption(obsOpts)), logs
}

func Test_GetZapLogLevel(t *testing.T) {
	t.Parallel()

	levels := []struct {
		level    string
		zapLevel zapcore.Level
	}{
		{
			level:    "info",
			zapLevel: zap.InfoLevel,
		},
		{
			level:    "debug",
			zapLevel: zap.DebugLevel,
		},
		{
			level:    "warn",
			zapLevel: zap.WarnLevel,
		},
		{
			level:    "error",
			zapLevel: zap.ErrorLevel,
		},
		{
			level:    "fatal",
			zapLevel: zap.FatalLevel,
		},
		{
			level:    "panic",
			zapLevel: zap.PanicLevel,
		},
	}

	for i := range levels {
		testCase := levels[i]
		t.Run(testCase.level, func(t *testing.T) {
			is := require.New(t)
			l := log.GetZapLogLevel(testCase.level)
			is.Equal(l, testCase.zapLevel)
		})
	}

}

func Test_New(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	logger, _ := setupLogger()
	defer logger.Sync()

	is.NotEmpty(logger)
	is.Implements((*log.Logger)(nil), logger)
}

func Test_With(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	logger, logs := setupLogger()
	defer logger.Sync()

	logger = logger.With(zapStringField)
	logger.Info("")

	if logs.Len() != 1 {
		t.Errorf("No logs")
	} else {
		entry := logs.All()[0]
		is.Equal(entry.Context, []zapcore.Field{
			zapStringField,
		})
	}
}

func Test_Debug(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	logger, logs := setupLogger()
	defer logger.Sync()

	logger.Debug(message)

	if logs.Len() != 1 {
		t.Errorf("No logs")
	} else {
		entry := logs.All()[0]
		is.Equal(entry.Level, zap.DebugLevel)
		is.Equal(message, entry.Message)
	}
}

func Test_Info(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	logger, logs := setupLogger()
	defer logger.Sync()

	logger.Info(message)

	if logs.Len() != 1 {
		t.Errorf("No logs")
	} else {
		entry := logs.All()[0]
		is.Equal(entry.Level, zap.InfoLevel)
		is.Equal(message, entry.Message)
	}
}

func Test_Warn(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	logger, logs := setupLogger()
	defer logger.Sync()

	logger.Warn(message)

	if logs.Len() != 1 {
		t.Errorf("No logs")
	} else {
		entry := logs.All()[0]
		is.Equal(entry.Level, zap.WarnLevel)
		is.Equal(message, entry.Message)
	}
}

func Test_Error(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	logger, logs := setupLogger()
	defer logger.Sync()

	logger.Error(message)

	if logs.Len() != 1 {
		t.Errorf("No logs")
	} else {
		entry := logs.All()[0]
		is.Equal(entry.Level, zap.ErrorLevel)
		is.Equal(message, entry.Message)
	}
}

func Test_Fatal(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	logger, _ := setupLogger()
	defer logger.Sync()

	is.Panics(func() {
		logger.Fatal("I should panic")
	})
}

func Test_Panic(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	logger, logs := setupLogger()
	defer logger.Sync()

	is.Panics(func() {
		logger.Panic(message)
	})

	if logs.Len() != 1 {
		t.Errorf("No logs")
	} else {
		entry := logs.All()[0]
		is.Equal(entry.Level, zap.PanicLevel)
		is.Equal(message, entry.Message)
	}
}

package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Object is a wrappers of zap.Object.
func Object(key string, val zapcore.ObjectMarshaler) zap.Field {
	return zap.Object(key, val)
}

// Bool is a wrappers of zap.Bool.
func Bool(key string, val bool) zap.Field {
	return zap.Bool(key, val)
}

// Bools is a wrappers of zap.Bools.
func Bools(key string, val []bool) zap.Field {
	return zap.Bools(key, val)
}

// String is a wrappers of zap.String.
func String(key string, val string) zap.Field {
	return zap.String(key, val)
}

// Strings is a wrappers of zap.Strings.
func Strings(key string, val []string) zap.Field {
	return zap.Strings(key, val)
}

// ByteString is a wrappers of zap.ByteString.
func ByteString(key string, val []byte) zap.Field {
	return zap.ByteString(key, val)
}

// ByteStrings is a wrappers of zap.ByteStrings.
func ByteStrings(key string, val [][]byte) zap.Field {
	return zap.ByteStrings(key, val)
}

// Binary is a wrappers of zap.Binary.
func Binary(key string, val []byte) zap.Field {
	return zap.Binary(key, val)
}

// Int is a wrappers of zap.Int.
func Int(key string, val int) zap.Field {
	return zap.Int(key, val)
}

// Int64 is a wrappers of zap.Int64.
func Int64(key string, val int64) zap.Field {
	return zap.Int64(key, val)
}

// Float32 is a wrappers of zap.Float32.
func Float32(key string, val float32) zap.Field {
	return zap.Float32(key, val)
}

// Float32s is a wrappers of zap.Float32s.
func Float32s(key string, val []float32) zap.Field {
	return zap.Float32s(key, val)
}

// Float64 is a wrappers of zap.Float64.
func Float64(key string, val float64) zap.Field {
	return zap.Float64(key, val)
}

// Float64s is a wrappers of zap.Float64s.
func Float64s(key string, val []float64) zap.Field {
	return zap.Float64s(key, val)
}

// Time is a wrappers of zap.Time.
func Time(key string, val time.Time) zap.Field {
	return zap.Time(key, val)
}

// Times is a wrappers of zap.Times.
func Times(key string, val []time.Time) zap.Field {
	return zap.Times(key, val)
}

// Duration is a wrappers of zap.Duration.
func Duration(key string, val time.Duration) zap.Field {
	return zap.Duration(key, val)
}

// Durations is a wrappers of zap.Durations.
func Durations(key string, val []time.Duration) zap.Field {
	return zap.Durations(key, val)
}

// Any is a wrappers of zap.Any.
func Any(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}

// Error is a wrappers of zap.Error.
// If given error is not empty and captureOpts is defined, the error will be captured in sentry.
// The first argument of captureOpts is to send the error to sentry.
// The second one is to wait an acknowledgement from sentry.
func Error(err error) zap.Field {
	return zap.NamedError("error", err)
}

// NamedError is a wrappers of zap.NamedError.
// If given error is not empty and captureOpts is defined, the error will be captured in sentry.
// The first argument of captureOpts is to send the error to sentry.
// The second one is to wait an acknowledgement from sentry.
func NamedError(key string, err error, captureOpts ...bool) zap.Field {
	return zap.NamedError(key, err)
}

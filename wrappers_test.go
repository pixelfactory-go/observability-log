package log_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/observability/log"
	"go.pixelfactory.io/pkg/observability/log/fields"
)

func Test_Wrapper_Object(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Object("key", &fields.Service{Name: "testSvc", Version: "1.0.0"})
	is.NotEmpty(b)
	is.Equal(b, zap.Object("key", &fields.Service{Name: "testSvc", Version: "1.0.0"}))
}

func Test_Wrapper_Bool(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Bool("key", true)
	is.NotEmpty(b)
	is.Equal(b, zap.Bool("key", true))
}

func Test_Wrapper_Bools(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Bools("key", []bool{true})
	is.NotEmpty(b)
	is.Equal(b, zap.Bools("key", []bool{true}))
}

func Test_Wrapper_String(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.String("key", "string")
	is.NotEmpty(b)
	is.Equal(b, zap.String("key", "string"))
}

func Test_Wrapper_Strings(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Strings("key", []string{"string"})
	is.NotEmpty(b)
	is.Equal(b, zap.Strings("key", []string{"string"}))
}

func Test_Wrapper_ByteString(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	s := []byte("value")
	b := log.ByteString("key", s)
	is.NotEmpty(b)
	is.Equal(b, zap.ByteString("key", s))
}

func Test_Wrapper_ByteStrings(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	s := make([][]byte, 1)
	b := log.ByteStrings("key", s)
	is.NotEmpty(b)
	is.Equal(b, zap.ByteStrings("key", s))
}

func Test_Wrapper_Binary(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	s := make([]byte, 1)
	b := log.Binary("key", s)
	is.NotEmpty(b)
	is.Equal(b, zap.Binary("key", s))
}

func Test_Wrapper_Int(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Int("key", 1)
	is.NotEmpty(b)
	is.Equal(b, zap.Int("key", 1))
}

func Test_Wrapper_Int64(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Int64("key", 1)
	is.NotEmpty(b)
	is.Equal(b, zap.Int64("key", 1))
}

func Test_Wrapper_Float32(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Float32("key", 1.1)
	is.NotEmpty(b)
	is.Equal(b, zap.Float32("key", 1.1))
}

func Test_Wrapper_Float64(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Float64("key", 1.1)
	is.NotEmpty(b)
	is.Equal(b, zap.Float64("key", 1.1))
}

func Test_Wrapper_Float32s(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Float32s("key", []float32{1.1})
	is.NotEmpty(b)
	is.Equal(b, zap.Float32s("key", []float32{1.1}))
}

func Test_Wrapper_Float64s(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Float64s("key", []float64{1.1})
	is.NotEmpty(b)
	is.Equal(b, zap.Float64s("key", []float64{1.1}))
}

func Test_Wrapper_Time(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	now := time.Now()
	b := log.Time("key", now)
	is.NotEmpty(b)
	is.Equal(b, zap.Time("key", now))
}

func Test_Wrapper_Times(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	now := time.Now()
	b := log.Times("key", []time.Time{now, now})
	is.NotEmpty(b)
	is.Equal(b, zap.Times("key", []time.Time{now, now}))
}

func Test_Wrapper_Duration(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Duration("key", 1*time.Second)
	is.NotEmpty(b)
	is.Equal(b, zap.Duration("key", 1*time.Second))
}

func Test_Wrapper_Durations(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	b := log.Durations("key", []time.Duration{1 * time.Second})
	is.NotEmpty(b)
	is.Equal(b, zap.Durations("key", []time.Duration{1 * time.Second}))
}

func Test_Wrapper_Any(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	any := struct {
		test string
	}{test: "value"}
	b := log.Any("key", any)
	is.NotEmpty(b)
	is.Equal(b, zap.Any("key", any))
}

func Test_Wrapper_Error(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	err := fmt.Errorf("test error")
	b := log.Error(err)
	is.NotEmpty(b)
	is.Equal(b, zap.Error(err))
}

func Test_Wrapper_NamedError(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	err := fmt.Errorf("test error")
	b := log.NamedError("key", err)
	is.NotEmpty(b)
	is.Equal(b, zap.NamedError("key", err))
}

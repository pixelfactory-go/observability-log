package zapsentry

import (
	"testing"

	"github.com/getsentry/sentry-go"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func TestSentryCore_NewCore(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	sentryCore := NewCore(zapcore.InfoLevel, &sentry.Client{})
	is.NotEmpty(sentryCore)
	is.Implements((*zapcore.Core)(nil), sentryCore)
}

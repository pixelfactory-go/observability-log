package zapsentry_test

import (
	"testing"

	"github.com/getsentry/sentry-go"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	zapsentry "go.pixelfactory.io/pkg/observability/log/sentry"
)

func TestSentryCore_NewCore(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	sentryCore := zapsentry.NewCore(zapcore.InfoLevel, &sentry.Client{})
	is.NotEmpty(sentryCore)
	is.Implements((*zapcore.Core)(nil), sentryCore)
}

package log_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/observability/log"
	"go.pixelfactory.io/pkg/observability/log/fields"
)

func Test_UserAgent(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	uaString := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11"
	ua := log.UserAgent(uaString)
	is.NotEmpty(ua)
	is.Equal(ua, zap.Object("user_agent", &fields.UserAgent{Original: uaString}))
}

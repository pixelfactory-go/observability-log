package log_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/observability/log"
	"go.pixelfactory.io/pkg/observability/log/fields"
)

func Test_Service(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	service := log.Service("testSvc", "0.0.1")
	is.NotEmpty(service)
	is.Equal(service, zap.Object("service", &fields.Service{Name: "testSvc", Version: "0.0.1"}))
}

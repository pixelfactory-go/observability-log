package fields_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/observability/log/fields"
)

func Test_Service(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	service := fields.Service("testSvc", "0.0.1")
	is.NotEmpty(service)
	is.Equal(service, zap.Object("service", &fields.ServiceField{Name: "testSvc", Version: "0.0.1"}))
}

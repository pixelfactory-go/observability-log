package fields_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/observability/log/fields"
)

func Test_Source(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	source := fields.Source("10.0.0.1", 8080)
	is.NotEmpty(source)
	is.Equal(source, zap.Object("source", &fields.SourceField{IP: "10.0.0.1", Port: 8080}))
}

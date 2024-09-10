package fields_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/observability/log/fields"
)

func Test_URL(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	u, err := url.Parse("http://test/search?q=dotnet")
	is.NoError(err)

	urlField := fields.URL(u)
	is.NotEmpty(urlField)
	is.Equal(urlField, zap.Object("url", &fields.URLField{URL: u}))
}

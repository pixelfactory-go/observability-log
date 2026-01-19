package fields_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.pixelfactory.io/pkg/observability/log/fields"
)

func Test_HTTPRequest(t *testing.T) {
	t.Parallel()
	is := require.New(t)
	req := httptest.NewRequest(http.MethodGet, "http://test/foo", http.NoBody)
	reqField := fields.HTTPRequest(req)
	is.NotEmpty(reqField)
	is.Equal(reqField, zap.Object("http.request", &fields.HTTPRequestField{Request: req}))
}

func Test_HTTPResponse(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	handler := func(w http.ResponseWriter, _ *http.Request) {
		_, err := io.WriteString(w, "<html><body>Hello Test!</body></html>")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	req := httptest.NewRequest(http.MethodGet, "http://test/foo", http.NoBody)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	defer func() { _ = resp.Body.Close() }()

	respField := fields.HTTPResponse(resp)
	is.NotEmpty(respField)
	is.Equal(respField, zap.Object("http.response", &fields.HTTPResponseField{Response: resp}))
}

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
	req := httptest.NewRequest("GET", "http://test/foo", nil)
	reqField := fields.HTTPRequest(req)
	is.NotEmpty(reqField)
	is.Equal(reqField, zap.Object("http.request", &fields.HTTPRequestField{Request: req}))
}

func Test_HTTPResponse(t *testing.T) {
	t.Parallel()
	is := require.New(t)

	handler := func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "<html><body>Hello Test!</body></html>")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	req := httptest.NewRequest("GET", "http://test/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	// body, _ := ioutil.ReadAll(resp.Body)

	respField := fields.HTTPResponse(resp)
	is.NotEmpty(respField)
	is.Equal(respField, zap.Object("http.response", &fields.HTTPResponseField{Response: resp}))
}

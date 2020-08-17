package fields

import (
	"net/http"

	"go.uber.org/zap/zapcore"
)

// HTTPResponse struct represents ECS http.response object
// https://www.elastic.co/guide/en/ecs/current/ecs-http.html
type HTTPResponse struct {
	Response *http.Response
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (r *HTTPResponse) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("status_code", r.Response.StatusCode)
	enc.AddInt64("bytes", r.Response.ContentLength)
	return nil
}

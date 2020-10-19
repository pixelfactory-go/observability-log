package fields

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// HTTPResponseField struct represents ECS http.response object
// https://www.elastic.co/guide/en/ecs/current/ecs-http.html
type HTTPResponseField struct {
	Response *http.Response
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (r *HTTPResponseField) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt("status_code", r.Response.StatusCode)
	enc.AddInt64("bytes", r.Response.ContentLength)
	return nil
}

// HTTPResponse returns ECS http.response as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-http.html
func HTTPResponse(resp *http.Response) zapcore.Field {
	return zap.Object(
		"http.response",
		&HTTPResponseField{Response: resp},
	)
}

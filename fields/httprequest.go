// Package fields provides ECS-compliant field helpers for structured logging.
package fields

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// HTTPRequestField struct represents ECS http.request object
// https://www.elastic.co/guide/en/ecs/current/ecs-http.html
type HTTPRequestField struct {
	Request *http.Request
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (r *HTTPRequestField) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if r.Request == nil {
		return nil
	}

	enc.AddString("method", r.Request.Method)
	enc.AddString("version", r.Request.Proto)
	enc.AddString("referrer", r.Request.Referer())
	enc.AddInt64("bytes", r.Request.ContentLength)

	if err := enc.AddReflected("headers", r.Request.Header); err != nil {
		return err
	}
	return nil
}

// HTTPRequest returns ECS http.request as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-http.html
func HTTPRequest(req *http.Request) zapcore.Field {
	return zap.Object(
		"http.request",
		&HTTPRequestField{Request: req},
	)
}

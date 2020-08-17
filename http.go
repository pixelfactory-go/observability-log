package log

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.pixelfactory.io/pkg/observability/log/fields"
)

// HTTPRequest returns ECS http.request as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-http.html
func HTTPRequest(req *http.Request) zapcore.Field {
	return zap.Object(
		"http.request",
		&fields.HTTPRequest{Request: req},
	)
}

// HTTPResponse returns ECS http.response as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-http.html
func HTTPResponse(resp *http.Response) zapcore.Field {
	return zap.Object(
		"http.response",
		&fields.HTTPResponse{Response: resp},
	)
}

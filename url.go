package log

import (
	"net/url"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.pixelfactory.io/pkg/observability/log/fields"
)

// URL returns ECS url as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-url.html
func URL(url *url.URL) zapcore.Field {
	return zap.Object("url", &fields.URL{URL: url})
}

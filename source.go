package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.pixelfactory.io/pkg/observability/log/fields"
)

// Source returns ECS source as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-source.html
func Source(ip string, port int) zapcore.Field {
	return zap.Object("source", &fields.Source{IP: ip, Port: port})
}

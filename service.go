package log

import (
	"go.uber.org/zap/zapcore"

	"go.pixelfactory.io/pkg/observability/log/fields"
)

// Service returns ECS service as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-service.html
func Service(name string, version string) zapcore.Field {
	return Object("service", &fields.Service{Name: name, Version: version})
}

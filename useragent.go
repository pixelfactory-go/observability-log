package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.pixelfactory.io/pkg/observability/log/fields"
)

// UserAgent returns ECS user_agent as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-user_agent.html
func UserAgent(original string) zapcore.Field {
	return zap.Object("user_agent", &fields.UserAgent{Original: original})
}

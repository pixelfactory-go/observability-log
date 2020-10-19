package fields

import (
	"github.com/mssola/user_agent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// UserAgentField struct represents ECS user_agent object
// https://www.elastic.co/guide/en/ecs/current/ecs-user_agent.html
type UserAgentField struct {
	Original string
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (u *UserAgentField) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	ua := user_agent.New(u.Original)
	enc.AddString("original", u.Original)
	name, version := ua.Browser()
	enc.AddString("name", name)
	enc.AddString("version", version)
	return nil
}

// UserAgent returns ECS user_agent as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-user_agent.html
func UserAgent(original string) zapcore.Field {
	return zap.Object("user_agent", &UserAgentField{Original: original})
}

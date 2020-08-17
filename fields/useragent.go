package fields

import (
	"github.com/mssola/user_agent"
	"go.uber.org/zap/zapcore"
)

// UserAgent struct represents ECS user_agent object
// https://www.elastic.co/guide/en/ecs/current/ecs-user_agent.html
type UserAgent struct {
	Original string
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (u *UserAgent) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	ua := user_agent.New(u.Original)
	enc.AddString("original", u.Original)
	name, version := ua.Browser()
	enc.AddString("name", name)
	enc.AddString("version", version)
	return nil
}

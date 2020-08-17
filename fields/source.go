package fields

import (
	"go.uber.org/zap/zapcore"
)

// Source struct represents ECS source object
// https://www.elastic.co/guide/en/ecs/current/ecs-source.html
type Source struct {
	IP   string
	Port int
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (s *Source) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("ip", s.IP)
	enc.AddInt("port", s.Port)
	return nil
}

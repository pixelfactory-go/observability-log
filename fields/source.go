package fields

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// SourceField struct represents ECS source object
// https://www.elastic.co/guide/en/ecs/current/ecs-source.html
type SourceField struct {
	IP   string
	Port int
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (s *SourceField) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("ip", s.IP)
	enc.AddInt("port", s.Port)
	return nil
}

// Source returns ECS source as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-source.html
func Source(ip string, port int) zapcore.Field {
	return zap.Object("source", &SourceField{IP: ip, Port: port})
}

package fields

import (
	"go.uber.org/zap/zapcore"
)

// Service struct represents ECS service object
// https://www.elastic.co/guide/en/ecs/current/ecs-service.html
type Service struct {
	Name    string
	Version string
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (s *Service) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", s.Name)
	enc.AddString("version", s.Version)
	return nil
}

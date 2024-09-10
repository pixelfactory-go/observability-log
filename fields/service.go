package fields

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ServiceField struct represents ECS service object
// https://www.elastic.co/guide/en/ecs/current/ecs-service.html
type ServiceField struct {
	Name    string
	Version string
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (s *ServiceField) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", s.Name)
	enc.AddString("version", s.Version)
	return nil
}

// Service returns ECS service as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-service.html
func Service(name, version string) zapcore.Field {
	return zap.Object("service", &ServiceField{Name: name, Version: version})
}

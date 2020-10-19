package fields

import (
	"net/url"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// URLField struct represents ECS url object
// https://www.elastic.co/guide/en/ecs/current/ecs-url.html
type URLField struct {
	URL *url.URL
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (u *URLField) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("path", u.URL.Path)
	enc.AddString("query", u.URL.RawQuery)
	return nil
}

// URL returns ECS url as zap.Field
// https://www.elastic.co/guide/en/ecs/current/ecs-url.html
func URL(url *url.URL) zapcore.Field {
	return zap.Object("url", &URLField{URL: url})
}

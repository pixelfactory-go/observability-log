package fields

import (
	"net/url"

	"go.uber.org/zap/zapcore"
)

// URL struct represents ECS url object
// https://www.elastic.co/guide/en/ecs/current/ecs-url.html
type URL struct {
	URL *url.URL
}

// MarshalLogObject implements zapcore ObjectMarshaler.
func (u *URL) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("path", u.URL.Path)
	enc.AddString("query", u.URL.RawQuery)
	return nil
}

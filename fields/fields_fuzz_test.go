package fields_test

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"go.pixelfactory.io/pkg/observability/log/fields"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

// FuzzUserAgent fuzzes the UserAgent field with arbitrary user agent strings.
func FuzzUserAgent(f *testing.F) {
	// Add seed corpus with various user agent strings
	f.Add("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	f.Add("curl/7.64.1")
	f.Add("")
	f.Add("invalid")
	f.Add("Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X)")
	f.Add("PostmanRuntime/7.26.8")
	f.Add(strings.Repeat("A", 1000))
	f.Add("特殊字符 🎉 测试")

	f.Fuzz(func(t *testing.T, userAgent string) {
		// Should not panic with any user agent string
		logger := zaptest.NewLogger(t)
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("UserAgent panicked with input %q: %v", userAgent, r)
			}
		}()

		// Create field and log it
		field := fields.UserAgent(userAgent)
		logger.Info("test", field)
	})
}

// FuzzService fuzzes the Service field with arbitrary service names and versions.
func FuzzService(f *testing.F) {
	// Add seed corpus
	f.Add("my-service", "1.0.0")
	f.Add("", "")
	f.Add("service-name", "v2.3.4-beta")
	f.Add("特殊服务", "版本1.0")
	f.Add("service!@#$%", "version!@#$%")
	f.Add(strings.Repeat("A", 500), strings.Repeat("B", 500))

	f.Fuzz(func(t *testing.T, name, version string) {
		// Should not panic with any input
		logger := zaptest.NewLogger(t)
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Service panicked with input name=%q version=%q: %v", name, version, r)
			}
		}()

		field := fields.Service(name, version)
		logger.Info("test", field)
	})
}

// FuzzSource fuzzes the Source field with arbitrary IPs and ports.
func FuzzSource(f *testing.F) {
	// Add seed corpus
	f.Add("192.168.1.1", 8080)
	f.Add("", 0)
	f.Add("::1", 443)
	f.Add("invalid-ip", -1)
	f.Add("256.256.256.256", 99999)
	f.Add("2001:0db8:85a3:0000:0000:8a2e:0370:7334", 65535)

	f.Fuzz(func(t *testing.T, ip string, port int) {
		// Should not panic with any input
		logger := zaptest.NewLogger(t)
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Source panicked with input ip=%q port=%d: %v", ip, port, r)
			}
		}()

		field := fields.Source(ip, port)
		logger.Info("test", field)
	})
}

// FuzzURL fuzzes the URL field with arbitrary URL strings.
func FuzzURL(f *testing.F) {
	// Add seed corpus
	f.Add("https://example.com/path?query=value")
	f.Add("http://localhost:8080")
	f.Add("/relative/path")
	f.Add("//example.com")
	f.Add("https://example.com/路径?参数=值")
	f.Add("invalid://url")
	f.Add("")
	f.Add("ftp://ftp.example.com/file.txt")

	f.Fuzz(func(t *testing.T, urlString string) {
		// Parse URL - may fail, but shouldn't panic
		parsedURL, err := url.Parse(urlString)
		if err != nil {
			// If URL parsing fails, skip this input
			return
		}

		// Should not panic with any valid parsed URL
		logger := zaptest.NewLogger(t)
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("URL panicked with input %q: %v", urlString, r)
			}
		}()

		field := fields.URL(parsedURL)
		logger.Info("test", field)
	})
}

// FuzzHTTPRequest fuzzes the HTTPRequest field with arbitrary HTTP request data.
func FuzzHTTPRequest(f *testing.F) {
	// Add seed corpus
	f.Add("GET", "/path", "HTTP/1.1", "https://example.com", int64(100))
	f.Add("POST", "/", "HTTP/2.0", "", int64(0))
	f.Add("", "", "", "", int64(-1))
	f.Add("DELETE", "/api/v1/resource", "HTTP/1.0", "http://localhost", int64(1024))
	f.Add("特殊方法", "/特殊路径", "HTTP/3", "https://例子.com", int64(999999))

	f.Fuzz(func(t *testing.T, method, path, proto, referer string, contentLength int64) {
		// Create HTTP request
		req, err := http.NewRequest(method, "http://example.com"+path, nil)
		if err != nil {
			// If request creation fails, skip this input
			return
		}

		// Set additional fields
		req.Proto = proto
		req.Header.Set("Referer", referer)
		req.ContentLength = contentLength

		// Should not panic with any valid HTTP request
		logger := zaptest.NewLogger(t)
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("HTTPRequest panicked: %v", r)
			}
		}()

		field := fields.HTTPRequest(req)
		logger.Info("test", field)
	})
}

// FuzzHTTPRequestNil fuzzes the HTTPRequest field with nil request.
func FuzzHTTPRequestNil(f *testing.F) {
	f.Add(true)

	f.Fuzz(func(t *testing.T, useNil bool) {
		if !useNil {
			return
		}

		// Should not panic with nil request
		logger := zaptest.NewLogger(t)
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("HTTPRequest panicked with nil: %v", r)
			}
		}()

		field := fields.HTTPRequest(nil)
		logger.Info("test", field)
	})
}

// FuzzUserAgentField fuzzes the UserAgentField MarshalLogObject method.
func FuzzUserAgentField(f *testing.F) {
	// Add seed corpus
	f.Add("Mozilla/5.0")
	f.Add("")
	f.Add(strings.Repeat("X", 2000))

	f.Fuzz(func(t *testing.T, original string) {
		field := &fields.UserAgentField{Original: original}

		// Create a mock encoder and ensure MarshalLogObject doesn't panic
		enc := zapcore.NewMapObjectEncoder()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("UserAgentField.MarshalLogObject panicked: %v", r)
			}
		}()

		err := field.MarshalLogObject(enc)
		if err != nil {
			t.Errorf("MarshalLogObject returned error: %v", err)
		}
	})
}

// FuzzServiceField fuzzes the ServiceField MarshalLogObject method.
func FuzzServiceField(f *testing.F) {
	f.Add("name", "version")
	f.Add("", "")

	f.Fuzz(func(t *testing.T, name, version string) {
		field := &fields.ServiceField{Name: name, Version: version}

		enc := zapcore.NewMapObjectEncoder()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("ServiceField.MarshalLogObject panicked: %v", r)
			}
		}()

		err := field.MarshalLogObject(enc)
		if err != nil {
			t.Errorf("MarshalLogObject returned error: %v", err)
		}
	})
}

// FuzzSourceField fuzzes the SourceField MarshalLogObject method.
func FuzzSourceField(f *testing.F) {
	f.Add("127.0.0.1", 8080)
	f.Add("", 0)

	f.Fuzz(func(t *testing.T, ip string, port int) {
		field := &fields.SourceField{IP: ip, Port: port}

		enc := zapcore.NewMapObjectEncoder()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("SourceField.MarshalLogObject panicked: %v", r)
			}
		}()

		err := field.MarshalLogObject(enc)
		if err != nil {
			t.Errorf("MarshalLogObject returned error: %v", err)
		}
	})
}

// FuzzURLField fuzzes the URLField MarshalLogObject method.
func FuzzURLField(f *testing.F) {
	f.Add("https://example.com/path?query=value")
	f.Add("")

	f.Fuzz(func(t *testing.T, urlString string) {
		parsedURL, err := url.Parse(urlString)
		if err != nil {
			return
		}

		field := &fields.URLField{URL: parsedURL}

		enc := zapcore.NewMapObjectEncoder()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("URLField.MarshalLogObject panicked: %v", r)
			}
		}()

		err = field.MarshalLogObject(enc)
		if err != nil {
			t.Errorf("MarshalLogObject returned error: %v", err)
		}
	})
}

// FuzzHTTPRequestField fuzzes the HTTPRequestField MarshalLogObject method.
func FuzzHTTPRequestField(f *testing.F) {
	f.Add("GET", "/path", int64(100))

	f.Fuzz(func(t *testing.T, method, path string, contentLength int64) {
		req, err := http.NewRequest(method, "http://example.com"+path, nil)
		if err != nil {
			return
		}
		req.ContentLength = contentLength

		field := &fields.HTTPRequestField{Request: req}

		enc := zapcore.NewMapObjectEncoder()
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("HTTPRequestField.MarshalLogObject panicked: %v", r)
			}
		}()

		err = field.MarshalLogObject(enc)
		if err != nil {
			// Some errors are acceptable (e.g., reflection errors)
			// but the function should not panic
			t.Logf("MarshalLogObject returned error (acceptable): %v", err)
		}
	})
}

// FuzzWrappers fuzzes various zap field wrapper functions.
func FuzzWrappers(f *testing.F) {
	f.Add("key", "value", int64(42), true, float64(3.14))

	f.Fuzz(func(t *testing.T, key, strVal string, intVal int64, boolVal bool, floatVal float64) {
		logger := zaptest.NewLogger(t)
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Field wrapper panicked: %v", r)
			}
		}()

		// Test various field types
		logger.Info("test",
			zap.String(key, strVal),
			zap.Int64(key+"_int", intVal),
			zap.Bool(key+"_bool", boolVal),
			zap.Float64(key+"_float", floatVal),
		)
	})
}

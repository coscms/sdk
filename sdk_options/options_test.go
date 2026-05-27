package sdk_options

import (
	"net/url"
	"testing"
)

type testAppInfo struct {
	AppID       string
	AppSecret   string
	ApiEndpoint string
}

func (t *testAppInfo) GetAppID() string       { return t.AppID }
func (t *testAppInfo) GetAppSecret() string   { return t.AppSecret }
func (t *testAppInfo) GetApiEndpoint() string { return t.ApiEndpoint }

func TestNewOptions(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(TypeOAuth, app)

	if opts.GetAppID() != "test-id" {
		t.Errorf("expected test-id, got %s", opts.GetAppID())
	}
	if opts.GetAppSecret() != "test-secret" {
		t.Errorf("expected test-secret, got %s", opts.GetAppSecret())
	}
	if opts.GetApiEndpoint() != "https://api.example.com" {
		t.Errorf("expected https://api.example.com, got %s", opts.GetApiEndpoint())
	}
	if opts.Type != TypeOAuth {
		t.Errorf("expected TypeOauth, got %s", opts.Type)
	}
}

func TestNewOptionsWithOptions(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com"}
	opts := New(TypePayment, app, OptionUserAgent("test-agent"), OptionClientIP("127.0.0.1"))

	if opts.UserAgent != "test-agent" {
		t.Errorf("expected test-agent, got %s", opts.UserAgent)
	}
	if opts.ClientIP != "127.0.0.1" {
		t.Errorf("expected 127.0.0.1, got %s", opts.ClientIP)
	}
}

func TestOptionsGetAppIDNoGetter(t *testing.T) {
	opts := &Options{}
	if opts.GetAppID() != "" {
		t.Errorf("expected empty, got %s", opts.GetAppID())
	}
	if opts.GetAppSecret() != "" {
		t.Errorf("expected empty, got %s", opts.GetAppSecret())
	}
	if opts.GetApiEndpoint() != "" {
		t.Errorf("expected empty, got %s", opts.GetApiEndpoint())
	}
}

func TestSetAppInfoGetter(t *testing.T) {
	opts := &Options{}
	app := &testAppInfo{AppID: "new-id", AppSecret: "new-secret", ApiEndpoint: "https://api.example.com"}
	opts.SetAppInfoGetter(app)

	if opts.GetAppID() != "new-id" {
		t.Errorf("expected new-id, got %s", opts.GetAppID())
	}
}

func TestGetApiEndpointTrimSlash(t *testing.T) {
	app := &testAppInfo{AppID: "test-id", AppSecret: "test-secret", ApiEndpoint: "https://api.example.com/"}
	opts := New(TypeOAuth, app)

	if opts.GetApiEndpoint() != "https://api.example.com" {
		t.Errorf("expected https://api.example.com, got %s", opts.GetApiEndpoint())
	}
}

func TestResponseIsSuccess(t *testing.T) {
	r := Response{Code: 1}
	if !r.IsSuccess() {
		t.Error("expected IsSuccess()=true for Code=1")
	}

	r = Response{Code: 0}
	if r.IsSuccess() {
		t.Error("expected IsSuccess()=false for Code=0")
	}
}

func TestGetValueByKey(t *testing.T) {
	m := map[string]any{"key1": "val1", "key2": "val2"}

	if v := GetValueByKey(m, "key1"); v != "val1" {
		t.Errorf("expected val1, got %v", v)
	}

	if v := GetValueByKey(m, "nonexistent"); v != nil {
		t.Errorf("expected nil, got %v", v)
	}

	if v := GetValueByKey(m, "nonexistent", "key2"); v != "val2" {
		t.Errorf("expected val2 (fallback), got %v", v)
	}

	if v := GetValueByKey(nil, "key1"); v != nil {
		t.Errorf("expected nil for nil map, got %v", v)
	}
}

func TestOptionsToURLWithGeneratorNil(t *testing.T) {
	opts := &Options{}
	_, _, err := opts.ToURLWithGenerator(nil, "/test")
	if err == nil {
		t.Error("expected error for nil generator")
	}
}

func TestSetSigner(t *testing.T) {
	opts := &Options{}
	opts.SetSigner(func(vals url.Values, secret string) string {
		return "sig"
	})
}

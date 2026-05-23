package sdk_options

import (
	"net/url"
	"testing"
)

func TestSignString(t *testing.T) {
	sig := SignString("raw", "key")
	if sig == "" {
		t.Error("SignString returned empty string")
	}
}

func TestCheckSign(t *testing.T) {
	sig := SignString("raw", "key")
	err := CheckSign("raw", sig, "key")
	if err != nil {
		t.Errorf("CheckSign should pass: %v", err)
	}

	err = CheckSign("raw", "wrong", "key")
	if err != ErrInvalidSign {
		t.Errorf("CheckSign should fail with ErrInvalidSign, got: %v", err)
	}
}

func TestBuildURLValues(t *testing.T) {
	v := url.Values{}
	v.Set("foo", "bar")
	result := BuildURLValues(v, "secret", nil)

	if result.Get("foo") != "bar" {
		t.Errorf("expected foo=bar, got foo=%s", result.Get("foo"))
	}
	if result.Get("timestamp") == "" {
		t.Error("timestamp should be set")
	}
	if result.Get("sign") == "" {
		t.Error("sign should be set")
	}
}

func TestBuildURLValuesWithCustomSigner(t *testing.T) {
	v := url.Values{}
	v.Set("foo", "bar")
	customSigner := func(vals url.Values, secret string) string {
		return "custom_" + secret
	}
	result := BuildURLValues(v, "mysecret", Signer(customSigner))

	if result.Get("sign") != "custom_mysecret" {
		t.Errorf("expected sign=custom_mysecret, got sign=%s", result.Get("sign"))
	}
}

func TestBuildURLValuesNil(t *testing.T) {
	result := BuildURLValues(nil, "secret", nil)
	if result == nil {
		t.Error("expected non-nil result")
	}
}

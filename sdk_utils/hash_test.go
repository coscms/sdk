package sdk_utils

import (
	"testing"
)

func TestMd5(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"md5", "1bc29b36f623ba82aaf6724fd3b16718"},
	}
	for _, tt := range tests {
		got := Md5(tt.input)
		if got != tt.want {
			t.Errorf("Md5(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

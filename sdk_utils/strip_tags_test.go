package sdk_utils

import "testing"

func TestStripTags(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"hello world", "hello world"},
		{"<b>bold</b>", "bold"},
		{"<a href=\"link\">click</a>", "click"},
		{"<script>alert('xss')</script>", ""},
		{"<style>body { color: red; }</style>text", "text"},
		{"<!-- comment -->visible", "visible"},
		{"<p>para</p>", "para"},
		{"< br />text", "< br />text"},
		{"escape < notag", "escape < notag"},
		{"<div><span>nested</span></div>", "nested"},
	}
	for _, tt := range tests {
		got := StripTags(tt.input)
		if got != tt.want {
			t.Errorf("StripTags(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

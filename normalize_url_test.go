package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove extra slashes",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove too many extra slashes",
			inputURL: "blog.boot.dev/path/////",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "bad protocol, still works",
			inputURL: "httttttps://blog.boot.dev/path/////",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove query parameters",
			inputURL: "https://blog.boot.dev/path/?somequery=true",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "works without host",
			inputURL: "/path/?somequery=true",
			expected: "/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

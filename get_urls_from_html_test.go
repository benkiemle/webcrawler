package main

import (
	"reflect"
	"testing"
)

func TestGetURLSFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "multiple attributes",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a target="_blank" href="/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "empty href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name:     "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a>
			<span>Boot.dev></span>
		</a>
	</body>
</html>
`,
			expected: make([]string, 0),
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLSFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

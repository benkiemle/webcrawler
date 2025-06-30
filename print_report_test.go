package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	tests := []struct {
		name       string
		inputPages map[string]int
		expected   []kv
	}{
		{
			name:       "simple sort",
			inputPages: map[string]int{"hi": 1, "yes": 2},
			expected:   []kv{{"yes", 2}, {"hi", 1}},
		},
		{
			name:       "matching counts, sort alphabetical",
			inputPages: map[string]int{"hi": 2, "area": 3, "test": 1, "hayes": 2},
			expected:   []kv{{"area", 3}, {"hayes", 2}, {"hi", 2}, {"test", 1}},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortPages(tc.inputPages)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

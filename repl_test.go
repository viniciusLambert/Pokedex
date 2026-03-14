package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"extra spaces": {
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		"only spaces": {
			input:    "        ",
			expected: []string{},
		},
		"no spaces": {
			input:    "nonenone",
			expected: []string{"nonenone"},
		},
		"lower cases": {
			input:    "AAADD ASAA SSSAA",
			expected: []string{"aaadd", "asaa", "sssaa"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(tc.expected, got)
			if diff != "" {
				t.Fatalf("%s", diff)
			}
		})
	}
}

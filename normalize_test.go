package goblok

import "testing"

func TestNormalize(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty_string",
			input:    "",
			expected: "",
		},
		{
			name:     "with_leetspeak",
			input:    "1n1 4d4l4h k0nt0l.",
			expected: "ini adalah kontol.",
		},
		{
			name:     "with_leetspeak_and_punctuation",
			input:    "1n1 4d4l4h ku(1n9.",
			expected: "ini adalah kucing.",
		},
		{
			name:     "no_leetspeak",
			input:    "Ini adalah kucing.",
			expected: "Ini adalah kucing.",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := normalize(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, actual)
			}
		})
	}
}

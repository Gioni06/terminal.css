package build

import "testing"

func TestGetMinifiedFileName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"lib/terminal.css", "terminal.min.css"},
		{"lib/terminal.min.css", "terminal.min.css"},
		{"lib/terminal", "terminal.min"},
		{"lib/app.v2.css", "app.v2.min.css"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := getMinifiedFileName(tt.input)
			if output != tt.expected {
				t.Errorf("got %s, want %s", output, tt.expected)
			}
		})
	}
}

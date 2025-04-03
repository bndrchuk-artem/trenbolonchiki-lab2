package lab2

import (
	"bytes"
	"strings"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		err      bool
	}{
		{"Valid input", "2 2 +", "4", false},
		{"Invalid input", "2 x", "", true},
		{"Negative numbers", "2 -2 +", "0", false},
		{"Complex expression", "2 3 * 4 +", "10", false},
		{"Leading/trailing spaces", "   2 2 +   ", "4", false},
		{"Empty input", "", "0", false}, 
		{"Division", "6 2 /", "3", false},
		{"Division by zero", "6 0 /", "", true},
	}	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			output := &bytes.Buffer{}
			handler := &ComputeHandler{Input: input, Output: output}

			err := handler.Compute()
			if tt.err {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if output.String() != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, output.String())
				}
			}
		})
	}
}
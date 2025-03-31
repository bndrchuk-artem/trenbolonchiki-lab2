package lab2

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		err      bool
		errType  error
	}{
		{"Valid input", "2 2 +", "4", false, nil},
		{"Invalid input", "2 x", "", true, errors.New("invalid character")},
		{"Negative numbers", "2 -2 +", "0", false, nil},
		{"Complex expression", "2 3 * 4 +", "10", false, nil},
		{"Leading/trailing spaces", "   2 2 +   ", "4", false, nil},
		{"Empty input", "", "0", false, nil}, 
		{"Division", "6 2 /", "3", false, nil},
		{"Division by zero", "6 0 /", "", true, errors.New("division by zero")},

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			output := &bytes.Buffer{}
			handler := &ComputeHandler{Input: input, Output: output}

			err := handler.Compute()

			if tt.err {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				} else if tt.errType != nil {
					if !errors.Is(err, tt.errType) && err.Error() != tt.errType.Error() {
						t.Errorf("Expected error of type %v, got %v", tt.errType, err)
					}
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if output.String() != tt.expected {
					t.Errorf("Expected %q, got %q", tt.expected, output.String())
				}
			}
		})
	}
}
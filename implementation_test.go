package lab2

import (
	"fmt"
	"testing"
)

func TestEvaluatePostfix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		err      bool
	}{
		{"Simple addition", "2 2 +", 5, false},//
		{"Simple subtraction", "5 3 -", 2, false},
		{"Simple multiplication", "3 4 *", 12, false},
		{"Complex expression", "4 2 - 3 * 5 +", 11, false},
		{"Division", "6 2 /", 3, false},
		{"Power", "2 3 ^", 8, false},
		{"Empty input", "", 0, true},
		{"Invalid operator", "2 2 x", 0, true},
		{"Too few operands", "2 +", 0, true},
		{"Division by zero", "2 0 /", 0, true},
		{"Too many operands", "2 3 4 +", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := EvaluatePostfix(tt.input)
			if tt.err {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if res != tt.expected {
					t.Errorf("expected %d, got %d", tt.expected, res)
				}
			}
		})
	}
}

func ExampleEvaluatePostfix() {
	res, _ := EvaluatePostfix("2 2 +")
	fmt.Println(res)
	// Output: 4
}

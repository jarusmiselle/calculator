package unaryoperation_test

import (
	"testing"

	"github.com/jarusmiselle/calculator/internal/expression/unaryoperation"
)

func TestFactorialOperation(t *testing.T) {
	tests := []struct {
		name     string
		operand  float64
		expected float64
		hasError bool
	}{
		{
			name:     "factorial of 5",
			operand:  5,
			expected: 120,
			hasError: false,
		},
		{
			name:     "factorial of 3",
			operand:  3,
			expected: 6,
			hasError: false,
		},
		{
			name:     "factorial of 0",
			operand:  0,
			expected: 1,
			hasError: false,
		},
		{
			name:     "factorial of negative number",
			operand:  -3,
			expected: 0,
			hasError: true,
		},
		{
			name:     "factorial of non-integer",
			operand:  4.5,
			expected: 0,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := unaryoperation.Factorial.Evaluate(tt.operand)
			if tt.hasError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("did not expect an error but got: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %f but got %f", tt.expected, result)
				}
			}
		})
	}
}

package calculator

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func TestFactorial(t *testing.T) {
	tt := []struct {
		input       int
		expected    int
		expectedErr error
	}{{
		input:       -2,
		expected:    0,
		expectedErr: errors.New("-2 should be zero or higher"),
	}, {
		input:    0,
		expected: 1,
	}, {
		input:    1,
		expected: 1,
	}, {
		input:    3,
		expected: 6,
	}, {
		input:    4,
		expected: 24,
	}, {
		input:    5,
		expected: 120,
	}}

	for _, c := range tt {
		name := fmt.Sprintf("input: %d, expected: %d, err: %v", c.input, c.expected, c.expectedErr)
		log.Println(name)
		t.Run(name, func(t *testing.T) {
			res, err := factorial(c.input)

			if res != c.expected {
				t.Fatal("response did not match expected")
			}

			if c.expectedErr != nil {
				if err == nil || c.expectedErr.Error() != err.Error() {
					t.Fatal("error did not match expectedErr")
				}
			} else {
				if err != nil {
					t.Fatal("error was not nil")
				}
			}
		})
	}
}

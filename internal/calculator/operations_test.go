package calculator

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	tt := []struct {
		input       []float64
		expected    float64
		expectedErr error
	}{{
		input:    []float64{5, 36},
		expected: 41,
	}, {
		input:    []float64{27, 42},
		expected: 69,
	}, {
		input:    []float64{47, -12},
		expected: 35,
	}, {
		input:       []float64{7, 5, 12},
		expected:    0,
		expectedErr: errors.New("usage: calc add left right"),
	}, {
		input:       []float64{42},
		expected:    0,
		expectedErr: errors.New("usage: calc add left right"),
	}, {
		input:       []float64{46, 29, 5, 7, 10},
		expected:    0,
		expectedErr: errors.New("usage: calc add left right"),
	}}

	for _, c := range tt {
		name := fmt.Sprintf("input: %f, expected: %f, err: %v", c.input, c.expected, c.expectedErr)
		log.Println(name)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res, err := add(c.input)

			if res != c.expected {
				t.Fatal("response did not match expected")
			}

			if c.expectedErr != nil {
				if err == nil || c.expectedErr.Error() != err.Error() {
					t.Fatalf("error %s did not match expectedErr %s", c.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Fatal("error was not nil")
				}
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	t.Parallel()

	tt := []struct {
		input       []float64
		expected    float64
		expectedErr error
	}{{
		input:       []float64{-2},
		expected:    0,
		expectedErr: errors.New("-2.000000 should be zero or higher"),
	}, {
		input:    []float64{0},
		expected: 1,
	}, {
		input:    []float64{1},
		expected: 1,
	}, {
		input:    []float64{3},
		expected: 6,
	}, {
		input:    []float64{4},
		expected: 24,
	}, {
		input:    []float64{5},
		expected: 120,
	}, {
		input:       []float64{4, 3},
		expected:    0,
		expectedErr: errors.New("usage: calc factorial arg"),
	}}

	for _, c := range tt {
		name := fmt.Sprintf("input: %f, expected: %f, err: %v", c.input, c.expected, c.expectedErr)
		log.Println(name)
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res, err := factorial(c.input)

			if res != c.expected {
				t.Fatal("response did not match expected")
			}

			if c.expectedErr != nil {
				if err == nil || c.expectedErr.Error() != err.Error() {
					t.Fatalf("error %s did not match expectedErr %s", c.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Fatal("error was not nil")
				}
			}
		})
	}
}

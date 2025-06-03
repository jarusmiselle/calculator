package calculator

import (
	"errors"
	"fmt"
)

type operation func(left, right float64) (float64, error)

var Operations = map[string]operation{
	"add":      add,
	"+":        add,
	"subtract": subtract,
	"sub":      subtract,
	"-":        subtract,
	"multiply": multiply,
	"mul":      multiply,
	"*":        multiply,
	"divide":   divide,
	"div":      divide,
	"/":        divide,
}

func add(left, right float64) (float64, error) {
	res := left + right
	return res, nil
}

func subtract(left, right float64) (float64, error) {
	res := left - right
	return res, nil
}

func multiply(left, right float64) (float64, error) {
	res := left * right
	return res, nil
}

func divide(left, right float64) (float64, error) {
	if right == 0 {
		return 0, errors.New("right cannot equal zero")
	}

	res := left / right
	return res, nil
}

func factorial(left float64) (float64, error) {
	if left < 0 {
		return 0, fmt.Errorf("%f should be zero or higher", left)
	}

	if left == 1 || left == 0 {
		return 1, nil
	}

	res, err := factorial(left - 1)
	return left * res, err
}

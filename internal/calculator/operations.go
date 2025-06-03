package calculator

import (
	"errors"
	"strconv"
)

type operation func(left, right int) (int, error)

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

func add(left int, right int) (int, error) {
	res := left + right
	return res, nil
}

func subtract(left int, right int) (int, error) {
	res := left - right
	return res, nil
}

func multiply(left int, right int) (int, error) {
	res := left * right
	return res, nil
}

func divide(left int, right int) (int, error) {
	if right == 0 {
		return 0, errors.New("right cannot equal zero")
	}

	res := left / right
	return res, nil
}

func factorial(left int) (int, error) {
	if left < 0 {
		return 0, errors.New(strconv.Itoa(left) + " should be zero or higher")
	}

	if left == 1 || left == 0 {
		return 1, nil
	}

	res, err := factorial(left - 1)
	return left * res, err
}

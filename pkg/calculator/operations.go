package calculator

import (
	"errors"
	"strconv"
)

func Add(left int, right int) int {
	res := left + right
	return res
}

func Subtract(left int, right int) int {
	res := left - right
	return res
}

func Multiply(left int, right int) int {
	res := left * right
	return res
}

func Divide(left int, right int) (int, error) {
	if right == 0 {
		return 0, errors.New("right cannot equal zero")
	}

	res := left / right
	return res, nil
}

func Factorial(left int) (int, error) {
	if left < 0 {
		return 0, errors.New(strconv.Itoa(left) + " should be zero or higher")
	}

	if left == 1 || left == 0 {
		return 1, nil
	}

	res, err := Factorial(left - 1)
	return left * res, err
}

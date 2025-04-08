package calculator

import "errors"

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

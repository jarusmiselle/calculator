package calculator

import (
	"errors"
	"strconv"
)

type Expression struct {
	Op    string
	Left  int
	Right int
}

func GetExpression(args []string) (Expression, error) {
	var exp Expression

	if len(args) != 3 {
		return exp, errors.New("usage: calc operation left right")
	}

	exp.Op = args[0]

	left, err := strconv.Atoi(args[1])
	if err != nil {
		return exp, err
	}
	exp.Left = left

	right, err := strconv.Atoi(args[2])
	if err != nil {
		return exp, err
	}
	exp.Right = right

	return exp, nil
}

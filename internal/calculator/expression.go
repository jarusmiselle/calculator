package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Expression struct {
	opFunk operation
	args   []float64
}

func GetExpression(args []string) (Expression, error) {
	var exp Expression

	if len(args) <= 1 {
		return exp, errors.New("usage: calc operation ..args")
	}

	opf, ok := operations[strings.ToLower(args[0])]
	if !ok {
		return exp, errors.New("invalid operation")
	}
	exp.opFunk = opf

	for _, arg := range args[1:] {
		f, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return exp, fmt.Errorf("parse args: %w", err)
		}
		exp.args = append(exp.args, f)
	}

	return exp, nil
}

func (e Expression) Evaluate() (float64, error) {
	return e.opFunk(e.args)
}

package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Expression struct {
	OpFunk operation
	Op     string
	Args   []float64
}

func GetExpression(args []string) (Expression, error) {
	var exp Expression

	if len(args) <= 1 {
		return exp, errors.New("usage: calc operation ..args")
	}

	exp.Op = strings.ToLower(args[0])

	opf, ok := Operations[exp.Op]
	if !ok {
		panic(errors.New("invalid operation"))
	}
	exp.OpFunk = opf

	for _, arg := range args[1:] {
		f, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return exp, fmt.Errorf("parse args: %w", err)
		}
		exp.Args = append(exp.Args, f)
	}

	return exp, nil
}

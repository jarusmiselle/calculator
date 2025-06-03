package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Expression struct {
	Op   string
	Args []float64
}

func GetExpression(args []string) (Expression, error) {
	var exp Expression

	if len(args) <= 1 {
		return exp, errors.New("usage: calc operation ..args")
	}

	exp.Op = strings.ToLower(args[0])

	for _, arg := range args[1:] {
		f, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return exp, fmt.Errorf("parse args: %w", err)
		}
		exp.Args = append(exp.Args, f)
	}

	return exp, nil
}

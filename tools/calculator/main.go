package main

import (
	"fmt"

	"github.com/jarusmiselle/calculator/internal/expression"
)

func main() {
	ue, err := expression.NewUnaryExpression(
		"!",
		expression.ValueExpression(3),
	)
	if err != nil {
		panic(err)
	}

	exp, err := expression.NewBinaryExpression(
		"+",
		ue,
		expression.ValueExpression(62),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(ue)
	fmt.Println(ue.Evaluate())

	fmt.Println(exp)
	fmt.Println(exp.Evaluate())
}

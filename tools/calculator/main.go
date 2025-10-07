package main

import (
	"fmt"

	"github.com/jarusmiselle/calculator/internal/expression"
)

func main() {
	s, t := ctof(30)

	x, y := ftoc(70)

	fmt.Println(s)
	fmt.Println(t)

	fmt.Println(x)
	fmt.Println(y)
}

func ctof(c float64) (string, float64) {
	sco, _ := expression.NewBinaryExpression(
		"/",
		expression.ValueExpression(9),
		expression.ValueExpression(5),
	)

	sct, _ := expression.NewBinaryExpression(
		"*",
		expression.ValueExpression(c),
		sco,
	)

	f, _ := expression.NewBinaryExpression(
		"+",
		sct,
		expression.ValueExpression(32),
	)

	res, _ := f.Evaluate()

	return f.String(), res
}

func ftoc(f float64) (string, float64) {
	sco, _ := expression.NewBinaryExpression(
		"/",
		expression.ValueExpression(5),
		expression.ValueExpression(9),
	)

	sct, _ := expression.NewBinaryExpression(
		"-",
		expression.ValueExpression(f),
		expression.ValueExpression(32),
	)

	c, _ := expression.NewBinaryExpression(
		"*",
		sct,
		sco,
	)

	res, _ := c.Evaluate()

	return c.String(), res
}

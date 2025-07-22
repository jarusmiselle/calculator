package main

import (
	"fmt"

	"github.com/rhafezi1/calculator/internal/expression"
)

func main() {
	ve, _ := expression.NewValueExpression(23)
	ue, err := expression.NewUnaryExpression("!", ve)
	if err != nil {
		panic(err)
	}

	ve, _ = expression.NewValueExpression(62)
	exp, err := expression.NewBinaryExpression("+", ue, ve)
	if err != nil {
		panic(err)
	}

	fmt.Println(ue)
	fmt.Println(ue.Evaluate())

	fmt.Println(exp)
	fmt.Println(exp.Evaluate())
}

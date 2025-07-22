package main

import (
	"fmt"

	et "github.com/rhafezi1/calculator/internal/expressiontree"
)

func main() {
	ue, err := et.NewUnaryExpression("!", et.ValueExpression{
		Value: 23,
	})
	if err != nil {
		panic(err)
	}

	exp := et.BinaryExpression{
		Left: ue,
		Right: et.ValueExpression{
			Value: 62,
		},
	}

	fmt.Println(ue)
	fmt.Println(ue.Evaluate())

	fmt.Println(exp)
	fmt.Println(exp.Evaluate())
}

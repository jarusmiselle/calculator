package main

import (
	"fmt"

	"github.com/rhafezi1/calculator/internal/expressiontree"
)

func main() {
	ue := expressiontree.UnaryExpression{
		Value: expressiontree.ValueExpression{
			Value: 23,
		},
	}

	exp := expressiontree.BinaryExpression{
		Left: ue,
		Right: expressiontree.ValueExpression{
			Value: 62,
		},
	}

	fmt.Println(ue)
	fmt.Println(ue.Evaluate())

	fmt.Println(exp)
	fmt.Println(exp.Evaluate())
}

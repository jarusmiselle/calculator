package main

import (
	"fmt"

	et "github.com/rhafezi1/calculator/internal/expressiontree"
)

func main() {
	ve, _ := et.NewValueExpression(23)
	ue, err := et.NewUnaryExpression("!", ve)
	if err != nil {
		panic(err)
	}

	ve, _ = et.NewValueExpression(62)
	exp := et.BinaryExpression{
		Left:  ue,
		Right: ve,
	}

	fmt.Println(ue)
	fmt.Println(ue.Evaluate())

	fmt.Println(exp)
	fmt.Println(exp.Evaluate())
}

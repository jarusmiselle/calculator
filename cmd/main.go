package main

import (
	"fmt"
	"os"

	"github.com/rhafezi1/calculator/internal/calculator"
)

func main() {
	exp, err := calculator.GetExpression(os.Args[1:])
	if err != nil {
		panic(err)
	}

	res, err := exp.OpFunk(exp.Args)
	if err != nil {
		panic(err)
	}

	fmt.Println(exp)
	fmt.Println(err)
	fmt.Println(res)
}

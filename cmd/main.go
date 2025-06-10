package main

import (
	"fmt"
	"os"

	"github.com/rhafezi1/calculator/internal/calculator"
)

func main() {
	exp, err := calculator.GetExpression(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := exp.Evaluate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(exp)
	fmt.Println(err)
	fmt.Println(res)
}

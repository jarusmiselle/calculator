package main

import (
	"fmt"
	"os"

	"github.com/rhafezi1/calculator/pkg/calculator"
)

func main() {
	exp, err := calculator.GetExpression(os.Args[1:])
	fmt.Println(exp)
	fmt.Println(err)
}

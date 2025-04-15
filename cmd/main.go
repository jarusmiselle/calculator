package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rhafezi1/calculator/pkg/calculator"
)

func main() {
	exp, err := calculator.GetExpression(os.Args[1:])
	if err != nil {
		panic(err)
	}

	lop := strings.ToLower(exp.Op)

	var res int
	switch lop {
	case "+":
		fallthrough
	case "add":
		res = calculator.Add(exp.Left, exp.Right)
	case "-":
		fallthrough
	case "subtract":
		fallthrough
	case "sub":
		res = calculator.Subtract(exp.Left, exp.Right)
	case "*":
		fallthrough
	case "multiply":
		fallthrough
	case "mult":
		res = calculator.Multiply(exp.Left, exp.Right)
	case "/":
		fallthrough
	case "divide":
		fallthrough
	case "div":
		res, err = calculator.Divide(exp.Left, exp.Right)
	default:
		// error
	}

	fmt.Println(exp)
	fmt.Println(err)
	fmt.Println(res)
}

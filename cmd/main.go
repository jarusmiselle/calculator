package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rhafezi1/calculator/internal/calculator"
)

func main() {
	exp, err := calculator.GetExpression(os.Args[1:])
	if err != nil {
		panic(err)
	}

	lop := strings.ToLower(exp.Op)

	var res int
	switch lop {
	case "+", "add":
		res = calculator.Add(exp.Left, exp.Right)
	case "-", "subtract", "sub":
		res = calculator.Subtract(exp.Left, exp.Right)
	case "*", "multiply", "mult":
		res = calculator.Multiply(exp.Left, exp.Right)
	case "/", "divide", "div":
		res, err = calculator.Divide(exp.Left, exp.Right)
	case "!", "factorial", "fact":
		res, err = calculator.Factorial(exp.Left)
	default:
		err = errors.New("operator incorrect")
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(exp)
	fmt.Println(err)
	fmt.Println(res)
}

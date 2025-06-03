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
	opf, ok := calculator.Operations[lop]
	if !ok {
		panic(errors.New("invalid operation"))
	}

	res, err := opf(exp.Left, exp.Right)
	if err != nil {
		panic(err)
	}

	fmt.Println(exp)
	fmt.Println(err)
	fmt.Println(res)
}

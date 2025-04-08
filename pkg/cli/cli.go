package cli

import (
	"fmt"
	"strings"
)

func DoOpSwitch(op string) {
	op = strings.ToLower(op)

	switch op {
	case "add":
		fmt.Println("add")
	case "subtract":
		fmt.Println("subtract")
	case "multiply":
		fmt.Println("multiply")
	case "divide":
		fmt.Println("divide")
	default:
		panic("Usage:operation must be one of: add, subtact, multiply, divide")
	}
}

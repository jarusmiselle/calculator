package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jarusmiselle/calculator/internal/expression"
	"github.com/jarusmiselle/calculator/internal/expression/unaryexpression"
	"github.com/jarusmiselle/calculator/internal/expression/unaryoperation"
)

func main() {
	// a()
	// b()
	// c()
	// d()
	// e()
	err := f()
	if err != nil {
		fmt.Println(err)
	}

	e := err
	for e = errors.Unwrap(e); e != nil; e = errors.Unwrap(e) {
		fmt.Println("caused by:", e)
	}

	e = errors.Unwrap(err)
	for e != nil {
		fmt.Println("caused by:", e)
		e = errors.Unwrap(e)
	}
}

func a() {
	num := 5
	bad := true
	blue := "red"

	out, _ := json.Marshal(num)
	fmt.Println(string(out))

	out, _ = json.Marshal(bad)
	fmt.Println(string(out))

	out, _ = json.Marshal(blue)
	fmt.Println(string(out))
}

func b() {
	num := "27"
	bad := "false"
	blue := "\"grey\""

	var numV int
	err := json.Unmarshal([]byte(num), &numV)
	if err != nil {
		panic(err)
	}
	fmt.Println(numV)

	var badV bool
	err = json.Unmarshal([]byte(bad), &badV)
	if err != nil {
		panic(err)
	}
	fmt.Println(badV)

	var blueV string
	err = json.Unmarshal([]byte(blue), &blueV)
	if err != nil {
		panic(err)
	}
	fmt.Println(blueV)
}

func c() {
	ex := must(expression.NewBinaryExpression(
		"+",
		expression.ValueExpression(4),
		must(unaryexpression.New(
			unaryoperation.Negation,
			expression.ValueExpression(3),
		)),
	))

	fmt.Println(ex)

	b := must(json.Marshal(ex))
	fmt.Println(string(b))
	fmt.Println(len(b))
}

func d() {
	v := expression.ValueExpression(4)
	u := must(unaryexpression.New(
		unaryoperation.Negation,
		expression.ValueExpression(3),
	))
	be := must(expression.NewBinaryExpression("+", v, u))

	b, err := json.MarshalIndent(be, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	fmt.Println(len(b))
}

func e() {
	data := []byte(`
{
	"type": "UnaryExpression",
	"operation": "-",
	"operand": {
		"type": "ValueExpression",
		"value": 3
	}
}`)

	var ue unaryexpression.UnaryExpression
	err := json.Unmarshal(data, &ue)
	if err != nil {
		fmt.Println(err)
	}
}

func f() error {
	data := []byte(`{"type": "ValueExpression","value": "3"}`)

	var ve expression.ValueExpression
	err := json.Unmarshal(data, &ve)
	if err != nil {
		return fmt.Errorf("f: %w", err)
	}
	fmt.Println(ve)
	return nil
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

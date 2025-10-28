package main

import (
	"encoding/json"
	"fmt"

	"github.com/jarusmiselle/calculator/internal/expression"
)

func main() {
	// a()
	// b()
	// c()
	d()
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
	json.Unmarshal([]byte(num), &numV)
	fmt.Println(numV)

	var badV bool
	json.Unmarshal([]byte(bad), &badV)
	fmt.Println(badV)

	var blueV string
	json.Unmarshal([]byte(blue), &blueV)
	fmt.Println(blueV)
}

func c() {
	ex := must(expression.NewBinaryExpression(
		"+",
		expression.ValueExpression(4),
		must(expression.NewUnaryExpression(
			"-",
			expression.ValueExpression(3),
		)),
	))

	fmt.Println(ex)

	b := must(json.Marshal(ex))
	fmt.Println(string(b))
}

func d() {
	v := expression.ValueExpression(4)

	b := must(json.Marshal(v))
	fmt.Println(string(b))
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

package unaryoperation

import (
	"fmt"
	"math"

	"github.com/jarusmiselle/calculator/internal/expression"
)

type UnaryOperation struct {
	symbol    string
	rank      int
	evaluate  func(float64) (float64, error)
	stringify func(expression.Expression) string
}

func (uo UnaryOperation) Symbol() string {
	return uo.symbol
}

func (uo UnaryOperation) Rank() int {
	return uo.rank
}

func (uo UnaryOperation) Evaluate(operand float64) (float64, error) {
	if !uo.IsValid() {
		return 0, fmt.Errorf("invalid unary operation: operation does not exist")
	}
	return uo.evaluate(operand)
}

func (uo UnaryOperation) Stringify(expr expression.Expression) string {
	if !uo.IsValid() {
		return fmt.Sprintf("<<%s>>", expr)
	}
	return uo.stringify(expr)
}

func (uo UnaryOperation) IsValid() bool {
	_, ok := unaryOperationRegistry[uo.symbol]
	return ok
}

var unaryOperationRegistry = make(map[string]UnaryOperation)

var (
	Negation = registerUnaryOperation(
		"-",
		1,
		func(operand float64) (float64, error) {
			return -operand, nil
		},
		func(expr expression.Expression) string {
			return fmt.Sprintf("-%s", expr)
		},
	)

	Factorial = registerUnaryOperation(
		"!",
		2,
		func(operand float64) (float64, error) {
			if operand != math.Trunc(operand) {
				return 0, fmt.Errorf("%f is invalid. factorial is only defined for integers", operand)
			}

			if operand < 0 {
				return 0, fmt.Errorf("%f should be zero or higher", operand)
			}

			intOperand := uint64(operand)
			var res, i uint64 = 1, 2
			for ; i <= intOperand; i++ {
				res *= i
			}

			return float64(res), nil
		},
		func(expr expression.Expression) string {
			return fmt.Sprintf("%s!", expr)
		},
	)
)

func registerUnaryOperation(
	symbol string,
	rank int,
	evaluate func(float64) (float64, error),
	stringify func(expression.Expression) string,
) UnaryOperation {
	unaryOperationRegistry[symbol] = UnaryOperation{
		symbol:    symbol,
		rank:      rank,
		evaluate:  evaluate,
		stringify: stringify,
	}

	return unaryOperationRegistry[symbol]
}

func Parse(symbol string) (UnaryOperation, error) {
	op, ok := unaryOperationRegistry[symbol]
	if !ok {
		return op, fmt.Errorf("invalid unary operator: %s", symbol)
	}

	return op, nil
}

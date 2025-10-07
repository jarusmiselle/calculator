package expression

import (
	"fmt"
	"math"
)

type UnaryExpression struct {
	operation unaryOperation
	operand   Expression
}

func NewUnaryExpression(operation string, operand Expression) (UnaryExpression, error) {
	op, ok := unaryOperations[operation]
	if !ok {
		return UnaryExpression{}, fmt.Errorf("invalid unary operation: %s", operation)
	}

	ue := UnaryExpression{
		operation: op,
		operand:   operand,
	}
	return ue, nil
}

func (ue UnaryExpression) Evaluate() (float64, error) {
	v, err := ue.operand.Evaluate()
	if err != nil {
		return 0, err
	}

	return ue.operation.evaluate(v)
}

func (ue UnaryExpression) Rank() int {
	return ue.operation.rank
}

func (ue UnaryExpression) String() string {
	return ue.operation.stringify(ue.operand)
}

type unaryOperation struct {
	rank      int
	evaluate  func(float64) (float64, error)
	stringify func(Expression) string
}

func factorial(operand float64) (float64, error) {
	if operand != math.Trunc(operand) {
		return 0, fmt.Errorf("%f is invalid. factorial is only defined for integers", operand)
	}

	if operand < 0 {
		return 0, fmt.Errorf("%f should be zero or higher", operand)
	}

	return float64(factorialInt(uint64(operand))), nil
}

func factorialInt(operand uint64) uint64 {
	var res uint64 = 1
	var i uint64 = 2
	for ; i <= operand; i++ {
		res *= i
	}
	return res
}

func negate(operand float64) (float64, error) {
	return -operand, nil
}

var unaryOperations = map[string]unaryOperation{
	"!": {
		rank:      1,
		evaluate:  factorial,
		stringify: func(e Expression) string { return fmt.Sprintf("%s!", e) },
	},
	"-": {
		rank:      2,
		evaluate:  negate,
		stringify: func(e Expression) string { return fmt.Sprintf("-%s", e) },
	},
}

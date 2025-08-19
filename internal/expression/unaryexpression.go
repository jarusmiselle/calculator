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
		return v, err
	}

	return ue.operation.evaluate(v)
}

func (ue UnaryExpression) String() string {
	return ue.operation.stringify(ue.operand)
}

type unaryOperation struct {
	rank      uint8
	evaluate  func(float64) (float64, error)
	stringify func(Expression) string
}

func factorial(operand float64) (float64, error) {
	var res float64

	if operand != math.Trunc(operand) {
		return res, fmt.Errorf("%f is invalid. factorial is only defined for integers", operand)
	}

	if operand < 0 {
		return res, fmt.Errorf("%f should be zero or higher", operand)
	}

	return float64(factorialInt(uint64(operand))), nil
}

func factorialInt(operand uint64) uint64 {
	if operand <= 1 {
		return 1
	}

	return operand * factorialInt(operand)
}

func factorialString(op Expression) string {
	return fmt.Sprintf("%s!", op)
}

func negate(operand float64) (float64, error) {
	return -operand, nil
}

func negateString(op Expression) string {
	return fmt.Sprintf("-%s", op)
}

var unaryOperations = map[string]unaryOperation{
	"!": {
		rank:      1,
		evaluate:  factorial,
		stringify: factorialString,
	},
	"-": {
		rank:      2,
		evaluate:  negate,
		stringify: negateString,
	},
}

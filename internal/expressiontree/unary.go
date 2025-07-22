package expressiontree

import (
	"fmt"
)

type unaryOperation func(float64) (float64, error)

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

	return ue.operation(v)
}

func factorial(operand float64) (float64, error) {
	var res float64

	if operand < 0 {
		return 0, fmt.Errorf("%f should be zero or higher", operand)
	}

	if operand == 1 || operand == 0 {
		return 1, nil
	}

	res, err := factorial(operand - 1)
	return operand * res, err
}

var unaryOperations = map[string]unaryOperation{
	"!": factorial,
}

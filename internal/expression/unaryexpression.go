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

	return ue.operation.function(v)
}

func (ue UnaryExpression) String() string {
	return ue.operation.string(fmt.Sprint(ue.operand))
}

type unaryOperationFunc func(float64) (float64, error)

type unaryOperation struct {
	symbol   rune
	rank     uint8
	function unaryOperationFunc
	string   func(string) string
}

func factorial(operand float64) (float64, error) {
	var res float64

	if operand != math.Trunc(operand) {
		return res, fmt.Errorf("%f is invalid. factorial is only defined for integers", operand)
	}

	if operand < 0 {
		return res, fmt.Errorf("%f should be zero or higher", operand)
	}

	if operand <= 1 {
		return 1, nil
	}

	res, err := factorial(operand - 1)
	return operand * res, err
}

func negate(operand float64) (float64, error) {
	return -operand, nil
}

var unaryOperations = map[string]unaryOperation{
	"!": {
		symbol:   '!',
		rank:     1,
		function: factorial,
		string:   func(s string) string { return fmt.Sprintf("%s!", s) },
	},
	"-": {
		symbol:   '-',
		rank:     2,
		function: negate,
		string:   func(s string) string { return fmt.Sprintf("-%s", s) },
	},
}

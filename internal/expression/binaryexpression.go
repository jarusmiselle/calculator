package expression

import "fmt"

type BinaryExpression struct {
	operation binaryOperation
	left      Expression
	right     Expression
}

func NewBinaryExpression(operation string, left, right Expression) (BinaryExpression, error) {
	op, ok := binaryOperations[operation]
	if !ok {
		return BinaryExpression{}, fmt.Errorf("invalid binary operation: %s", operation)
	}

	be := BinaryExpression{
		operation: op,
		left:      left,
		right:     right,
	}
	return be, nil
}

func (e BinaryExpression) Evaluate() (float64, error) {
	var v float64

	l, err := e.left.Evaluate()
	if err != nil {
		return v, err
	}

	r, err := e.right.Evaluate()
	if err != nil {
		return v, err
	}

	v = l + r
	return v, nil
}

type binaryOperation func(float64, float64) (float64, error)

var binaryOperations = map[string]binaryOperation{}

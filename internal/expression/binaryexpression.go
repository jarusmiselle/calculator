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

	return e.operation(l, r)
}

type binaryOperation func(float64, float64) (float64, error)

func add(left, right float64) (float64, error) {
	res := left + right

	return res, nil
}

func subtract(left, right float64) (float64, error) {
	res := left - right

	return res, nil
}

func multiply(left, right float64) (float64, error) {
	res := left * right

	return res, nil
}

func divide(left, right float64) (float64, error) {
	if right == 0 {
		return 0, fmt.Errorf("invalid: cannot divide by zero")
	}

	res := left / right

	return res, nil
}

var binaryOperations = map[string]binaryOperation{
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

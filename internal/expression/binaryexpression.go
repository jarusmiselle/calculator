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
	l, err := e.left.Evaluate()
	if err != nil {
		return 0, err
	}

	r, err := e.right.Evaluate()
	if err != nil {
		return 0, err
	}

	return e.operation.evaluate(l, r)
}

func (be BinaryExpression) String() string {
	return be.operation.stringify(be.left, be.right)
}

type binaryOperation struct {
	rank      uint8
	evaluate  func(float64, float64) (float64, error)
	stringify func(Expression, Expression) string
}

func add(left, right float64) (float64, error) {
	res := left + right

	return res, nil
}

func addString(left, right Expression) string {
	return fmt.Sprintf("%s+%s", left, right)
}

func subtract(left, right float64) (float64, error) {
	res := left - right

	return res, nil
}

func subtractString(left, right Expression) string {
	return fmt.Sprintf("%s-%s", left, right)
}

func multiply(left, right float64) (float64, error) {
	res := left * right

	return res, nil
}

func multiplyString(left, right Expression) string {
	return fmt.Sprintf("%s*%s", left, right)
}

func divide(left, right float64) (float64, error) {
	if right == 0 {
		return 0, fmt.Errorf("invalid: cannot divide by zero")
	}

	res := left / right

	return res, nil
}

func divideString(left, right Expression) string {
	return fmt.Sprintf("%s/%s", left, right)
}

var binaryOperations = map[string]binaryOperation{
	"+": {
		rank:      4,
		evaluate:  add,
		stringify: addString,
	},
	"-": {
		rank:      4,
		evaluate:  subtract,
		stringify: subtractString,
	},
	"*": {
		rank:      3,
		evaluate:  multiply,
		stringify: multiplyString,
	},
	"/": {
		rank:      3,
		evaluate:  divide,
		stringify: divideString,
	},
}

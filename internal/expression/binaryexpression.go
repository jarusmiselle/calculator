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

func (be BinaryExpression) Rank() int {
	return be.operation.rank
}

func (be BinaryExpression) String() string {
	return fmt.Sprintf(
		be.operation.formatString,
		fmt.Sprintf(
			parens[be.left.Rank() > be.Rank()],
			be.left,
		),
		fmt.Sprintf(
			parens[be.right.Rank() > be.Rank()],
			be.right,
		),
	)
}

type binaryOperation struct {
	rank         int
	evaluate     func(float64, float64) (float64, error)
	formatString string
}

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
	"+": {
		rank:         4,
		evaluate:     add,
		formatString: "%s + %s",
	},
	"-": {
		rank:         4,
		evaluate:     subtract,
		formatString: "%s - %s",
	},
	"*": {
		rank:         3,
		evaluate:     multiply,
		formatString: "%s * %s",
	},
	"/": {
		rank:         3,
		evaluate:     divide,
		formatString: "%s / %s",
	},
}

var parens = map[bool]string{
	true:  "(%s)",
	false: "%s",
}

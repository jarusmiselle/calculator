package expression

import "fmt"

type ValueExpression struct {
	value float64
}

func NewValueExpression(value float64) (ValueExpression, error) {
	return ValueExpression{value: value}, nil
}

func (e ValueExpression) Evaluate() (float64, error) {
	return e.value, nil
}

func (e ValueExpression) String() string {
	return fmt.Sprintf("%f", e.value)
}

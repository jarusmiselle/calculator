package expressiontree

type ValueExpression struct {
	value float64
}

func NewValueExpression(value float64) (ValueExpression, error) {
	return ValueExpression{value: value}, nil
}

func (e ValueExpression) Evaluate() (float64, error) {
	return e.value, nil
}

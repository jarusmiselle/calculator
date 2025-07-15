package expressiontree

type Expression interface {
	Evaluate() (float64, error)
}

type ValueExpression struct {
	Value float64
}

func (e ValueExpression) Evaluate() (float64, error) {
	return e.Value, nil
}

type BinaryExpression struct {
	Left  Expression
	Right Expression
}

func (e BinaryExpression) Evaluate() (float64, error) {
	var v float64

	l, err := e.Left.Evaluate()
	if err != nil {
		return v, err
	}

	r, err := e.Right.Evaluate()
	if err != nil {
		return v, err
	}

	v = l + r
	return v, nil
}

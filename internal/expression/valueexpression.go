package expression

import "fmt"

type ValueExpression float64

func (e ValueExpression) Evaluate() (float64, error) {
	return float64(e), nil
}

func (e ValueExpression) String() string {
	return fmt.Sprintf("%f", e)
}

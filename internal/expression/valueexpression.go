package expression

import (
	"encoding/json"
	"fmt"
)

type ValueExpression float64

var _ Expression = ValueExpression(0)

func (e ValueExpression) Evaluate() (float64, error) {
	return float64(e), nil
}

func (e ValueExpression) Rank() int {
	return 0
}

func (e ValueExpression) String() string {
	return fmt.Sprintf("%f", e)
}

func (v ValueExpression) MarshalJSON() ([]byte, error) {
	j := map[string]any{
		"type":  "ValueExpression",
		"value": float64(v),
	}

	return json.Marshal(j)
}

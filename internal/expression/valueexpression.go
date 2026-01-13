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

const valueExpressionTypeKey = "ValueExpression"

func (v ValueExpression) MarshalJSON() ([]byte, error) {
	type marshaler struct {
		Type  string  `json:"type"`
		Value float64 `json:"value"`
	}

	j := marshaler{
		Type:  valueExpressionTypeKey,
		Value: float64(v),
	}

	return json.Marshal(j)
}

func (v *ValueExpression) UnmarshalJSON(data []byte) error {
	type unmarshaler struct {
		Value float64 `json:"value"`
	}

	var u unmarshaler
	err := json.Unmarshal(data, &u)
	if err != nil {
		return fmt.Errorf("unmarshal ValueExpression: %w", err)
	}

	*v = ValueExpression(u.Value)
	return nil
}

package unaryexpression

import (
	"encoding/json"
	"errors"

	"github.com/jarusmiselle/calculator/internal/expression"
	"github.com/jarusmiselle/calculator/internal/expression/unaryoperation"
)

var ErrInvalidUnaryOperation = errors.New("invalid unary operation")

type UnaryExpression struct {
	operation unaryoperation.UnaryOperation
	operand   expression.Expression
}

func New(operation unaryoperation.UnaryOperation, operand expression.Expression) (UnaryExpression, error) {
	if !operation.IsValid() {
		return UnaryExpression{}, ErrInvalidUnaryOperation
	}

	ue := UnaryExpression{
		operation: operation,
		operand:   operand,
	}
	return ue, nil
}

func (ue UnaryExpression) Evaluate() (float64, error) {
	v, err := ue.operand.Evaluate()
	if err != nil {
		return 0, err
	}

	return ue.operation.Evaluate(v)
}

func (ue UnaryExpression) Rank() int {
	return ue.operation.Rank()
}

func (ue UnaryExpression) String() string {
	return ue.operation.Stringify(ue.operand)
}

const unaryExpressionTypeKey = "UnaryExpression"

func (u UnaryExpression) MarshalJSON() ([]byte, error) {
	type marshaler struct {
		Type      string                `json:"type"`
		Operation string                `json:"operation"`
		Operand   expression.Expression `json:"operand"`
	}

	data := marshaler{
		Type:      unaryExpressionTypeKey,
		Operation: u.operation.Symbol(),
		Operand:   u.operand,
	}

	return json.Marshal(data)
}

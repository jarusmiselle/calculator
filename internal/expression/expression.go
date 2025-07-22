package expression

type Expression interface {
	Evaluate() (float64, error)
}

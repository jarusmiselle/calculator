package calculator

import (
	"errors"
)

type operation func([]float64) (float64, error)

var operations = map[string]operation{
	"add":      add,
	"+":        add,
	"subtract": subtract,
	"sub":      subtract,
	"-":        subtract,
	"multiply": multiply,
	"mul":      multiply,
	"*":        multiply,
	"divide":   divide,
	"div":      divide,
	"/":        divide,
	// "factorial": factorial,
	// "fact":      factorial,
	// "!":         factorial,
}

func add(args []float64) (float64, error) {
	var res float64

	if len(args) != 2 {
		return res, errors.New("usage: calc add left right")
	}

	res = args[0] + args[1]
	return res, nil
}

func subtract(args []float64) (float64, error) {
	var res float64

	if len(args) != 2 {
		return res, errors.New("usage: calc subtract left right")
	}

	res = args[0] - args[1]
	return res, nil
}

func multiply(args []float64) (float64, error) {
	var res float64

	if len(args) != 2 {
		return res, errors.New("usage: calc multiply left right")
	}

	res = args[0] * args[1]
	return res, nil
}

func divide(args []float64) (float64, error) {
	var res float64

	if len(args) != 2 {
		return res, errors.New("usage: calc divide left right")
	}

	if args[1] == 0 {
		return 0, errors.New("right cannot equal zero")
	}

	res = args[0] / args[1]
	return res, nil
}

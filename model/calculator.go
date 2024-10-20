package model

import "errors"

type Calculator interface {
	Add(a, b float64) float64
	Subtract(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) (float64, error)
}

type MyCalculator struct{
	Name string
}

func (mc MyCalculator) Add(a, b float64) float64 {
	return a + b
}

func (mc MyCalculator) Subtract(a, b float64) float64 {
	return a - b
}

func (mc MyCalculator) Multiply(a, b float64) float64 {
	return a * b
}

func (mc MyCalculator) divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Tidak bisa dibagi dengan 0")
	}
	return a / b, nil
}

func (mc MyCalculator) Divide(a, b float64) (float64, error) {
	return mc.divide(a, b)
}
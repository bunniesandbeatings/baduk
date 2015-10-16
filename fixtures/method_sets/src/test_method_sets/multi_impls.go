package test_method_sets

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Circle int

func (r Circle) Perimeter() float64 {
	return 2 * math.Pi * float64(r)
}

func (r Circle) Area() float64 {
	return math.Pi * float64(r*r)
}

type Rectangle struct {
	l int
	w int
}

func (r Rectangle) Perimeter() float64 {
	return 2 * float64(r.l+r.w)
}

func (r Rectangle) Area() float64 {
	return float64(r.l * r.w)
}

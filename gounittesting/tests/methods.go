package tests

import "math"

type vertex struct {
	X, Y float64
}

//a method is just a function with a receiver argument.
func (v vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

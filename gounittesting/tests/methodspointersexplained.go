package tests

import (
	"math"
)

func abs(v vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func scale(v *vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

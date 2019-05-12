package tests

import (
	"fmt"
	"math"
)

type abser interface {
	absTwo() float64
}

type myFloat float64

func (f myFloat) absTwo() float64 {
	fmt.Println("myFloat")
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f myFloat) absThree() float64 {
	fmt.Println("myFloat")
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *vertex) absTwo() float64 {
	fmt.Println("*vertex")
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

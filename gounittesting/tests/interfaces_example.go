package tests

import (
	"fmt"
	"math"
)

// interface for basic shapes:
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circ struct {
	radius float64
}

// implement all methods for type geometry interface:
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circ) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circ) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Printf("Value %v\n", g)
	fmt.Printf("Area %v\n", g.area())
	fmt.Printf("Perimeter %v\n", g.perim())
}

package tests

import (
	"fmt"
	"testing"
)

func TestMeasure(t *testing.T) {
	// 1 Set up
	rectangle := rect{width: 5, height: 10}
	circle := circ{radius: 5}

	// Function Call
	fmt.Println("Testing Rectangle")
	measure(rectangle)
	fmt.Println("Testing Circle")
	measure(circle)

	// Verify
	//assert.True()
}

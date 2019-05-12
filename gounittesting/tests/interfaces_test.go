package tests

import (
	"fmt"
	"math"
	"testing"
)

func TestInterfaceFunctions(t *testing.T) {
	// 1: Set up

	var a abser

	f := myFloat(-math.Sqrt2)
	v := vertex{3, 4}

	a = f  // a myFloat implements abser
	a = &v // a *vertex implements abser

	// In the following line, v is a vertex (not *vertex)
	// and does NOT implement abser.

	// a = v

	// 2: Function Call

	fmt.Println(a.absTwo())

}

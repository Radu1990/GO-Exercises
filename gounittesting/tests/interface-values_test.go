package tests

import (
	"math"
	"testing"
)

func TestInterfaceValues(t *testing.T) {
	var i I
	var ii I

	// 1
	i = &T{"Hello", "Kitty"} // M method has pointer receiver
	describe(i)
	i.M()

	// 2
	ii = F(math.Pi)
	describe(ii)
	ii.M()
}

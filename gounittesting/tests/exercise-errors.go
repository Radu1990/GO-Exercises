package tests

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	/* Computers typically compute the square root
	of x using a loop. Starting with some guess
	z, we can adjust z based on how close z²
	is to x, producing a better guess: */
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	counter := float64(0)
	for {
		counter = counter + float64(1)
		cb := (z*z - x) / (2 * z)
		// separating this piece from the equation
		cbAbs := math.Abs(cb)
		z -= cb
		if cbAbs < 0.001 {
			fmt.Println("Steps needed", counter)
			break
		}
	}
	return z, nil
}

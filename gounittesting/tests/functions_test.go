package tests

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHypotResultValues(t *testing.T) {
	// 1: Set up

	// 2: Function Call

	hypotResult := hypot(5, 12)

	fmt.Println("Result=", hypotResult)

	// 3: Verify

	assert.Equal(t, float64(13), hypotResult, "Expecting 13 for x=5 y=12")
}

func TestComputeHypot(t *testing.T) {
	// 1: Set up

	// 2: Function Call

	computeHypot := compute(hypot)

	fmt.Println("Result=", computeHypot)

	// 3: Verify

	assert.Equal(t, float64(5), computeHypot, "Expecting 5 for fn(3,4)")
}

func TestComputeMathPow(t *testing.T) {
	// 1: Set up

	// 2: Function Call

	computeMathPow := compute(math.Pow)

	fmt.Println("Result=", computeMathPow)

	// 3: Verify

	assert.Equal(t, float64(81), computeMathPow, "Expecting 5 for fn(3,4)")
}

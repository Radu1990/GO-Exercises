package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionScale(t *testing.T) {

	// 1: Set up

	v := vertex{3, 4}

	// 2: Function Call

	scale(&v, 10)
	fmt.Println(abs(v))

	// 3: verify
	assert.Equal(t, float64(50), abs(v), "Expecting 50 for v = vertex{3,4} amd scale f=10")
}

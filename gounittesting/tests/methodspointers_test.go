package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethodsPointersScale(t *testing.T) {
	// 1: Set up

	v := vertex{3, 4}
	fmt.Println("v before scale", v)
	// 2: Function Call

	v.Scale(10)

	fmt.Println("v after scale", v)
	fmt.Println(v.abs())
	fmt.Println("v after abs", v)

	// 3: Verify
	assert.Equal(t, float64(50), v.abs(), "Expecting result = 50 if v.Scale(10) points at the local vertex")
}

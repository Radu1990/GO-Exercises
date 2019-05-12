package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVertexSquareRoot(t *testing.T) {
	// 1: Set up
	v := vertex{3, 4}
	// 2: Function Call
	fmt.Println(v.abs())
	// 3: Verify
	assert.Equal(t, float64(5), v.abs(), "Expecting result = 5 if vertex X=3 Y=4")
}

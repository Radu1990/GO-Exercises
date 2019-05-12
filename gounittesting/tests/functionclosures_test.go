package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnclosureFunctionAdder(t *testing.T) {
	// 1: Set up

	pos := adder()

	// 2: Function Call

	for i := 0; i < 10; i++ {
		fmt.Printf("for i = %d Result = %d\n", i, pos(i))
	}

	// 3: Verify
	assert.Equal(t, 55, pos(10),
		"Expecting result = 55 if 'i' in the row would be 10")
}

func TestEnclosureFunctionAdderSimpleValue(t *testing.T) {
	// 1: Set up

	pos := adder()

	// 2: Function Call

	// 3: Verify
	assert.Equal(t, 10, pos(10), "Expecting result = 10 if i=10")
	assert.Equal(t, 20, pos(10), "Expecting result = 20 if i=10")
	assert.Equal(t, 30, pos(10), "Expecting result = 30 if i=10")
	assert.Equal(t, 40, pos(10), "Expecting result = 40 if i=10")
}

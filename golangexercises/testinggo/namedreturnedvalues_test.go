package testinggo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitGoodValues(t *testing.T) {
	// 1: Set Up

	// 2: Function Call
	x, y := Split(100)

	// 3: Verify
	// $ go get github.com/stretchr/testify
	assert.Equal(t, 44, x, "Custom error message")
	assert.Equal(t, 56, y, "Gooby pls !!!")
}

func TestSplitZero(t *testing.T) {
	// 1: Set Up

	// 2: Function Call
	x, y := Split(0)

	// 3: Verify
	// $ go get github.com/stretchr/testify
	assert.Equal(t, 0, x, "Expected zero as valid input / x output")
	assert.Equal(t, 0, y, "Expected zero as valid input / y output")
}

func TestTrue(t *testing.T) {
	assert.True(t, 0 < 1)
}

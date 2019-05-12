package tests

import "testing"

func TestEmptyInterface(t *testing.T) {

	var i interface{}
	describeEmptyInterface(i)

	i = 42
	describeEmptyInterface(i)

	i = "hello"
	describeEmptyInterface(i)
}


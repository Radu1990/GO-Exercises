package tests

import (
	"fmt"
	"testing"
)

func TestErrNegativeSqrtError(t *testing.T) {

	value1, error1 := Sqrt(2)
	value2, error2 := Sqrt(-2)

	fmt.Println(value1)
	fmt.Println(error1)
	fmt.Println(value2)
	fmt.Println(error2)
}

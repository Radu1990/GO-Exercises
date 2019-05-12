package tests

import (
	"fmt"
	"testing"
)

func TestMyError_Error(t *testing.T) {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

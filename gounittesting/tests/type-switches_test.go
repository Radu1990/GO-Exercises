package tests

import (
	"fmt"
	"testing"
)

func TestDo(t *testing.T) {
	do(21)
	do("hello")
	do(true)
}

func TestPerson(t *testing.T) {
	a := Person{"Radu Andrei", 28}
	b := Person{"Bence",24}
	fmt.Println(a,b)
}

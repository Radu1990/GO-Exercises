package tests

import "fmt"

type I interface {
	M()
}

type T struct {
	R string
	S string
}

func (t *T) M() {
	fmt.Printf("Printing %v and %v\n", t.R, t.S)
}

type F float64

func (f F) M() {
	fmt.Printf("Printing %v\n", f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

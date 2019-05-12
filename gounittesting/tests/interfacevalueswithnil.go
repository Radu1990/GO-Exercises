package tests

import "fmt"

type J interface {
	M()
}

type U struct {
	S string
}

func (t *U) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describeTwo(i J) {
	fmt.Printf("(%v, %T)\n", i, i)
}

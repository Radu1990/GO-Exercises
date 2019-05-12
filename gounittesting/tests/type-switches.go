package tests

import "fmt"

// type switch
func do(i interface{}) {
	switch v:= i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n",v, v*2 )
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// type stringer

type Person struct{
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}


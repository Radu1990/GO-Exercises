package effective_go

import (
	"fmt"
)

type ByteSlice []byte

func pointerFoo(p *ByteSlice) {
	fmt.Println("-p", p)
	fmt.Println("--*p", *p)
	fmt.Println("---&p",&p)
	*p = ByteSlice{6,7,8,0}

}

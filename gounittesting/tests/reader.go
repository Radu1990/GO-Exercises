package tests

import (
	"fmt"
	"io"
	"strings"
)

func reader(s string) {
	r := strings.NewReader(s)
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v v = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}



package effective_go

import (
	"fmt"
	"testing"
)

func TestPointerFoo (t *testing.T) {
		x := ByteSlice{1,2,0,3}

		fmt.Println("x before", x)

		pointerFoo(&x)

		fmt.Println("x after", x)

}

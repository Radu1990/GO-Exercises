package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow { // here _ skips the index, can be replaced with key
		fmt.Printf("%d\n", value) // and then key added here for having both values
	}
}

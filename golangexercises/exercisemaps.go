package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

/*Implement WordCount.
It should return a map of the counts of each “word”
in the string s. The wc.Test function runs a test
 suite against the provided function and prints success or failure. */

var s string = "I am learning Go!"

func WordCount(s string) map[string]int {
	var b []string = strings.Fields(s)
	var slst = make(map[string]int)
	for i := range b {

		if slst[b[i]] == 0 {
			slst[b[i]] = 1
		} else {
			slst[b[i]] = slst[b[i]] + 1
		}

	}
	return slst
}

func main() {
	wc.Test(WordCount)
}

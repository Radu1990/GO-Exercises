package camelproblem

import (
	"fmt"
	"strings"
)
// We are running with go run main.go < camel.in in the Terminal
func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1
	for _, ch := range input {
		str := string(ch)
		if strings.ToUpper(str) == str {
			// It is a capital letter!
			answer++

		}
	}
	fmt.Println(answer)
}

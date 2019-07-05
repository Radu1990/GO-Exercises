package main

import "fmt"

// Maria has n candies and each can have a different flavour (n is always even)
// she gives half to her brother but keeps the highest variety of types of candy
// how many different types of candy can she eat ?
// eg. input = [1, 2, 3, 4, 5, 6] input = [1,1,1,1,1,1] input = [3, 4, 5, 66, 66, 77, 77]


func main(){
	candyArr := [6]int{1,2,3,4,5,6} // array of different types of candies
	requiredNumber:= len(candyArr) / 2 // amount of candies Maria will eat if she gives half away
	fmt.Println("bla bla", candyArr)

}

func makeHistogram (n []int) int {
	m := make(map[int]int)
	for i := 0; i < len(n); i++ {
		if v, _ := m


	}
}
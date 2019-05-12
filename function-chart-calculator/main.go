package function_chart_calculator

import (
	"fmt"
	"math"
)

type convert func(float64) float64

func main() {
	var listNvalues []float64
	listNvalues = generateNvalues(10)
	// n^2*log(n) Computing n raised to the power 2 times log(n)
	fmt.Println("n^2*log(n): ")
	swapFunc(listNvalues, computeNsquareLogN)

	// 2^n Computing 2 raised to the power n
	fmt.Println("2^n: ")
	swapFunc(listNvalues, computeTwoPowN)

	// 2^2^n Computing 2 raised to the power 2 raised to the power n
	fmt.Println("2^2^n: ")
	swapFunc(listNvalues, computeTwoPowTwoPowN)

	// n^log(n) Computing n raised to the power log(n)
	fmt.Println("n^log(n): ")
	swapFunc(listNvalues, computeNpowLogN)
	// n^2 Computing n squared
	fmt.Println("n^2: ")
	swapFunc(listNvalues, computeNsquared)
	//


}
// This function is used to swap the function as needed
func swapFunc (a []float64, fn convert) []float64 {
	var listResults []float64
	for _, val := range a {
		res := fn(val) // desired function changes here
		listResults = append(listResults, res)
	}
	for i:=0; i<len(a); i++{
		fmt.Printf("n=%f res=%f\n", a[i], listResults[i])
	}
	fmt.Println(listResults)
	return listResults
}

// Functions for every equation (more will be added in the future)

func computeNsquareLogN (n float64) float64 {
	ret := math.Pow(n, 2) * math.Log(n)
	return ret
}

func computeTwoPowN (n float64) float64 {
	ret := math.Pow(2, n)
	return ret
}

func computeTwoPowTwoPowN (n float64) float64 {
	a := math.Pow(2, n)
	ret := math.Pow(2, a)
	return ret
}

func computeNpowLogN (n float64) float64 {
	ret := math.Pow(n, math.Log(n))
	return ret
}

func computeNsquared (n float64) float64 {
	ret := math.Pow(n, 2)
	return ret
}

func generateNvalues (a int) []float64 {
	var nval []float64
	for i:=0; i<a+1; i++{
		 nval = append(nval, float64(i))
	}
	return nval
}
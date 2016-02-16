package main

import "fmt"
import "golang_book/chapter5/math"

func main() {
	// A slice of initial length 2 with a total capacity of 2
	t := make([]float64, 2)
	t[0] = 56
	t[1] = 62
	// t[3] = 44			// This would throw an error because the array can only hold 2
	fmt.Println("Array t Average:", math.Average(t))

	// A slice of initial length 5 with a total capacity of 10
	x := make([]float64, 5, 10)
	x[0] = 93
	x[1] = 91
	// the rest of the elements are 0 because they are initialized but never assigned

	fmt.Println("Array x Length:", len(x))
	fmt.Println("Array x Average:", math.Average(x))
}

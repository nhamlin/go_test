package main

import "fmt"

func main() {
	// Using the copy function to get the first two elements in a
	a := []float64{1, 2, 3}
	b := make([]float64, 2)
	copy(b, a)
	fmt.Println(b)
}

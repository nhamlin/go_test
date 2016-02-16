package main

import "fmt"

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}

// The following function names the return type in the declaration line.
// Calling the variable in the return is optional
// If you do not set the variable in the function, and either explicitly call it or have a blank return, the default value is returned
// This is a tricky format because a value is returned whether it's set or not, which could be prone to code smell or unknown bugs
func average(xs []float64) (avg float64) {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	avg = total / float64(len(xs))
	return
}

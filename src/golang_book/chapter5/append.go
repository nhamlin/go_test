package main

import "fmt"

func main() {
	// Using the append function to add three elements to the initialized but never assigned unlimited array
	a := []float64{}
	a = append(a, 1, 2, 3)
	fmt.Println(a)
}

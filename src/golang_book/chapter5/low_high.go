package main

import "fmt"

func main() {
	// Using the [low : high] expression to get the first three elements of arr
	arr := [5]float64{1, 2, 3, 4, 5}
	y := arr[0:2]
	fmt.Println(y)

	// Using the [low : high] expression to start at the 3rd element through the end of the array
	y = arr[2:]
	fmt.Println(y)

	// Using the [low : high] expression to get all elements in the array
	y = arr[:]
	fmt.Println(y)

}

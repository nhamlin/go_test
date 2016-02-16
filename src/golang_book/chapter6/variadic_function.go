package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, 3, 4))

	//We can also pass a slice of ints by following the slice with the ellipsis
	xs := []int{1, 2, 3, 4}
	fmt.Println(add(xs...))
	// The previous line will explode the xs variable into its individual parts and send those parts to the function

	// Doing this will fail, so if it seems like it would work, try exploding the slice into its parts
	//fmt.Println(add(xs))
}

// By using the ellipsis before the type name of the last parameter, you indicate that it takes zero or more of those parameters.
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

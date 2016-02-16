package main

import "fmt"

func main() {
	for x := uint(1); x < 5; x++ {
		fmt.Println("Factorial of", x, "is:", factorial(x))
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

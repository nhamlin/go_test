package main

import "fmt"

func main() {
	xs := []int{1, 2, 3, 4, 5}
	a, b := fbar(xs)
	fmt.Println(a, b)
}

func fbar(xs []int) (int, int) {
	return xs[0], xs[1]
}

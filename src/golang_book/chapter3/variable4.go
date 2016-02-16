package main

import "fmt"

func main() {
	var x string = "hello"
	var y string = "world"
	fmt.Println(x == y)

	x = "hello"
	y = "hello"
	fmt.Println(x == y)

	z := "Hello, World"
	fmt.Println(z)

	a := 5
	fmt.Println(a)

	var b = 7
	fmt.Println(b)
}

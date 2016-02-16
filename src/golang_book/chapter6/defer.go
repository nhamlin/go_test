package main

import "fmt"

func main() {
	defer first()
	defer third()
	second()
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func third() {
	fmt.Println("3rd")
}

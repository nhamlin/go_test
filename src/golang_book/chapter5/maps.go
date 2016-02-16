package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	x["elephant"] = 1
	fmt.Println("Key:", x["key"])

	delete(x, "Key")
	fmt.Println("Key:", x["Key"])
	fmt.Println("elephant", x["elephant"])
}

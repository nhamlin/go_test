package main

import (
	"fmt"
	m "golang_book/chapter8/math" // Use the alias when we want to use two packages with the same name, such as "math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)
	//fmt.Println(m.rand())			// This line won't run because "rand" is not exposed publicly. It's a private function only available to math/math.go
}

package main

import (
	"fmt"
	"math"
)

func main() {
	c1 := new(Circle)              // This creates a pointer which allocates memory for all the fields and sets the to their default value
	c2 := Circle{x: 0, y: 0, r: 5} // This creates a Circle with the defined values
	c3 := Circle{0, 0, 5}          // If you know the order that the values are in, you don't need to include them
	c4 := &Circle{0, 0, 5}         // This creates a pointer to the Circle (like c1) and allocates values to the fields. not a common integration

	fmt.Println(c1.x, c1.y, c1.r)
	fmt.Println(c2.x, c2.y, c2.r)
	fmt.Println(c3.x, c3.y, c3.r)
	fmt.Println(c4.x, c4.y, c4.r)

	fmt.Println(circleArea(c3))
	fmt.Println(c3.r) //c3.r is still 5

	//fmt.Println(circleAreaModifiable(c4))

	fmt.Println(circleAreaModifiable(&c2))
	fmt.Println(c2.r) //c2.r is now 2
}

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	c.r = 2
	return math.Pi * (c.r * c.r)
}

func circleAreaModifiable(c *Circle) float64 {
	c.r = 2
	return math.Pi * (c.r * c.r)
}

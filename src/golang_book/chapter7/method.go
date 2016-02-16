package main

import (
	"fmt"
	"math"
)

func main() {
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
	fmt.Println(area(&r)) //When using the stand-alone function, you must declare the same type (in this case the pointer to a Rectangle)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// The following function attaches itself to the Rectangle struct because it first declares the struct then the name
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

// The following is a stand-alone function that takes a Rectangle as a parameter
func area(r *Rectangle) float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

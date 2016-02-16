package main

import (
	"fmt"
	"math"
)

func main() {
	r1 := Rectangle{2, 2, 4, 4}
	c1 := Circle{0, 0, 5}
	fmt.Println(totalArea([]Circle{c1}, []Rectangle{r1}))
}

// Interfaces define the structure of a type and state what all inherited types MUST have
type Shape interface {
	// The "method set" follows instead of fields
	area() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func totalArea(circles []Circle, rectangles []Rectangle) float64 {
	var total float64
	for _, c := range circles {
		total += c.area()
	}
	for _, r := range rectangles {
		total += r.area()
	}
	return total
}

// The following is a stand-alone function that takes a Rectangle as a parameter
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (c *Circle) area() float64 {
	return math.Pi * (c.r * c.r)
}

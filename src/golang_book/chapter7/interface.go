package main

import (
	"fmt"
	"math"
)

// Note: Go is unique when it comes to how we implement the interfaces we want our types to support.
// Go does not require us to explicitly state that our types implement an interface.
// If every method that belongs to an interface's method set is implemented by the type, then the type is said to implement the interface.

func main() {
	r1 := Rectangle{1, 1, 2, 2}
	c1 := Circle{1, 2, 2}
	fmt.Println(totalArea(&c1, &r1))

	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{1, 1, 2},
			&Rectangle{1, 1, 2, 2},
		},
	}

	fmt.Println(totalArea(&multiShape))
}

// TYPES ----------------------------------------------------------------------

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

// We can also declare an interface made up of many interfaces
type MultiShape struct {
	shapes []Shape
}

// METHODS ------------------------------------------------------------------

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * (c.r * c.r)
}

// FUNCTIONS ------------------------------------------------------------------
func totalArea(shapes ...Shape) float64 { // Using an interface here. Unlike c#, the interface isn't specifically mentioned through IMPLEMENTS or EXTENDS keywords.
	var total float64
	for _, s := range shapes {
		total += s.area()
	}

	return total
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

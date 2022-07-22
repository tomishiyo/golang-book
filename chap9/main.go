/*Add a new method to the Shape interface called perimeter which calculates the
perimeter of a shape. Implement the method for Circle and Rectangle */

/* Note: The point of an interface is that sometimes we need to manipulate
data that all implements the same methods. For instance, in the book example
we want to calculate the total area of many figures, regardless of the
particular details of their nature and calculation. */
package main

import (
	"fmt"
	"math"
)

// Circle
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

// Rectangle
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	a := r.x2 - r.x1
	b := r.y2 - r.y1

	return a * b

}

func (r *Rectangle) perimeter() float64 {
	a := r.x2 - r.x1
	b := r.y2 - r.y1

	return 2*a + 2*b
}

// Interface
type Shape interface {
	area() float64
	perimeter() float64
}

// Interface functions

func total_area(shapes ...Shape) float64 {
	result := 0.0
	for _, shape := range shapes {
		result += shape.area()
	}

	return result

}

func total_perimeter(shapes ...Shape) float64 {
	result := 0.0
	for _, shape := range shapes {
		result += shape.perimeter()
	}

	return result
}

func main() {
	c := Circle{x: 0, y: 0, r: 3}
	r := Rectangle{x1: 0, y1: 0, x2: 3, y2: 5}

	fmt.Println(total_perimeter(&c, &r))

}

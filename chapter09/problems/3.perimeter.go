//Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape. Implement the method for Circle and Rectangle.
package main

import (
	"log"
	"math"
)

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	log.Println(totalArea(&c, &r))
	log.Println(c.perimeter())
	log.Println(r.perimeter())
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

// area() method for Circle
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

// perimeter() method for Circle
func (c *Circle) perimeter() float64 {
	return math.Pi * (c.r * 2)
}

// area() method for Rectangle
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

// perimeter() method for Rectangle
func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return (l + w) * 2
}

// function to get distance between any two points
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
  return math.Sqrt(a*a + b*b)
}

// a function taking a Shape interface
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
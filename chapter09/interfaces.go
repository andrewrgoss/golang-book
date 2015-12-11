package main

import (
	"log"
	"math"
)

func main() {
	c := Circle{0, 0, 5}
	log.Println(c.area()) 
	r := Rectangle{0, 0, 10, 10}
	log.Println(r.area())
	log.Println(totalArea(&c, &r))
}

// area() method for Circle
func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
}

// area() method for Rectangle
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
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

// func (m *Multishape) area() float64 {
// 	var area float64
// 	for _, s := range m.shapes {
// 		area += s.area()
// 	}
// 	return area
// }

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface { //Like a struct an interface is created using the type keyword, followed by a name and the keyword interface. But instead of defining fields, we define a “method set”. A method set is a list of methods that a type must have in order to “implement” the interface.
	area() float64
}

type MultiShape struct {
	shapes []Shape //interfaces can also be used as fields
}
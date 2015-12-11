package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area()) //Go automatically knows to pass a pointer to the circle for this method
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
func (c *Circle) area() float64 {  //Method example. (c *Circle) = receiver. The receiver is like a parameter - it has a name and a type - but by creating the function this way it allows us to call the function using the . operator
  return math.Pi * c.r*c.r
}
func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
  return math.Sqrt(a*a + b*b)
}
type Circle struct {
	x, y, r float64
}
type Rectangle struct {
	x1, y1, x2, y2 float64
}
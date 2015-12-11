package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.x, c.y, c.r) //struct fields can be accessed using the . operator
	fmt.Println(circleArea(&c)) //& operator returns the location in memory where the c variable is stored
}
func circleArea(c *Circle) float64 {
  return math.Pi * c.r*c.r
}
type Circle struct { //a struct is a type which contains named fields
	x, y, r float64
}
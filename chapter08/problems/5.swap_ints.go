//Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).
package main

import "log"

func main() {
	x := 1
	y := 2
	swap(&x, &y)
	log.Println("x=", x)
	log.Println("y=", y)
}

func swap(x *int, y *int) {
	xVal := *x
	yVal := *y
	*x = yVal
	*y = xVal
}
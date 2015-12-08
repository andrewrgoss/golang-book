// What is the value of x after running this program?
package main

import "log"

func main() {
	x := 1.5
	square(&x)
	log.Println(x) // x = 2.25
}

func square(x *float64) {
	*x = *x * *x
}
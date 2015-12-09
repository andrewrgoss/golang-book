package main

import "fmt"

func main() {
  var x [5]int
  x[4] = 100 //set the 5th element of the array to 100
  fmt.Println(x) //[0 0 0 0 100]
}
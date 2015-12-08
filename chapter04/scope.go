package main

import "fmt"

var x := "Hello World" // since variable is outside the main function, it can be accessed by other functions

func main() {
  fmt.Println(x)
}

func f() {
  fmt.Println(x)
}
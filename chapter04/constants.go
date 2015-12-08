package main

import "fmt"

func main() {
  const x string = "Hello World" // Constants are basically variables whose values cannot be changed later. They are created in the same way you create variables but instead of using the var keyword we use the const keyword
  x = "Some other string" // this will produce a compile-time error.
  fmt.Println(x)
  //Constants are a good way to reuse common values in a program without writing them out each time. For example Pi in the math package is defined as a constant.
}
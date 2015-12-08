package main

import "fmt"

//Go shorthand if you need to define multiple variables
var (
	a = 5
	b = 10
	c = 15
)

//Use the keyword var (or const) followed by parentheses with each variable on its own line.

func main() {
	fmt.Println(a, b, c)
}
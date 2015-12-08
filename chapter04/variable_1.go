package main

import "fmt"

func main() {

	//   var x string
	//   x = "first"
	//   fmt.Println(x)
	//   x = "second"
	//   fmt.Println(x)

	var x string
	x = "first "
	fmt.Println(x)
	x += "second" // The x = x + y form is so common in programming that Go has a special assignment statement: +=
	fmt.Println(x)
	
}
package main

import "fmt"

/*Ctrl-K, Ctrl-C	Marks the current line or selected lines of code as a comment, using the correct comment syntax for the programming language
Ctrl-K, Ctrl-U	Removes the comment syntax from the current line or currently selected lines of code
http://www.dofactory.com/reference/visual-studio-shortcuts*/

func main() {
	// var x = "Hello World"
	//x:= "Hello World"
	// fmt.Println(x)
	// var x string
	// x = "first "
	// fmt.Println(x)
	// x += "second"
	// fmt.Println(x)
	// var x = "hello"
	// var y = "hello"
	// fmt.Println(x == y)
	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName) //In this last case we use a special way to represent multiple words in a variable name known as lower camel case (also know as mixed case, bumpy caps, camel back or hump back). The first letter of the first word is lowercase, the first letter of the subsequent words is uppercase and all the other letters are lowercase.
	
	/*According to the language specification “Go is lexically scoped using blocks”. Basically this means that the variable exists within the nearest curly braces { } (a block) including any nested curly braces (blocks), but not outside of them. This is scope.*/
	
	/*Constants are basically variables whose values cannot be changed later.*/
	
	const x string = "Hello World"
	
	var (
  	a = 5
  	b = 10
  	c = 15
	)
	
	fmt.Println(a, b, c)
	fmt.Println("=================================================")
	
	fmt.Print("Enter a number: ")
  	var input float64
  	fmt.Scanf("%f", &input)
  	output := input * 2
  	fmt.Println("Doubled number= ", output)
	fmt.Println("=================================================")
	
	fmt.Print("Enter a number for Fahrenheit to Celsius conversion: ")
	var tempInput float64
	fmt.Scanf("%f", &tempInput)
	tempOutput := ((tempInput - 32) * 5/9)
	fmt.Println("Converted temp to Celsius = ", tempOutput)
	fmt.Println("=================================================")
	
	fmt.Print("Enter a number for feet to meter conversion: ")
	var lenInput float64
	fmt.Scanf("%f", &lenInput)
	lenOutput := lenInput * 0.3048
	fmt.Println("Converted to meters = ", lenOutput)
}

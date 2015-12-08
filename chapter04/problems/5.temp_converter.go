//Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
package main

import "fmt"

func main() {
	fmt.Print("Enter a number for Fahrenheit to Celsius conversion: ")
	var tempInput float64
	fmt.Scanf("%f", &tempInput)
	tempOutput := ((tempInput - 32) * 5/9)
	fmt.Println("Converted temp to Celsius = ", tempOutput)
}
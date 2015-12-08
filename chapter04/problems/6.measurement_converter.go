//Write another program that converts from feet into meters. (1 ft = 0.3048 m)
package main

import "fmt"

func main() {
	fmt.Print("Enter a number for feet to meter conversion: ")
	var lenInput float64
	fmt.Scanf("%f", &lenInput)
	lenOutput := lenInput * 0.3048
	fmt.Println("Converted to meters = ", lenOutput)
}
//What does the following program print:
package main

import "fmt"

func main() { //program will print "Small"
	i := 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}
}
//Modify the program we wrote so that instead of printing Hello World it prints Hello, my name is followed by your name.
package main

import "fmt"

func main() {
	fmt.Print("Enter your name: ")
  	var myName string
  	fmt.Scanf("%s", &myName) //%s = reads in the uninterpreted bytes of the string or slice
	  //&myName - the & operand finds the memory location of the 'myName' variable and then using a pointer, modifies the original variable value (to the user input)

/*func Scanf(format string, a ...interface{}) (n int, err error)
Scanf scans text read from standard input, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why. Newlines in the input must match newlines in the format. The one exception: the verb %c always scans the next rune in the input, even if it is a space (or tab etc.) or newline.*/

  	fmt.Println("Hello, my name is", myName)
}
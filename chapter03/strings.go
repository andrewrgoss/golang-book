package main

import "fmt"

func main() {
	fmt.Println(len("Hello World")) // A space is also considered a character, so the string's length is 11 not 10 and the 3rd line has "Hello " instead of "Hello".
	fmt.Println("Hello World"[1]) // Strings are “indexed” starting at 0 not 1. [1] gives you the 2nd element not the 1st. Also notice that you see 101 instead of e when you run this program. This is because the character is represented by a byte (remember a byte is an integer).
	fmt.Println("Hello " + "World") // Concatenation uses the same symbol as addition. The Go compiler figures out what to do based on the types of the arguments. Since both sides of the + are strings the compiler assumes you mean concatenation and not addition. (Addition is meaningless for strings)
}
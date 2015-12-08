//https://www.golang-book.com/books/intro/7#section5
package main

import "log"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1) //factorial calls itself, which is what makes this function recursive
}

func main() {
	log.Println(factorial(2))
}
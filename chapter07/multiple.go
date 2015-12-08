//https://www.golang-book.com/books/intro/7#section2
package main

//import "log"

func f() (int, int) {
	return 5,6
}

func main() {
	x,y := f()
}
//https://www.golang-book.com/books/intro/7#section3
package main

import "log"

func add(args ...int) int { //By using ... before the type name of the last parameter you can indicate that it takes zero or more of those parameters. In this case we take zero or more ints. We invoke the function like any other function except we can pass as many ints as we want.
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	log.Println(add(1,2,3))
	xs := []int{1,2,3}
	log.Println(add(xs...))
}
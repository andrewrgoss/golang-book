//https://www.golang-book.com/books/intro/7#section4
package main

import "log"

//It is possible to create functions inside of functions:
func main() {
	add := func(x, y int) int { //function that takes two ints and returns an int
		return x + y
	}
	
	x := 0
	increment := func() int {
		x++
		return x
	}
	
	//A function like this together with the non-local variables it references is known as a closure. In this case increment and the variable x form the closure.
	
	log.Println(add(1,1))
	log.Println(increment())
	log.Println(increment())
	
	nextEven := makeEvenGenerator()
	log.Println(nextEven()) // 0
	log.Println(nextEven()) // 2
	log.Println(nextEven()) // 4
	
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
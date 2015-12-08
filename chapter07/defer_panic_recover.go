//https://www.golang-book.com/books/intro/7#section6
package main

import "log"

func first() {
	log.Println("1st")
}
func second() {
	log.Println("2nd")
}
func main() {
	defer second() //defer is often used when resources need to be freed in some way.
	first()
	
	// For example, when we open a file we need to make sure to close it later. With defer:
	
	// f, _ := os.Open(filename)
	// defer f.Close()
	
	/*This has 3 advantages: (1) it keeps our Close call near our Open call so it's easier to understand, (2) if our function had multiple return statements (perhaps one in an if and one in an else) Close will happen before both of them and (3) deferred functions are run even if a run-time panic occurs.*/
	
	defer func() {
		str := recover()
		log.Println(str)
	}()
	panic("PANIC")
	
	/*A panic generally indicates a programmer error (for example attempting to access an index of an array that's out of bounds, forgetting to initialize a map, etc.) or an exceptional condition that there's no easy way to recover from. (Hence the name “panic”)*/
}
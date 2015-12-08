package main

import "log"

func main() {
	x := 5
	zero(&x)
	log.Println(x) // x is 0
	
	yPtr := new(int) //new takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
	one(yPtr)
	log.Println(*yPtr) // y is 1
}

//Pointers reference a location in memory where a value is stored rather than the value itself. (They point to something else) By using a pointer (*int) the zero function is able to modify the original variable. A pointer is a special data type.

func zero(xPtr *int) {
	*xPtr = 0 //When we write *xPtr = 0 we are saying “store the int 0 in the memory location xPtr refers to”. If we try xPtr = 0 instead we will get a compiler error because xPtr is not an int it's a *int, which can only be given another *int.
}

func one(yPtr *int) {
	*yPtr = 1
}
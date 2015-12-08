/*Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).*/
package main

import "log"

func main() {
	log.Println(half(12))
}

func half(num int)(int, bool) {
	if num % 2 == 0 {
		return num/2,true	
	} 
		return num/2,false
}
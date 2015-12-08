package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(fib(uint(i)))
	}
}

func fib(n uint)uint {
	if n == 0 || n == 1 {
		return n
	}
	return (n-1) + (n-2)
}
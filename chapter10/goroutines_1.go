package main

import "fmt"

func main() { // implicit goroutine
	go f(0) // 2nd goroutine gets created here. Normally when we invoke a function our program will execute all the statements in a function and then return to the next line following the invocation. With a goroutine we return immediately to the next line and don't wait for the function to complete. This is why the call to the Scanln function has been included; without it the program would exit before being given the opportunity to print all the numbers.
	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
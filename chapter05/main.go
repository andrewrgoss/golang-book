package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
	
	fmt.Println("\n======================================================")
	
	for j := 0; j <= 5; j++ {
		switch j {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			default: fmt.Println("Unknown Number")
		}
	}
	
	fmt.Println("\n======================================================")
	
	for k := 1; k <= 100; k++ {
		if k % 3 == 0 {
			fmt.Print(k, ",")
		}
	}
	
	fmt.Println("\n======================================================")
	
	for l := 1; l <= 100; l++ {
		if l % 3 == 0 {
			fmt.Println(l, "Fizz")
		}
	 	if l % 5 == 0 {
			fmt.Println(l, "Buzz")
		}
		if l % 3 == 0 && l % 5 == 0 {
			fmt.Println(l, "FizzBuzz")
		}
	}
}